package recipes

import (
	resource "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

func Rbac() []*Resource {
	resources := []*Resource{
		{
			SubService:     "cluster_roles",
			Struct:         &resource.ClusterRole{},
			ResourceFunc:   resourceType.ClusterRolesGetter.ClusterRoles,
			GlobalResource: true,
		},
		{
			SubService:     "cluster_role_bindings",
			Struct:         &resource.ClusterRoleBinding{},
			ResourceFunc:   resourceType.ClusterRoleBindingsGetter.ClusterRoleBindings,
			GlobalResource: true,
		},
		{
			SubService:   "roles",
			Struct:       &resource.Role{},
			ResourceFunc: resourceType.RolesGetter.Roles,
		},
		{
			SubService:   "role_bindings",
			Struct:       &resource.RoleBinding{},
			ResourceFunc: resourceType.RoleBindingsGetter.RoleBindings,
		},
	}

	for _, resource := range resources {
		resource.Service = "rbac"
		resource.ServiceFunc = kubernetes.Interface.RbacV1
	}

	return resources
}
