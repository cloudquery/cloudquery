package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func CoreResources() []*Resource {
	resources := []*Resource{
		{
			SubService:     "component_statuses",
			Struct:         &corev1.ComponentStatus{},
			ResourceFunc:   v1.ComponentStatusesGetter.ComponentStatuses,
			GlobalResource: true,
		},
		{
			SubService:   "config_maps",
			Struct:       &corev1.ConfigMap{},
			ResourceFunc: v1.ConfigMapsGetter.ConfigMaps,
		},
		{
			SubService:   "endpoints",
			Struct:       &corev1.Endpoints{},
			ResourceFunc: v1.EndpointsGetter.Endpoints,
		},
		{
			SubService:   "events",
			Struct:       &corev1.Event{},
			ResourceFunc: v1.EventsGetter.Events,
		},
		{
			SubService:   "limit_ranges",
			Struct:       &corev1.LimitRange{},
			ResourceFunc: v1.LimitRangesGetter.LimitRanges,
		},
		{
			SubService:     "namespaces",
			Struct:         &corev1.Namespace{},
			GlobalResource: true,
			ResourceFunc:   v1.NamespacesGetter.Namespaces,
		},
		{
			SubService:     "namespaces",
			Struct:         &corev1.Namespace{},
			GlobalResource: true,
			ResourceFunc:   v1.NamespacesGetter.Namespaces,
		},
		{
			SubService:     "nodes",
			Struct:         corev1.Node{},
			GlobalResource: true,
			ResourceFunc:   v1.NodesGetter.Nodes,
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
			MockFieldsValue: map[string]string{
				"PodCIDR":  `"8.8.8.8"`,
				"PodCIDRs": `[]string{"1.1.1.1"}`,
			},
		},
		{
			SubService:     "pvs",
			Struct:         &corev1.PersistentVolume{},
			ResourceFunc:   v1.PersistentVolumesGetter.PersistentVolumes,
			GlobalResource: true,
		},
		{
			SubService:   "pvcs",
			Struct:       &corev1.PersistentVolumeClaim{},
			ResourceFunc: v1.PersistentVolumeClaimsGetter.PersistentVolumeClaims,
		},
		{
			SubService:   "pods",
			Struct:       &corev1.Pod{},
			ResourceFunc: v1.PodsGetter.Pods,
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
			MockFieldsValue: map[string]string{
				"HostIP": `"8.8.8.8"`,
				"PodIP":  `"8.8.8.8"`,
				"PodIPs": `[]resource.PodIP{resource.PodIP{IP: "8.8.8.8"}}`,
			},
		},
		{
			SubService:   "replication_controllers",
			Struct:       &corev1.ReplicationController{},
			ResourceFunc: v1.ReplicationControllersGetter.ReplicationControllers,
			// MockFieldsValue: map[string]string{
			// 	"":
			// },
		},
		{
			SubService:   "resource_quotas",
			Struct:       &corev1.ResourceQuota{},
			ResourceFunc: v1.ResourceQuotasGetter.ResourceQuotas,
		},
		{
			SubService:   "secrets",
			Struct:       &corev1.Secret{},
			SkipFields:   []string{"Data", "StringData"},
			ResourceFunc: v1.SecretsGetter.Secrets,
		},
		{
			SubService:   "services",
			Struct:       &corev1.Service{},
			ResourceFunc: v1.ServicesGetter.Services,
			SkipFields:   []string{"ClusterIP", "ClusterIPs", "ExternalIPs", "LoadBalancerIP"},
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
			MockFieldsValue: map[string]string{
				"ExternalIPs":    `[]string{"8.8.8.8"}`,
				"ClusterIPs":     `[]string{"8.8.8.8"}`,
				"ClusterIP":      `"8.8.8.8"`,
				"LoadBalancerIP": `"8.8.8.8"`,
			},
		},
		{
			SubService:   "service_accounts",
			Struct:       &corev1.ServiceAccount{},
			ResourceFunc: v1.ServiceAccountsGetter.ServiceAccounts,
		},
	}

	for _, resource := range resources {
		resource.Service = "core"
		resource.ServiceFunc = kubernetes.Interface.CoreV1
		resource.SkipMockTypeFields = []string{"IntOrString"}
		resource.SkipMockFields = []string{"FieldsV1"}
	}

	return resources
}
