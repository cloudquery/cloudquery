package recipes

import rbacv1 "k8s.io/api/rbac/v1"

func RbacResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "role_bindings",
			Struct:     &rbacv1.RoleBinding{},
		},
		{
			SubService: "roles",
			Struct:     &rbacv1.Role{},
		},
	}

	for _, resource := range resources {
		resource.Service = "rbac"
	}

	return resources
}
