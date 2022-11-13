package recipes

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

func RbacResources() []*Resource {
	resources := []*Resource{
		{
			SubService:     "cluster_roles",
			Struct:         &rbacv1.ClusterRole{},
			ResourceFunc:   v1.ClusterRolesGetter.ClusterRoles,
			GlobalResource: true,
		},
		{
			SubService:     "cluster_role_bindings",
			Struct:         &rbacv1.ClusterRoleBinding{},
			ResourceFunc:   v1.ClusterRoleBindingsGetter.ClusterRoleBindings,
			GlobalResource: true,
		},
		{
			SubService:   "roles",
			Struct:       &rbacv1.Role{},
			ResourceFunc: v1.RolesGetter.Roles,
		},
		{
			SubService:   "role_bindings",
			Struct:       &rbacv1.RoleBinding{},
			ResourceFunc: v1.RoleBindingsGetter.RoleBindings,
		},
	}

	for _, resource := range resources {
		resource.Service = "rbac"
		resource.ServiceFunc = kubernetes.Interface.RbacV1
	}

	return resources
}
