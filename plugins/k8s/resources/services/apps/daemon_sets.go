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

func DaemonSets() *schema.Table {
	return &schema.Table{
		Name:         "k8s_apps_daemon_sets",
		Description:  "DaemonSet represents the configuration of a daemon set.",
		Resolver:     fetchDaemonSets,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		IgnoreError:  client.IgnoreForbiddenNotFound,
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
				Name:          "deletion_grace_period_seconds",
				Description:   "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
				IgnoreInTests: true,
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
				Name:          "owner_references",
				Description:   "List of objects depended by this object",
				Type:          schema.TypeJSON,
				Resolver:      resolveDaemonSetsOwnerReferences,
				IgnoreInTests: true,
			},
			{
				Name:          "finalizers",
				Description:   "Must be empty before the object is deleted from the registry",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("ObjectMeta.Finalizers"),
				IgnoreInTests: true,
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
				Resolver:    resolveDaemonSetsManagedFields,
			},
			{
				Name:        "selector_match_labels",
				Description: "matchLabels is a map of {key,value} pairs",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Selector.MatchLabels"),
			},
			{
				Name:        "template",
				Description: "An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template",
				Type:        schema.TypeJSON,
				Resolver:    resolveDaemonSetsTemplate,
			},
			{
				Name:        "update_strategy_type",
				Description: "Type of daemon set update",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.UpdateStrategy.Type"),
			},
			{
				Name:     "update_strategy_rolling_update_max_unavailable_type",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.MaxUnavailable.Type"),
			},
			{
				Name:     "update_strategy_rolling_update_max_unavailable_int_val",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.MaxUnavailable.IntVal"),
			},
			{
				Name:     "update_strategy_rolling_update_max_unavailable_str_val",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.MaxUnavailable.StrVal"),
			},
			{
				Name:     "update_strategy_rolling_update_max_surge_type",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.MaxSurge.Type"),
			},
			{
				Name:     "update_strategy_rolling_update_max_surge_int_val",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.MaxSurge.IntVal"),
			},
			{
				Name:     "update_strategy_rolling_update_max_surge_str_val",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.UpdateStrategy.RollingUpdate.MaxSurge.StrVal"),
			},
			{
				Name:        "min_ready_seconds",
				Description: "The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.MinReadySeconds"),
			},
			{
				Name:        "revision_history_limit",
				Description: "The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.RevisionHistoryLimit"),
			},
			{
				Name:        "status_current_number_scheduled",
				Description: "The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.CurrentNumberScheduled"),
			},
			{
				Name:        "status_number_misscheduled",
				Description: "The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.NumberMisscheduled"),
			},
			{
				Name:        "status_desired_number_scheduled",
				Description: "The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.DesiredNumberScheduled"),
			},
			{
				Name:        "status_number_ready",
				Description: "The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.NumberReady"),
			},
			{
				Name:        "status_observed_generation",
				Description: "The most recent generation observed by the daemon set controller.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:        "status_updated_number_scheduled",
				Description: "The total number of nodes that are running updated daemon pod",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.UpdatedNumberScheduled"),
			},
			{
				Name:        "status_number_available",
				Description: "The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.NumberAvailable"),
			},
			{
				Name:        "status_number_unavailable",
				Description: "The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.NumberUnavailable"),
			},
			{
				Name:          "status_collision_count",
				Description:   "Count of hash collisions for the DaemonSet",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("Status.CollisionCount"),
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				IgnoreInTests: true,
				Name:          "k8s_apps_daemon_set_selector_match_expressions",
				Description:   "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:      fetchDaemonSetSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "daemon_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_daemon_sets table (FK)",
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
				IgnoreInTests: true,
				Name:          "k8s_apps_daemon_set_status_conditions",
				Description:   "DaemonSetCondition describes the state of a DaemonSet at a certain point.",
				Resolver:      fetchDaemonSetStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "daemon_set_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_daemon_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of DaemonSet condition.",
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

func fetchDaemonSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client).Services().DaemonSets
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
func resolveDaemonSetsOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.DaemonSet)
	if !ok {
		return fmt.Errorf("not a appsv1.DaemonSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveDaemonSetsManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.DaemonSet)
	if !ok {
		return fmt.Errorf("not a appsv1.DaemonSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveDaemonSetsTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.DaemonSet)
	if !ok {
		return fmt.Errorf("not a appsv1.DaemonSet instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchDaemonSetSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(appsv1.DaemonSet)
	if !ok {
		return fmt.Errorf("not a appsv1.DaemonSet instance: %T", parent.Item)
	}
	if p.Spec.Selector == nil {
		return nil
	}
	res <- p.Spec.Selector.MatchExpressions
	return nil
}
func fetchDaemonSetStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(appsv1.DaemonSet)
	if !ok {
		return fmt.Errorf("not a appsv1.DaemonSet instance: %T", parent.Item)
	}

	res <- p.Status.Conditions
	return nil
}
