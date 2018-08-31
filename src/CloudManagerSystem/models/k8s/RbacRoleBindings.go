package k8s

import (
	rbac "k8s.io/api/rbac/v1"
)
// RbacRoleBindingList contains a list of Roles and ClusterRoles in the cluster.
type RbacRoleBindingList struct {
	ListMeta ListMeta `json:"listMeta"`

	// Unordered list of RbacRoleBindings
	Items []RbacRoleBinding `json:"items"`
}

// RbacRoleBinding provides the simplified, combined presentation layer view of Kubernetes' RBAC RoleBindings and ClusterRoleBindings.
// ClusterRoleBindings will be referred to as RoleBindings for the namespace "all namespaces".
type RbacRoleBinding struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`
	Subjects   []rbac.Subject `json:"subjects"`
	RoleRef    rbac.RoleRef   `json:"roleRef"`
	Name       string         `json:"name"`
	Namespace  string         `json:"namespace"`
}

