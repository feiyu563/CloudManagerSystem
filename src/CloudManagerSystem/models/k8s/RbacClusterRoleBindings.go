package k8s

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
	authv1 "k8s.io/api/rbac/v1"
	//"fmt"
)

func K8sCreateClusterRoleBinding(client client.Interface, ClusterRoleName string, GroupName string,namespace string) error {
	rolbinding := &authv1.ClusterRoleBinding{
		ObjectMeta: metaV1.ObjectMeta{
			Annotations: map[string]string{authv1.AutoUpdateAnnotationKey: "true"},
			Labels:      map[string]string{"kubernetes.io/bootstrapping": "rbac-cloudmanager"},
			Name:        GroupName,
			Namespace:   namespace,
			//ResourceVersion:authv1.ResourceAll,//"76",
		},
		RoleRef: authv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     ClusterRoleName,
		},
		Subjects: []authv1.Subject{
			{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     authv1.GroupKind, //the first add group
				Name:     GroupName,
			},
		},
	}
	_, err := client.RbacV1().ClusterRoleBindings().Create(rolbinding) //("default")
	if err != nil {
		// TODO(bryk): Roll back created resources in case of error.
		return err
	}
	return nil
}


func K8sDeleteClusterRoleBinding(client client.Interface, GroupName string) error {

	err := client.RbacV1().ClusterRoleBindings().Delete(GroupName,&metaV1.DeleteOptions{}) //("default")
	if err != nil {
		// TODO(bryk): Roll back created resources in case of error.
		return err
	}
	return nil
}
