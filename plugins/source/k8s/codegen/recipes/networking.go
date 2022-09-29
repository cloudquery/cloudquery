package recipes

import networkingv1 "k8s.io/api/networking/v1"

func NetworkingResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "network_policies",
			Struct:     &networkingv1.NetworkPolicy{},
		},
	}

	for _, resource := range resources {
		resource.Service = "networking"
	}

	return resources
}
