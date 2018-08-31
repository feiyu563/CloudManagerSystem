package k8s

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
	authv1 "k8s.io/api/rbac/v1"
)

func K8sCreateClusterRole(client client.Interface, RoleName string) error {

	clusterRole := &authv1.ClusterRole{
		ObjectMeta: metaV1.ObjectMeta{
			Annotations: map[string]string{authv1.AutoUpdateAnnotationKey: "true"},
			Labels:      map[string]string{"kubernetes.io/bootstrapping": "rbac-cloudmanager"},
			Name:        RoleName,
		},
		Rules: []authv1.PolicyRule{
			{
				APIGroups: []string{authv1.APIGroupAll},
				Resources: []string{authv1.ResourceAll},//the first version add all resource
				Verbs:     []string{authv1.VerbAll},//the first version add all
				//NonResourceURLs:[]string{authv1.NonResourceAll},
			},
		},
	}
	_, err := client.RbacV1().ClusterRoles().Create(clusterRole)
	if err != nil {
		// TODO(bryk): Roll back created resources in case of error.
		return err
	}
	return nil
}

func K8sDeleteClusterRole(client client.Interface, RoleName string) error {

	err := client.RbacV1().ClusterRoles().Delete(RoleName,&metaV1.DeleteOptions{})
	if err != nil {
		// TODO(bryk): Roll back created resources in case of error.
		return err
	}
	return nil
}
