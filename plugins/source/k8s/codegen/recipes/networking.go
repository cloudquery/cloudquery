package recipes

import (
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

func NetworkingResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "network_policies",
			Struct:     &networkingv1.NetworkPolicy{},
			ResourceFunc: v1.NetworkPoliciesGetter.NetworkPolicies,
		},
	}

	for _, resource := range resources {
		resource.Service = "networking"
		resource.ServiceFunc = kubernetes.Interface.NetworkingV1
		resource.SkipMockFields = []string{"Port"}
	}

	return resources
}
