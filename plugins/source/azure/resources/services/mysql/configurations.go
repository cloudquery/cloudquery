// Auto generated code - DO NOT EDIT.

package mysql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
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
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
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
