// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func managedInstanceEncryptionProtectors() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_managed_instance_encryption_protectors",
		Resolver: fetchSQLManagedInstanceEncryptionProtectors,
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

func fetchSQLManagedInstanceEncryptionProtectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedInstanceEncryptionProtectors

	instance := parent.Item.(sql.ManagedInstance)
	resourceDetails, err := client.ParseResourceID(*instance.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.ListByInstance(ctx, resourceDetails.ResourceGroup, *instance.Name)

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
