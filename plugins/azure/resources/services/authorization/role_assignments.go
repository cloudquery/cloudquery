package authorization

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AuthorizationRoleAssignments() *schema.Table {
	return &schema.Table{
		Name:         "azure_authorization_role_assignments",
		Description:  "RoleAssignment role Assignments",
		Resolver:     fetchAuthorizationRoleAssignments,
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
				Description: "The role assignment ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The role assignment name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The role assignment type",
				Type:        schema.TypeString,
			},
			{
				Name:        "scope",
				Description: "The role assignment scope",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Scope"),
			},
			{
				Name:        "role_definition_id",
				Description: "The role definition ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.RoleDefinitionID"),
			},
			{
				Name:        "principal_id",
				Description: "The principal ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.PrincipalID"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAuthorizationRoleAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Authorization
	result, err := svc.RoleAssignments.List(ctx, "")
	if err != nil {
		return err
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
