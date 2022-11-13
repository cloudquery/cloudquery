package recipes

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func Apps() []*Resource {
	resources := []*Resource{
		{
			SubService:   "daemon_sets",
			Struct:       &appsv1.DaemonSet{},
			ResourceFunc: v1.DaemonSetsGetter.DaemonSets,
			MockImports:  []string{`corev1 "k8s.io/api/core/v1"`},
			FakerOverride: `
			r.Spec.Template = corev1.PodTemplateSpec{}
			r.Spec.UpdateStrategy = resource.DaemonSetUpdateStrategy{}
			`,
		},
		{
			SubService:   "deployments",
			Struct:       &appsv1.Deployment{},
			ResourceFunc: v1.DeploymentsGetter.Deployments,
			MockImports:  []string{`corev1 "k8s.io/api/core/v1"`},
			FakerOverride: `
			r.Spec.Template = corev1.PodTemplateSpec{}
			r.Spec.Strategy = resource.DeploymentStrategy{}
			`,
		},
		{
			SubService:   "replica_sets",
			Struct:       &appsv1.ReplicaSet{},
			ResourceFunc: v1.ReplicaSetsGetter.ReplicaSets,
			MockImports:  []string{`corev1 "k8s.io/api/core/v1"`},
			FakerOverride: `
			r.Spec.Template = corev1.PodTemplateSpec{}
			`,
		},
		{
			SubService:   "stateful_sets",
			Struct:       &appsv1.StatefulSet{},
			ResourceFunc: v1.StatefulSetsGetter.StatefulSets,
			MockImports:  []string{`corev1 "k8s.io/api/core/v1"`},
			FakerOverride: `
			r.Spec.Template = corev1.PodTemplateSpec{}
			r.Spec.UpdateStrategy = resource.StatefulSetUpdateStrategy{}
			r.Spec.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{}
			`,
		},
	}

	for _, resource := range resources {
		resource.Service = "apps"
		resource.ServiceFunc = kubernetes.Interface.AppsV1
	}

	return resources
}
