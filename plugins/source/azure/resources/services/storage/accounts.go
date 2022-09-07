// Auto generated code - DO NOT EDIT.

package storage

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/Azure/go-autorest/autorest"

	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "azure_storage_accounts",
		Resolver:  fetchStorageAccounts,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "extended_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedLocation"),
			},
			{
				Name:     "account_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccountProperties"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
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
				Name:     "blob_logging_settings",
				Type:     schema.TypeJSON,
				Resolver: fetchStorageAccountBlobLoggingSettings,
			},
			{
				Name:     "queue_logging_settings",
				Type:     schema.TypeJSON,
				Resolver: fetchStorageAccountQueueLoggingSettings,
			},
		},

		Relations: []*schema.Table{
			BlobServices(), Containers(),
		},
	}
}

func fetchStorageAccountBlobLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	acc := resource.Item.(storage.Account)

	if !isBlobSupported(&acc) {
		return nil
	}

	// fetch storageClient account keys for Shared Key authentication
	storageClient := meta.(*client.Client).Services().Storage
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	keysResult, err := storageClient.Accounts.ListKeys(ctx, details.ResourceGroup, *acc.Name, "")
	if err != nil {
		if client.IgnoreAccessDenied(err) {
			meta.Logger().Warn("received access denied on Accounts.ListKeys", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
			return nil
		}
		return errors.WithStack(err)
	}
	if keysResult.Keys == nil || len(*keysResult.Keys) == 0 {
		return nil
	}

	// use account key to create a new authorizer and then fetch service properties
	auth, err := autorest.NewSharedKeyAuthorizer(*acc.Name, *(*keysResult.Keys)[0].Value, autorest.SharedKeyLite)
	if err != nil {
		return errors.WithStack(err)
	}
	blobProps := storageClient.NewBlobServiceProperties(auth)
	result, err := blobProps.GetServiceProperties(ctx, *acc.Name)
	if err != nil {
		// For premium 'page blob' storage accounts, we sometimes get "authorization error", not sure why.
		// In any case, we can probably ignore this since it only happens for premium 'page blob' storage accounts.
		if client.IgnoreAccessDenied(err) {
			meta.Logger().Warn("received access denied on GetServiceProperties", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
			return nil
		}
		return errors.WithStack(err)
	}
	var logging *accounts.Logging
	if result.StorageServiceProperties != nil {
		logging = result.StorageServiceProperties.Logging
	}
	data, err := json.Marshal(logging)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, data))
}

func fetchStorageAccountQueueLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	acc := resource.Item.(storage.Account)
	if !isQueueSupported(&acc) {
		return nil
	}

	// fetch storage account keys for Shared Key authentication
	storageClient := meta.(*client.Client).Services().Storage
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	keysResult, err := storageClient.Accounts.ListKeys(ctx, details.ResourceGroup, *acc.Name, "")
	if err != nil {
		if client.IgnoreAccessDenied(err) {
			meta.Logger().Warn("received access denied on Accounts.ListKeys", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
			return nil
		}
	}
	if keysResult.Keys == nil || len(*keysResult.Keys) == 0 {
		return nil
	}

	// use account key to create a new authorizer and then fetch service properties
	auth, err := autorest.NewSharedKeyAuthorizer(*acc.Name, *(*keysResult.Keys)[0].Value, autorest.SharedKeyLite)
	if err != nil {
		return errors.WithStack(err)
	}
	blobProps := storageClient.NewQueueServiceProperties(auth)
	result, err := blobProps.GetServiceProperties(ctx, *acc.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	data, err := json.Marshal(result.Logging)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, data))
}

// isQueueSupported checks whether queues are supported for a storage account.
// Premium storage accounts don't support queues.
func isQueueSupported(account *storage.Account) bool {
	return account.Sku.Tier == storage.Standard && account.Kind == storage.StorageV2
}

// isBlobSupported checks whether blobs are supported for a storage account.
func isBlobSupported(account *storage.Account) bool {
	return (account.Kind == storage.Storage) || (account.Kind == storage.StorageV2) ||
		(account.Kind == storage.BlockBlobStorage) || (account.Kind == storage.BlobStorage)
}

func fetchStorageAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Storage.Accounts

	response, err := svc.List(ctx)

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
