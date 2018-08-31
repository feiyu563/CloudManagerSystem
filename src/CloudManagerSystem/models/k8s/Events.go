package k8s

import (
	//"CloudManagerSystem/models"
	//"fmt"
	"strings"
	"k8s.io/apimachinery/pkg/types"
	//metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
	api "k8s.io/api/core/v1"
	client "k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/fields"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"log"
)

// EmptyEventList is a empty list of events.
var EmptyEventList = &EventList{
	Events: make([]Event, 0),
	ListMeta: ListMeta{
		TotalItems: 0,
	},
}

type Event struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// A human-readable description of the status of related object.
	Message string `json:"message"`

	// Component from which the event is generated.
	SourceComponent string `json:"sourceComponent"`

	// Host name on which the event is generated.
	SourceHost string `json:"sourceHost"`

	// Reference to a piece of an object, which triggered an event. For example
	// "spec.containers{name}" refers to container within pod with given name, if no container
	// name is specified, for example "spec.containers[2]", then it refers to container with
	// index 2 in this pod.
	SubObject string `json:"object"`

	// The number of times this event has occurred.
	Count int32 `json:"count"`

	// The time at which the event was first recorded.
	FirstSeen metaV1.Time `json:"firstSeen"`

	// The time at which the most recent occurrence of this event was recorded.
	LastSeen metaV1.Time `json:"lastSeen"`

	// Short, machine understandable string that gives the reason
	// for this event being generated.
	Reason string `json:"reason"`

	// Event type (at the moment only normal and warning are supported).
	Type string `json:"type"`
}

type EventList struct {
	ListMeta ListMeta `json:"listMeta"`

	// List of events from given namespace.
	Events []Event `json:"events"`
}

// GetPodsEvents gets events targeting given list of pods.
func GetPodsEvents(client client.Interface, namespace string, pods []v1.Pod) (
	[]v1.Event, error) {

	nsQuery := NewSameNamespaceQuery(namespace)
	if namespace == v1.NamespaceAll {
		nsQuery = NewNamespaceQuery([]string{})
	}

	channels := &ResourceChannels{
		EventList: GetEventListChannel(client, nsQuery.ToRequestParam()),
	}

	eventList := <-channels.EventList.List
	//if err := <-channels.EventList.Error; err != nil {
	//	return nil, err
	//}

	events := filterEventsByPodsUID(eventList.Items, pods)

	return events, nil
}

// GetServiceEvents returns model events for a service with the given name in the given namespace.
func GetServiceEvents(client client.Interface, namespace, name string) (*EventList, error) {
	eventList := EventList{
		Events:   make([]Event, 0),
		ListMeta: ListMeta{TotalItems: 0},
	}

	serviceEvents, err := GetEvents(client, namespace, name)
	if err != nil {
		return &eventList, err
	}

	eventList = CreateEventList(FillEventsType(serviceEvents))
	log.Printf("Found %d events related to %s service in %s namespace", len(eventList.Events), name, namespace)
	return &eventList, nil
}


// Returns filtered list of event objects. Events list is filtered to get only events targeting
// pods on the list.
func filterEventsByPodsUID(events []api.Event, pods []api.Pod) []api.Event {
	result := make([]api.Event, 0)
	podEventMap := make(map[types.UID]bool, 0)

	if len(pods) == 0 || len(events) == 0 {
		return result
	}

	for _, pod := range pods {
		podEventMap[pod.UID] = true
	}

	for _, event := range events {
		if _, exists := podEventMap[event.InvolvedObject.UID]; exists {
			result = append(result, event)
		}
	}

	return result
}

// GetPodsEventWarnings returns warning pod events by filtering out events targeting only given pods
// TODO(floreks) : Import and use Set instead of custom function to get rid of duplicates
func GetPodsEventWarnings(events []api.Event, pods []api.Pod) []Event {
	result := make([]Event, 0)

	// Filter out only warning events
	events = getWarningEvents(events)
	failedPods := make([]api.Pod, 0)

	// Filter out ready and successful pods
	for _, pod := range pods {
		if !isReadyOrSucceeded(pod) {
			failedPods = append(failedPods, pod)
		}
	}

	// Filter events by failed pods UID
	events = filterEventsByPodsUID(events, failedPods)
	events = removeDuplicates(events)

	for _, event := range events {
		result = append(result, Event{
			Message: event.Message,
			Reason:  event.Reason,
			Type:    event.Type,
		})
	}

	return result
}

