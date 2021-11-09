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

func RbacRoles() *schema.Table {
	return &schema.Table{
		Name:         "k8s_rbac_roles",
		Description:  "Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.",
		Resolver:     fetchRbacRoles,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
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
				Resolver: resolveRbacRoleOwnerReferences,
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
				Resolver: resolveRbacRoleManagedFields,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_rbac_role_rules",
				Description: "PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.",
				Resolver:    fetchRbacRoleRules,
				Columns: []schema.Column{
					{
						Name:        "role_cq_id",
						Description: "Unique CloudQuery ID of k8s_rbac_roles table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "verbs",
						Description: "Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "api_groups",
						Description: "APIGroups is the name of the APIGroup that contains the resources",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("APIGroups"),
					},
					{
						Name:        "resources",
						Description: "Resources is a list of resources this rule applies to",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "resource_names",
						Description: "ResourceNames is an optional white list of names that the rule applies to",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "non_resource_urls",
						Description: "NonResourceURLs is a set of partial urls that a user should have access to",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("NonResourceURLs"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRbacRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client).Services().Roles
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
func resolveRbacRoleOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(rbacv1.Role)
	if !ok {
		return fmt.Errorf("not a rbacv1.Role instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveRbacRoleManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(rbacv1.Role)
	if !ok {
		return fmt.Errorf("not a rbacv1.Role instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchRbacRoleRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	role, ok := parent.Item.(rbacv1.Role)
	if !ok {
		return fmt.Errorf("not a rbacv1.Role instance: %T", parent.Item)
	}
	res <- role.Rules
	return nil
}
