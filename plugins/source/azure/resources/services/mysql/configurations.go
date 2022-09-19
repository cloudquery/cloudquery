// Auto generated code - DO NOT EDIT.

package mysql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/mysql/mgmt/mysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func configurations() *schema.Table {
	return &schema.Table{
		Name:     "azure_mysql_configurations",
		Resolver: fetchMySQLConfigurations,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "mysql_server_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Value"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "default_value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultValue"),
			},
			{
				Name:     "data_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataType"),
			},
			{
				Name:     "allowed_values",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllowedValues"),
			},
			{
				Name:     "source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Source"),
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

func fetchMySQLConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().MySQL.Configurations

	server := parent.Item.(mysql.Server)
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
