// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func virtualNetworkRules() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_virtual_network_rules",
		Resolver: fetchSQLVirtualNetworkRules,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sql_server_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "virtual_network_subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualNetworkSubnetID"),
			},
			{
				Name:     "ignore_missing_vnet_service_endpoint",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IgnoreMissingVnetServiceEndpoint"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
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

func fetchSQLVirtualNetworkRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.VirtualNetworkRules

	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)

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
