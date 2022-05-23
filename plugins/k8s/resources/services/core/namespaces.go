package core

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:         "k8s_core_namespaces",
		Description:  "Namespace provides a scope for Names. Use of multiple namespaces is optional.",
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Resolver:     fetchCoreNamespaces,
		IgnoreError:  client.IgnoreForbiddenNotFound,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources ",
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
				Resolver:      resolveCoreNamespacesOwnerReferences,
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
				Resolver:    resolveCoreNamespacesManagedFields,
			},
			{
				Name:        "spec_finalizers",
				Description: "Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/ ",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Spec.Finalizers"),
			},
			{
				Name:        "phase",
				Description: "Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/ ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.Phase"),
			},
			{
				Name:          "conditions",
				Description:   "Represents the latest available observations of a namespace's current state.  +patchMergeKey=type +patchStrategy=merge",
				Type:          schema.TypeJSON,
				Resolver:      resolveCoreNamespacesConditions,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				IgnoreInTests: true,
				Name:          "k8s_meta_owner_references",
				Description:   "OwnerReference contains enough information to let you identify an owning object",
				Resolver:      client.OwnerReferenceResolver,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"resource_uid", "owner_uid"}},
				Columns: []schema.Column{
					{
						Name:        "namespace_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_namespaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_uid",
						Description: "resources this owner object references",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceUID"),
					},
					{
						Name:        "api_version",
						Description: "API version of the referent.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("APIVersion"),
					},
					{
						Name:        "kind",
						Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
						Type:        schema.TypeString,
					},
					{
						Name:        "owner_uid",
						Description: "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UID"),
					},
					{
						Name:        "controller",
						Description: "If true, this reference points to the managing controller.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "block_owner_deletion",
						Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCoreNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	namespaces := meta.(*client.Client).Services().Namespaces
	opts := metav1.ListOptions{}
	for {
		result, err := namespaces.List(ctx, opts)
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
func resolveCoreNamespacesOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	n := resource.Item.(v1.Namespace)
	data, err := json.Marshal(n.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveCoreNamespacesManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	n := resource.Item.(v1.Namespace)
	data, err := json.Marshal(n.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveCoreNamespacesConditions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	n := resource.Item.(v1.Namespace)
	data, err := json.Marshal(n.Status.Conditions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
