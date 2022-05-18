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

func Deployments() *schema.Table {
	return &schema.Table{
		Name:         "k8s_apps_deployments",
		Description:  "Deployment enables declarative updates for Pods and ReplicaSets.",
		Resolver:     fetchAppsDeployments,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		IgnoreError:  client.IgnoreForbidden,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.APIVersion"),
			},
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
				Resolver:    resolveAppsDeploymentsOwnerReferences,
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
				Resolver:    resolveAppsDeploymentsManagedFields,
			},
			{
				Name:        "replicas",
				Description: "Number of desired pods",
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
				Description: "Template describes the pods that will be created.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAppsDeploymentsTemplate,
			},
			{
				Name:        "strategy_type",
				Description: "Type of deployment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Strategy.Type"),
			},
			{
				Name:     "strategy_rolling_update_max_unavailable_type",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxUnavailable.Type"),
			},
			{
				Name:     "strategy_rolling_update_max_unavailable_int_val",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxUnavailable.IntVal"),
			},
			{
				Name:     "strategy_rolling_update_max_unavailable_str_val",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal"),
			},
			{
				Name:     "strategy_rolling_update_max_surge_type",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxSurge.Type"),
			},
			{
				Name:     "strategy_rolling_update_max_surge_int_val",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxSurge.IntVal"),
			},
			{
				Name:     "strategy_rolling_update_max_surge_str_val",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Strategy.RollingUpdate.MaxSurge.StrVal"),
			},
			{
				Name:        "min_ready_seconds",
				Description: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.MinReadySeconds"),
			},
			{
				Name:        "revision_history_limit",
				Description: "The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.RevisionHistoryLimit"),
			},
			{
				Name:        "paused",
				Description: "Indicates that the deployment is paused.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Paused"),
			},
			{
				Name:        "progress_deadline_seconds",
				Description: "The maximum time in seconds for a deployment to make progress before it is considered to be failed",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.ProgressDeadlineSeconds"),
			},
			{
				Name:        "status_observed_generation",
				Description: "The generation observed by the deployment controller.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:        "status_replicas",
				Description: "Total number of non-terminated pods targeted by this deployment (their labels match the selector).",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Replicas"),
			},
			{
				Name:        "status_updated_replicas",
				Description: "Total number of non-terminated pods targeted by this deployment that have the desired template spec.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.UpdatedReplicas"),
			},
			{
				Name:        "status_ready_replicas",
				Description: "Total number of ready pods targeted by this deployment.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.ReadyReplicas"),
			},
			{
				Name:        "status_available_replicas",
				Description: "Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.AvailableReplicas"),
			},
			{
				Name:        "status_unavailable_replicas",
				Description: "Total number of unavailable pods targeted by this deployment",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.UnavailableReplicas"),
			},
			{
				Name:        "status_collision_count",
				Description: "Count of hash collisions for the Deployment",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.CollisionCount"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_apps_deployment_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchAppsDeploymentSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
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
				Name:        "k8s_apps_deployment_status_conditions",
				Description: "DeploymentCondition describes the state of a deployment at a certain point.",
				Resolver:    fetchAppsDeploymentStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "deployment_cq_id",
						Description: "Unique CloudQuery ID of k8s_apps_deployments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of deployment condition.",
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

func fetchAppsDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client).Services().Deployments
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
func resolveAppsDeploymentsOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentsManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveAppsDeploymentsTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.Spec.Template)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchAppsDeploymentSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	deployment, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	if deployment.Spec.Selector == nil {
		return nil
	}
	res <- deployment.Spec.Selector.MatchExpressions
	return nil
}
func fetchAppsDeploymentStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	deployment, ok := parent.Item.(appsv1.Deployment)
	if !ok {
		return fmt.Errorf("not a appsv1.Deployment instance: %T", parent.Item)
	}
	res <- deployment.Status.Conditions
	return nil
}
