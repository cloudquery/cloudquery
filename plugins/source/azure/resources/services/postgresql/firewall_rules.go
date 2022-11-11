// Auto generated code - DO NOT EDIT.

package postgresql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/postgresql/mgmt/postgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func firewallRules() *schema.Table {
	return &schema.Table{
		Name:        "azure_postgresql_firewall_rules",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql#FirewallRule`,
		Resolver:    fetchPostgreSQLFirewallRules,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "postgresql_server_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
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

func fetchPostgreSQLFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().PostgreSQL.FirewallRules

	server := parent.Item.(postgresql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
