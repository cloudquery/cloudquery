// Auto generated code - DO NOT EDIT.

package datalake

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
)

func StoreAccounts() *schema.Table {
	return &schema.Table{
		Name:                "azure_datalake_store_accounts",
		Resolver:            fetchDataLakeStoreAccounts,
		PreResourceResolver: getDataLakeStoreAccount,
		Multiplex:           client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "data_lake_store_account_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DataLakeStoreAccountProperties"),
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
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchDataLakeStoreAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().DataLake.StoreAccounts

	response, err := svc.List(ctx, "", nil, nil, "", "", nil)

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

func getDataLakeStoreAccount(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	svc := meta.(*client.Client).Services().DataLake.StoreAccounts

	account := r.Item.(account.DataLakeStoreAccount)
	resourceDetails, err := client.ParseResourceID(*account.ID)
	if err != nil {
		errors.WithStack(err)
	}
	item, err := svc.Get(ctx, resourceDetails.ResourceGroup, *account.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	r.SetItem(item)
	return nil
}
