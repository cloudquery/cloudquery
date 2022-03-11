package authorization

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AuthorizationRoleDefinitions() *schema.Table {
	return &schema.Table{
		Name:         "azure_authorization_role_definitions",
		Description:  "RoleDefinition role definition",
		Resolver:     fetchAuthorizationRoleDefinitions,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "id",
				Description: "The role definition ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The role definition name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The role definition type",
				Type:        schema.TypeString,
			},
			{
				Name:        "role_name",
				Description: "The role name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleDefinitionProperties.RoleName"),
			},
			{
				Name:        "description",
				Description: "The role definition description",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleDefinitionProperties.Description"),
			},
			{
				Name:        "role_type",
				Description: "The role type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleDefinitionProperties.RoleType"),
			},
			{
				Name:        "assignable_scopes",
				Description: "Role definition assignable scopes",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("RoleDefinitionProperties.AssignableScopes"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_authorization_role_definition_permissions",
				Description: "Permission role definition permissions",
				Resolver:    fetchAuthorizationRoleDefinitionPermissions,
				Columns: []schema.Column{
					{
						Name:        "role_definition_cq_id",
						Description: "Unique CloudQuery ID of azure_authorization_role_definitions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "actions",
						Description: "Allowed actions",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "not_actions",
						Description: "Denied actions",
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
func fetchAuthorizationRoleDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Authorization.RoleDefinitions
	result, err := svc.List(ctx, client.ScopeSubscription(cl.SubscriptionId), "")
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchAuthorizationRoleDefinitionPermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	def, ok := parent.Item.(authorization.RoleDefinition)
	if !ok {
		return fmt.Errorf("not an authorization.RoleDefinition instance: %T", parent.Item)
	}
	if def.Permissions == nil {
		return nil
	}
	res <- *def.Permissions
	return nil
}
