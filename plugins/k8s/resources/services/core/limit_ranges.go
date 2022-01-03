package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func LimitRanges() *schema.Table {
	return &schema.Table{
		Name:         "k8s_core_limit_ranges",
		Description:  "LimitRange sets resource usage limits for each kind of resource in a Namespace.",
		Resolver:     fetchCoreLimitRanges,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
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
				Resolver:      resolveCoreLimitRangesOwnerReferences,
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
				Resolver:    resolveCoreLimitRangesManagedFields,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_core_limit_range_limits",
				Description: "LimitRangeItem defines a min/max usage limit for any resource that matches on kind.",
				Resolver:    fetchCoreLimitRangeLimits,
				Columns: []schema.Column{
					{
						Name:        "limit_range_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_limit_ranges table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of resource that this limit applies to.",
						Type:        schema.TypeString,
					},
					{
						Name:        "max",
						Description: "Max usage constraints on this kind by resource name.",
						Type:        schema.TypeJSON,
					},
					{
						Name:          "min",
						Description:   "Min usage constraints on this kind by resource name.",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:          "default",
						Description:   "Default resource requirement limit value by resource name if resource limit is omitted.",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:          "default_request",
						Description:   "DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:          "max_limit_request_ratio",
						Description:   "MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCoreLimitRanges(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client).Services().LimitRanges
	opts := metav1.ListOptions{}
	for {
		result, err := c.List(ctx, opts)
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

func resolveCoreLimitRangesOwnerReferences(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.LimitRange)
	if !ok {
		return fmt.Errorf("not a corev1.LimitRange instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreLimitRangesManagedFields(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.LimitRange)
	if !ok {
		return fmt.Errorf("not a corev1.LimitRange instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchCoreLimitRangeLimits(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	limitRange, ok := parent.Item.(corev1.LimitRange)
	if !ok {
		return fmt.Errorf("not a corev1.LimitRange instance: %T", parent.Item)
	}
	res <- limitRange.Spec.Limits
	return nil
}
