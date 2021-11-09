package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RbacRoleBindings() *schema.Table {
	return &schema.Table{
		Name:        "k8s_rbac_role_bindings",
		Description: "RoleBinding references a role, but does not contain it",
		Resolver:    fetchRbacRoleBindings,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TypeMeta.APIVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:     "generate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.GenerateName"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.SelfLink"),
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:     "resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:     "deletion_grace_period_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:     "owner_references",
				Type:     schema.TypeJSON,
				Resolver: resolveRbacRoleBindingOwnerReferences,
			},
			{
				Name:     "finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ObjectMeta.Finalizers"),
			},
			{
				Name:     "cluster_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:     "managed_fields",
				Type:     schema.TypeJSON,
				Resolver: resolveRbacRoleBindingManagedFields,
			},
			{
				Name:        "role_ref_api_group",
				Description: "APIGroup is the group for the resource being referenced",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleRef.APIGroup"),
			},
			{
				Name:        "role_ref_kind",
				Description: "Kind is the type of resource being referenced",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleRef.Kind"),
			},
			{
				Name:        "role_ref_name",
				Description: "Name is the name of resource being referenced",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleRef.Name"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_rbac_role_binding_subjects",
				Description: "Subject contains a reference to the object or user identities a role binding applies to",
				Resolver:    fetchRbacRoleBindingSubjects,
				Columns: []schema.Column{
					{
						Name:        "role_binding_cq_id",
						Description: "Unique CloudQuery ID of k8s_rbac_role_bindings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Kind of object being referenced",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_group",
						Description: "APIGroup holds the API group of the referenced subject. Defaults to \"\" for ServiceAccount subjects. Defaults to \"rbac.authorization.k8s.io\" for User and Group subjects. +optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("APIGroup"),
					},
					{
						Name:        "name",
						Description: "Name of the object being referenced.",
						Type:        schema.TypeString,
					},
					{
						Name:        "namespace",
						Description: "Namespace of the referenced object",
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
func fetchRbacRoleBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client).Services().RoleBindings
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
func resolveRbacRoleBindingOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(rbacv1.RoleBinding)
	if !ok {
		return fmt.Errorf("not a rbacv1.Role instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveRbacRoleBindingManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(rbacv1.RoleBinding)
	if !ok {
		return fmt.Errorf("not a rbacv1.Role instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchRbacRoleBindingSubjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	role, ok := parent.Item.(rbacv1.RoleBinding)
	if !ok {
		return fmt.Errorf("not a rbacv1.RoleBinding instance: %T", parent.Item)
	}
	res <- role.Subjects
	return nil
}
