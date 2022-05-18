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

func StatefulSets() *schema.Table {
	return &schema.Table{
		Name:         "k8s_apps_stateful_sets",
		Description:  "StatefulSet represents a set of pods with consistent identities. Identities are defined as:  - Network: A single stable DNS and hostname.  - Storage: As many VolumeClaims as requested. The StatefulSet guarantees that a given network identity will always map to the same storage identity.",
		Resolver:     fetchAppsStatefulSets,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		IgnoreError:  client.IgnoreForbidden,
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
				Resolver:    resolveAppsStatefulSetsOwnerReferences,
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
				Resolver:    resolveAppsStatefulSetsManagedFields,
			},
			{
				Name:        "replicas",
				Description: "replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. TODO: Consider a rename of this field.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Replicas"),
			},
			{
				Name:        "selector_match_labels",
				Description: "matchLabels is a map of {key,value} pairs",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Selector.MatchLabels"),
			},
			{
				Name:        "template",
				Description: "template is the object that describes the pod that will be created if insufficient replicas are detected",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsStatefulSetsTemplate,
			},
			{
				Name:        "volume_claim_templates",
				Description: "volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsStatefulSetsVolumeClaimTemplates,
			},
			{
				Name:        "service_name",
				Description: "serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.ServiceName"),
			},
			{
				Name:        "pod_management_policy",
				Description: "podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.PodManagementPolicy"),
			},
			{
				Name:        "update_strategy_type",
				Description: "Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.UpdateStrategy.Type"),
			},
			{
				Name:        "update_strategy_rolling_update_partition",
				Description: "Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.Partition"),
			},
			{
				Name:        "revision_history_limit",
				Description: "revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.RevisionHistoryLimit"),
			},
			{
				Name:        "min_ready_seconds",
				Description: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.MinReadySeconds"),
			},
			{
				Name:        "status_observed_generation",
				Description: "observedGeneration is the most recent generation observed for this StatefulSet",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:        "status_replicas",
				Description: "replicas is the number of Pods created by the StatefulSet controller.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Replicas"),
			},
			{
				Name:        "status_ready_replicas",
				Description: "readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.ReadyReplicas"),
			},
			{
				Name:        "status_current_replicas",
				Description: "currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.CurrentReplicas"),
			},
			{
				Name:        "status_updated_replicas",
				Description: "updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.UpdatedReplicas"),
			},
			{
				Name:        "status_current_revision",
				Description: "currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.CurrentRevision"),
			},
			{
				Name:        "status_update_revision",
				Description: "updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.UpdateRevision"),
			},
			{
				Name:        "status_collision_count",
				Description: "collisionCount is the count of hash collisions for the StatefulSet",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.CollisionCount"),
			},
			{
				Name:        "status_available_replicas",
				Description: "Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset. This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate. Remove omitempty when graduating to beta",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.AvailableReplicas"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_apps_stateful_set_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchAppsStatefulSetSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
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
				Name:        "k8s_apps_stateful_set_status_conditions",
				Description: "StatefulSetCondition describes the state of a statefulset at a certain point.",
				Resolver:    fetchAppsStatefulSetStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "stateful_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of statefulset condition.",
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

func fetchAppsStatefulSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client).Services().StatefulSets
	opts := metav1.ListOptions{}
	for {
		result, err := cl.List(ctx, opts)
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
func resolveAppsStatefulSetsOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsStatefulSetsManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsStatefulSetsTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsStatefulSetsVolumeClaimTemplates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.VolumeClaimTemplates)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchAppsStatefulSetSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}
	if p.Spec.Selector == nil {
		return nil
	}
	res <- p.Spec.Selector.MatchExpressions
	return nil
}
func fetchAppsStatefulSetStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(appsv1.StatefulSet)
	if !ok {
		return fmt.Errorf("not a appsv1.StatefulSet instance: %T", parent.Item)
	}

	res <- p.Status.Conditions
	return nil
}
