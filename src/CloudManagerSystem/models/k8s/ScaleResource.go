package k8s

import (
	"strconv"
	"strings"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
)

// ReplicaCounts provide the desired and actual number of replicas.
type ReplicaCounts struct {
	DesiredReplicas int32 `json:"desiredReplicas"`
	ActualReplicas  int32 `json:"actualReplicas"`
}
// ScaleResource scales the provided resource using the client scale method in the case of Deployment,
// ReplicaSet, Replication Controller. In the case of a job we are using the jobs resource update
// method since the client scale method does not provide one for the job.
func ScaleResource(client client.Interface, kind, namespace, name, count string) (rc *ReplicaCounts, err error) {
	rc = new(ReplicaCounts)
	if strings.ToLower(kind) == "job" {
		err = scaleJobResource(client, namespace, name, count, rc)
	} else if strings.ToLower(kind) == "statefulset" {
		err = scaleStatefulSetResource(client, namespace, name, count, rc)
	} else {
		err = scaleGenericResource(client, kind, namespace, name, count, rc)
	}
	if err != nil {
		return nil, err
	}

	return
}

//ScaleGenericResource is used for Deployment, ReplicaSet, Replication Controller scaling.
func scaleGenericResource(client client.Interface, kind, namespace, name, count string, rc *ReplicaCounts) error {
	s, err := client.ExtensionsV1beta1().Scales(namespace).Get(kind, name)
	if err != nil {
		return err
	}
	c, err := strconv.Atoi(count)
	if err != nil {
		return err
	}
	s.Spec.Replicas = int32(c)
	s, err = client.ExtensionsV1beta1().Scales(namespace).Update(kind, s)
	if err != nil {
		return err
	}
	rc.DesiredReplicas = s.Spec.Replicas
	rc.ActualReplicas = s.Status.Replicas

	return nil
}

// scaleJobResource is exclusively used for jobs as it does not increase/decrease pods but jobs parallelism attribute.
func scaleJobResource(client client.Interface, namespace, name, count string, rc *ReplicaCounts) error {
	j, err := client.BatchV1().Jobs(namespace).Get(name, metaV1.GetOptions{})
	c, err := strconv.Atoi(count)
	if err != nil {
		return err
	}

	*j.Spec.Parallelism = int32(c)
	j, err = client.BatchV1().Jobs(namespace).Update(j)
	if err != nil {
		return err
	}

	rc.DesiredReplicas = *j.Spec.Parallelism
	rc.ActualReplicas = *j.Spec.Parallelism

	return nil
}

// scaleStatefulSet is exclusively used for statefulsets
func scaleStatefulSetResource(client client.Interface, namespace, name, count string, rc *ReplicaCounts) error {
	ss, err := client.AppsV1beta1().StatefulSets(namespace).Get(name, metaV1.GetOptions{})
	c, err := strconv.Atoi(count)
	if err != nil {
		return err
	}

	*ss.Spec.Replicas = int32(c)
	ss, err = client.AppsV1beta1().StatefulSets(namespace).Update(ss)
	if err != nil {
		return err
	}

	rc.DesiredReplicas = *ss.Spec.Replicas
	rc.ActualReplicas = ss.Status.Replicas

	return nil
}

