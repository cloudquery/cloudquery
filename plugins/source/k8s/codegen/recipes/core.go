package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	corev1 "k8s.io/api/core/v1"
)

func CoreResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "endpoints",
			Struct:     &corev1.Endpoints{},
		},
		{
			SubService: "limit_ranges",
			Struct:     &corev1.LimitRange{},
		},
		{
			SubService: "namespaces",
			Struct:     &corev1.Namespace{},
		},
		{
			SubService: "nodes",
			Struct:     corev1.Node{},
			SkipFields: []string{
				"PodCIDR",
				"PodCIDRs",
				"DoNotUseExternalID", // Deprecated
			},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "spec_pod_cidr",
					Type:     schema.TypeCIDR,
					Resolver: `client.StringToCidrPathResolver("Spec.PodCIDR")`,
				},
				{
					Name:     "spec_pod_cidrs",
					Type:     schema.TypeCIDRArray,
					Resolver: `schema.PathResolver("Spec.PodCIDRs")`,
				},
			},
		},
		{
			SubService: "pods",
			Struct:     &corev1.Pod{},
			SkipFields: []string{
				"HostIP",
				"PodIP",
				"PodIPs",
				"DeprecatedServiceAccount", // Deprecated
			},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "status_host_ip",
					Type:     schema.TypeInet,
					Resolver: `client.StringToInetPathResolver("Status.HostIP")`,
				},
				{
					Name:     "status_pod_ip",
					Type:     schema.TypeInet,
					Resolver: `client.StringToInetPathResolver("Status.PodIP")`,
				},
				{
					Name:     "status_pod_ips",
					Type:     schema.TypeInetArray,
					Resolver: `resolveCorePodPodIPs`,
				},
			},
		},
		{
			SubService: "resource_quotas",
			Struct:     &corev1.ResourceQuota{},
		},
		{
			SubService: "secrets",
			Struct:     &corev1.Secret{},
			SkipFields: []string{"Data", "StringData"},
		},
		{
			SubService: "service_accounts",
			Struct:     &corev1.ServiceAccount{},
		},
		{
			SubService: "services",
			Struct:     &corev1.Service{},
			SkipFields: []string{"ClusterIP", "ClusterIPs", "ExternalIPs", "LoadBalancerIP"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "spec_cluster_ip",
					Type:     schema.TypeInet,
					Resolver: `client.StringToInetPathResolver("Spec.ClusterIP")`,
				},
				{
					Name:     "spec_cluster_ips",
					Type:     schema.TypeInetArray,
					Resolver: `schema.PathResolver("Spec.ClusterIPs")`,
				},
				{
					Name:     "spec_external_ips",
					Type:     schema.TypeInetArray,
					Resolver: `schema.PathResolver("Spec.ExternalIPs")`,
				},
				{
					Name:     "spec_load_balancer_ip",
					Type:     schema.TypeInet,
					Resolver: `client.StringToInetPathResolver("Spec.LoadBalancerIP")`,
				},
			},
		},
	}

	for _, resource := range resources {
		resource.Service = "core"
	}

	return resources
}
