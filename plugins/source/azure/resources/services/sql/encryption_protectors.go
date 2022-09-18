// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func encryptionProtectors() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_encryption_protectors",
		Resolver: fetchSQLEncryptionProtectors,
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
				Name:     "subregion",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subregion"),
			},
			{
				Name:     "server_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerKeyName"),
			},
			{
				Name:     "server_key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerKeyType"),
			},
			{
				Name:     "uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URI"),
			},
			{
				Name:     "thumbprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumbprint"),
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

func fetchSQLEncryptionProtectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.EncryptionProtectors

	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.Get(ctx, resourceDetails.ResourceGroup, *server.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
}
