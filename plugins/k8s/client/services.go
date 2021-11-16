package client

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Services struct {
	Client          *kubernetes.Clientset
	Nodes           NodesClient
	Pods            PodsClient
	Services        ServicesClient
	DaemonSets      DaemonSetsClient
	StatefulSets    StatefulSetsClient
	Deployments     DeploymentsClient
	Namespaces      NamespacesClient
	ReplicaSets     ReplicaSetsClient
	Jobs            JobsClient
	Roles           RolesClient
	RoleBindings    RoleBindingsClient
	NetworkPolicies NetworkPoliciesClient
	CronJobs        CronJobsClient
}

//go:generate mockgen -package=mocks -destination=./mocks/namespaces.go . NamespacesClient
type NamespacesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/nodes.go . NodesClient
type NodesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.NodeList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/pods.go . PodsClient
type PodsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/services.go . ServicesClient
type ServicesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/deployments.go . DeploymentsClient
type DeploymentsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.DeploymentList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/stateful_sets.go . StatefulSetsClient
type StatefulSetsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.StatefulSetList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/network_policies.go . NetworkPoliciesClient
type NetworkPoliciesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.NetworkPolicyList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/replica_sets.go . ReplicaSetsClient
type ReplicaSetsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.ReplicaSetList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/jobs.go . JobsClient
type JobsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*batchv1.JobList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/cronjobs.go . CronJobsClient
type CronJobsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*batchv1.CronJobList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/daemon_sets.go . DaemonSetsClient
type DaemonSetsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.DaemonSetList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/roles.go . RolesClient
type RolesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleList, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/role_bindings.go . RoleBindingsClient
type RoleBindingsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleBindingList, error)
}
