package recipes

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func AppsResources() []*Resource {
	resources := []*Resource{
		{
			SubService:   "daemon_sets",
			Struct:       &appsv1.DaemonSet{},
			ResourceFunc: v1.DaemonSetsGetter.DaemonSets,
		},
		{
			SubService:   "deployments",
			Struct:       &appsv1.Deployment{},
			ResourceFunc: v1.DeploymentsGetter.Deployments,
		},
		{
			SubService:   "replica_sets",
			Struct:       &appsv1.ReplicaSet{},
			ResourceFunc: v1.ReplicaSetsGetter.ReplicaSets,
		},
		{
			SubService:   "stateful_sets",
			Struct:       &appsv1.StatefulSet{},
			ResourceFunc: v1.StatefulSetsGetter.StatefulSets,
		},
	}

	for _, resource := range resources {
		resource.Service = "apps"
		resource.ServiceFunc = kubernetes.Interface.AppsV1
		resource.SkipMockTypeFields = []string{"IntOrString", "*IntOrString", "*intstr.IntOrString"}
	}

	return resources
}
