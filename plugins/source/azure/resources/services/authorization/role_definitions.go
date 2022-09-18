// Auto generated code - DO NOT EDIT.

package authorization

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func RoleDefinitions() *schema.Table {
	return &schema.Table{
		Name:      "azure_authorization_role_definitions",
		Resolver:  fetchAuthorizationRoleDefinitions,
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
				Name:     "role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleName"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "role_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleType"),
			},
			{
				Name:     "permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Permissions"),
			},
			{
				Name:     "assignable_scopes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AssignableScopes"),
			},
		},
	}
}

func fetchAuthorizationRoleDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Authorization.RoleDefinitions

	response, err := svc.List(ctx, client.ScopeSubscription(meta.(*client.Client).SubscriptionId), "")

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
