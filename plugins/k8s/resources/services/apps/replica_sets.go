package apps

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ReplicaSets() *schema.Table {
	return &schema.Table{
		Name:         "k8s_apps_replica_sets",
		Description:  "ReplicaSet ensures that a specified number of pod replicas are running at any given time.",
		Resolver:     fetchAppsReplicaSets,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "name",
				Description: "Name must be unique within a namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:        "generate_name",
				Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.GenerateName"),
			},
			{
				Name:        "namespace",
				Description: "Namespace defines the space within which each name must be unique",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:        "self_link",
				Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.SelfLink"),
			},
			{
				Name:        "uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:        "resource_version",
				Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:        "generation",
				Description: "A sequence number representing a specific generation of the desired state. Populated by the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:        "deletion_grace_period_seconds",
				Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:        "labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:        "annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:        "owner_references",
				Description: "List of objects depended by this object",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsReplicaSetsOwnerReferences,
			},
			{
				Name:        "finalizers",
				Description: "Must be empty before the object is deleted from the registry",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "managed_fields",
				Description: "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsReplicaSetsManagedFields,
			},
			{
				Name:        "replicas",
				Description: "Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Replicas"),
			},
			{
				Name:        "min_ready_seconds",
				Description: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.MinReadySeconds"),
			},
			{
				Name:        "selector_match_labels",
				Description: "matchLabels is a map of {key,value} pairs",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Selector.MatchLabels"),
			},
			{
				Name:        "template",
				Description: "Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsReplicaSetsTemplate,
			},
			{
				Name:        "status_replicas",
				Description: "Replicas is the most recently observed number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Replicas"),
			},
			{
				Name:        "status_fully_labeled_replicas",
				Description: "The number of pods that have labels matching the labels of the pod template of the replicaset.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.FullyLabeledReplicas"),
			},
			{
				Name:        "status_ready_replicas",
				Description: "The number of ready replicas for this replica set.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.ReadyReplicas"),
			},
			{
				Name:        "status_available_replicas",
				Description: "The number of available replicas (ready for at least minReadySeconds) for this replica set.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.AvailableReplicas"),
			},
			{
				Name:        "status_observed_generation",
				Description: "ObservedGeneration reflects the generation of the most recently observed ReplicaSet.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_apps_replica_set_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchAppsReplicaSetSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "replica_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_replica_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key",
						Description: "key is the label key that the selector applies to. +patchMergeKey=key +patchStrategy=merge",
						Type:        schema.TypeString,
					},
					{
						Name:        "operator",
						Description: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "values is an array of string values",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "k8s_apps_replica_set_status_conditions",
				Description: "ReplicaSetCondition describes the state of a replica set at a certain point.",
				Resolver:    fetchAppsReplicaSetStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "replica_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_replica_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of replica set condition.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Status of the condition, one of True, False, Unknown.",
						Type:        schema.TypeString,
					},
					{
						Name:        "reason",
						Description: "The reason for the condition's last transition.",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "A human readable message indicating details about the transition.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAppsReplicaSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client).Services().ReplicaSets
	opts := metav1.ListOptions{}
	for {
		result, err := client.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}
func resolveAppsReplicaSetsOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.ReplicaSet)
	if !ok {
		return fmt.Errorf("not a appsv1.ReplicaSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsReplicaSetsManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.ReplicaSet)
	if !ok {
		return fmt.Errorf("not a appsv1.ReplicaSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsReplicaSetsTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.ReplicaSet)
	if !ok {
		return fmt.Errorf("not a appsv1.ReplicaSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchAppsReplicaSetSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(appsv1.ReplicaSet)
	if !ok {
		return fmt.Errorf("not a appsv1.ReplicaSet instance: %T", parent.Item)
	}
	if p.Spec.Selector == nil {
		return nil
	}
	res <- p.Spec.Selector.MatchExpressions
	return nil
}
func fetchAppsReplicaSetStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(appsv1.ReplicaSet)
	if !ok {
		return fmt.Errorf("not a appsv1.ReplicaSet instance: %T", parent.Item)
	}

	res <- p.Status.Conditions
	return nil
}
