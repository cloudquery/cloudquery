package core

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ResourceQuotas() *schema.Table {
	return &schema.Table{
		Name:         "k8s_core_resource_quotas",
		Description:  "ResourceQuota sets aggregate quota restrictions enforced per namespace",
		Resolver:     fetchCoreResourceQuotas,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		IgnoreError:  client.IgnoreForbiddenNotFound,
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
				Name:          "deletion_grace_period_seconds",
				Description:   "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
				IgnoreInTests: true,
			},
			{
				Name:          "labels",
				Description:   "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ObjectMeta.Labels"),
				IgnoreInTests: true,
			},
			{
				Name:          "annotations",
				Description:   "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ObjectMeta.Annotations"),
				IgnoreInTests: true,
			},
			{
				Name:          "owner_references",
				Description:   "List of objects depended by this object",
				Type:          schema.TypeJSON,
				Resolver:      resolveCoreResourceQuotasOwnerReferences,
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
				Name:          "managed_fields",
				Description:   "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow",
				Type:          schema.TypeJSON,
				Resolver:      resolveCoreResourceQuotasManagedFields,
				IgnoreInTests: true,
			},
			{
				Name:        "hard",
				Description: "hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Hard"),
			},
			{
				Name:        "scopes",
				Description: "A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Spec.Scopes"),
			},
			{
				Name:        "status_hard",
				Description: "Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Status.Hard"),
			},
			{
				Name:        "status_used",
				Description: "Used is the current observed total usage of the resource in the namespace.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Status.Used"),
			},
		},
		Relations: []*schema.Table{
			{
				IgnoreInTests: true,
				Name:          "k8s_core_resource_quota_scope_selector_match_expressions",
				Description:   "A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates the scope name and values.",
				Resolver:      fetchCoreResourceQuotaScopeSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "resource_quota_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_resource_quotas table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "scope_name",
						Description: "The name of the scope that the selector applies to.",
						Type:        schema.TypeString,
					},
					{
						Name:        "operator",
						Description: "Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "An array of string values",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCoreResourceQuotas(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client).Services().ResourceQuotas
	opts := metav1.ListOptions{}
	for {
		result, err := c.List(ctx, opts)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}

func resolveCoreResourceQuotasOwnerReferences(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(corev1.ResourceQuota)
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func resolveCoreResourceQuotasManagedFields(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(corev1.ResourceQuota)
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func fetchCoreResourceQuotaScopeSelectorMatchExpressions(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	resourceQuota := parent.Item.(corev1.ResourceQuota)
	if resourceQuota.Spec.ScopeSelector == nil {
		return nil
	}
	res <- resourceQuota.Spec.ScopeSelector.MatchExpressions
	return nil
}
