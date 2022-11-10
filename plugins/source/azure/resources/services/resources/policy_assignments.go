// Auto generated code - DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PolicyAssignments() *schema.Table {
	return &schema.Table{
		Name:        "azure_resources_policy_assignments",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy#Assignment`,
		Resolver:    fetchResourcesPolicyAssignments,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "policy_definition_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyDefinitionID"),
			},
			{
				Name:     "scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Scope"),
			},
			{
				Name:     "not_scopes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NotScopes"),
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Parameters"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "enforcement_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnforcementMode"),
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
		},
	}
}

func fetchResourcesPolicyAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Resources.PolicyAssignments

	response, err := svc.List(ctx, meta.(*client.Client).SubscriptionId, "", nil)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
