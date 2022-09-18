// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/sql/mgmt/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func firewallRules() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_firewall_rules",
		Resolver: fetchSQLFirewallRules,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "start_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StartIPAddress"),
			},
			{
				Name:     "end_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndIPAddress"),
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
		},
	}
}

func fetchSQLFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.FirewallRules

	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
