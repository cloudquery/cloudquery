package recipes

import (
	resource "k8s.io/api/networking/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/networking/v1"
)

func Networking() []*Resource {
	resources := []*Resource{
		{
			SubService:   "ingresses",
			Struct:       &resource.Ingress{},
			ResourceFunc: resourceType.IngressesGetter.Ingresses,
		},
		{
			SubService:   "network_policies",
			Struct:       &resource.NetworkPolicy{},
			ResourceFunc: resourceType.NetworkPoliciesGetter.NetworkPolicies,
			FakerOverride: `
			r.Spec.Ingress = []resource.NetworkPolicyIngressRule{}
			r.Spec.Egress = []resource.NetworkPolicyEgressRule{}
			`,
		},
		{
			SubService:     "ingress_classes",
			Struct:         &resource.IngressClass{},
			ResourceFunc:   resourceType.IngressClassesGetter.IngressClasses,
			GlobalResource: true,
		},
	}

	for _, resource := range resources {
		resource.Service = "networking"
		resource.ServiceFunc = kubernetes.Interface.NetworkingV1
	}

	return resources
}