// GetResourceEvents gets events associated to specified resource.
func GetResourceEvents(client client.Interface, namespace, name string) (*EventList, error) {
	resourceEvents, err := GetEvents(client, namespace, name)
	if err != nil {
		return EmptyEventList, err
	}

	events := CreateEventList(resourceEvents)
	return &events, nil
}

// CreateEventList converts array of api events to common EventList structure
func CreateEventList(events []v1.Event) EventList {
	eventList := EventList{
		Events:   make([]Event, 0),
		ListMeta: ListMeta{TotalItems: len(events)},
	}

	//events = fromCells(dataselect.GenericDataSelect(toCells(events), dsQuery))
	for _, event := range events {
		eventDetail := ToEvent(event)
		eventList.Events = append(eventList.Events, eventDetail)
	}

	return eventList
}

// ToEvent converts event api Event to Event model object.
func ToEvent(event v1.Event) Event {
	result := Event{
		ObjectMeta:      NewObjectMeta(event.ObjectMeta),
		TypeMeta:        NewTypeMeta(ResourceKindEvent),
		Message:         event.Message,
		SourceComponent: event.Source.Component,
		SourceHost:      event.Source.Host,
		SubObject:       event.InvolvedObject.FieldPath,
		Count:           event.Count,
		FirstSeen:       event.FirstTimestamp,
		LastSeen:        event.LastTimestamp,
		Reason:          event.Reason,
		Type:            event.Type,
	}

	return result
}

// GetEvents gets events associated to resource with given name.
func GetEvents(client client.Interface, namespace, resourceName string) ([]v1.Event, error) {
	fieldSelector, err := fields.ParseSelector("involvedObject.name" + "=" + resourceName)

	if err != nil {
		return nil, err
	}

	channels := &ResourceChannels{
		EventList: GetEventListChannelWithOptions(
			client,
			namespace,
			metaV1.ListOptions{
				LabelSelector: labels.Everything().String(),
				FieldSelector: fieldSelector.String(),
			},
			),
	}

	eventList := <-channels.EventList.List
	if err := <-channels.EventList.Error; err != nil {
		return nil, err
	}

	return FillEventsType(eventList.Items), nil
}


var FailedReasonPartials = []string{"failed", "err", "exceeded", "invalid", "unhealthy",
	"mismatch", "insufficient", "conflict", "outof", "nil", "backoff"}
// Returns filtered list of event objects.
// Event list object is filtered to get only warning events.
func getWarningEvents(events []api.Event) []api.Event {
	return filterEventsByType(FillEventsType(events), api.EventTypeWarning)
}

// Filters kubernetes API event objects based on event type.
// Empty string will return all events.
func filterEventsByType(events []api.Event, eventType string) []api.Event {
	if len(eventType) == 0 || len(events) == 0 {
		return events
	}

	result := make([]api.Event, 0)
	for _, event := range events {
		if event.Type == eventType {
			result = append(result, event)
		}
	}

	return result
}

// Based on event Reason fills event Type in order to allow correct filtering by Type.
func FillEventsType(events []v1.Event) []v1.Event {
	for i := range events {
		// Fill in only events with empty type.
		if len(events[i].Type) == 0 {
			if isFailedReason(events[i].Reason, FailedReasonPartials...) {
				events[i].Type = v1.EventTypeWarning
			} else {
				events[i].Type = v1.EventTypeNormal
			}
		}
	}

	return events
}

func isFailedReason(reason string, partials ...string) bool {
	for _, partial := range partials {
		if strings.Contains(strings.ToLower(reason), partial) {
			return true
		}
	}

	return false
}

// Returns true if given pod is in state ready or succeeded, false otherwise
func isReadyOrSucceeded(pod api.Pod) bool {
	if pod.Status.Phase == api.PodSucceeded {
		return true
	}
	if pod.Status.Phase == api.PodRunning {
		for _, c := range pod.Status.Conditions {
			if c.Type == api.PodReady {
				if c.Status == api.ConditionFalse {
					return false
				}
			}
		}

		return true
	}

	return false
}

// Removes duplicate strings from the slice
func removeDuplicates(slice []api.Event) []api.Event {
	visited := make(map[string]bool, 0)
	result := make([]api.Event, 0)

	for _, elem := range slice {
		if !visited[elem.Reason] {
			visited[elem.Reason] = true
			result = append(result, elem)
		}
	}

	return result
}
