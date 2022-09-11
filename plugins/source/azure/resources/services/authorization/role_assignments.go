// Auto generated code - DO NOT EDIT.

package authorization

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func RoleAssignments() *schema.Table {
	return &schema.Table{
		Name:      "azure_authorization_role_assignments",
		Resolver:  fetchAuthorizationRoleAssignments,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Scope"),
			},
			{
				Name:     "role_definition_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleDefinitionID"),
			},
			{
				Name:     "principal_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrincipalID"),
			},
		},
	}
}

func fetchAuthorizationRoleAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Authorization.RoleAssignments

	response, err := svc.List(ctx, "")

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
