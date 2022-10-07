package recipes

import appsv1 "k8s.io/api/apps/v1"

func AppsResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "daemon_sets",
			Struct:     &appsv1.DaemonSet{},
		},
		{
			SubService: "deployments",
			Struct:     &appsv1.Deployment{},
		},
		{
			SubService: "replica_sets",
			Struct:     &appsv1.ReplicaSet{},
		},
		{
			SubService: "stateful_sets",
			Struct:     &appsv1.StatefulSet{},
		},
	}

	for _, resource := range resources {
		resource.Service = "apps"
	}

	return resources
}
