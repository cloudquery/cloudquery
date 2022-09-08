package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

const fetchStorageAccountBlobLoggingSettings = `func fetchStorageAccountBlobLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
			meta.Logger().Warn().Msgf("received access denied on Accounts.ListKeys %s %s %s %s %s %s", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
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
			meta.Logger().Warn().Msgf("received access denied on GetServiceProperties %s %s %s %s %s %s", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
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
}`

const fetchStorageAccountQueueLoggingSettings = `func fetchStorageAccountQueueLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
			meta.Logger().Warn().Msgf("received access denied on Accounts.ListKeys %s %s %s %s %s %s", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
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
}`

const isQueueSupported = `// isQueueSupported checks whether queues are supported for a storage account.
// Premium storage accounts don't support queues.
func isQueueSupported(account *storage.Account) bool {
	return account.Sku.Tier == storage.Standard && account.Kind == storage.StorageV2
}`

const isBlobSupported = `// isBlobSupported checks whether blobs are supported for a storage account.
func isBlobSupported(account *storage.Account) bool {
	return (account.Kind == storage.Storage) || (account.Kind == storage.StorageV2) ||
		(account.Kind == storage.BlockBlobStorage) || (account.Kind == storage.BlobStorage)
}`

func Storage() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage", "github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &storage.Account{},
					listFunction: "List",
					helpers:      []string{fetchStorageAccountBlobLoggingSettings, fetchStorageAccountQueueLoggingSettings, isQueueSupported, isBlobSupported},
					customColumns: codegen.ColumnDefinitions{
						codegen.ColumnDefinition{Name: "blob_logging_settings", Type: schema.TypeJSON, Resolver: "fetchStorageAccountBlobLoggingSettings"},
						codegen.ColumnDefinition{Name: "queue_logging_settings", Type: schema.TypeJSON, Resolver: "fetchStorageAccountQueueLoggingSettings"},
					},
					relations: []string{"blobServices(),containers()"},
				},
				{
					azureStruct:  &storage.BlobServiceProperties{},
					listFunction: "List",
					listFunctionArgsInit: []string{`account := parent.Item.(storage.Account)
					if !isBlobSupported(&account) {
						return nil
					}
				
					resource, err := client.ParseResourceID(*account.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listFunctionArgs:   []string{"resource.ResourceGroup", "*account.Name"},
					listHandler:        valueHandler,
					subServiceOverride: "BlobServices",
					isRelation:         true,
				},
				{
					azureStruct:  &storage.ListContainerItem{},
					listFunction: "List",
					listFunctionArgsInit: []string{`account := parent.Item.(storage.Account)
					if !isBlobSupported(&account) {
						return nil
					}
				
					resource, err := client.ParseResourceID(*account.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listFunctionArgs:   []string{"resource.ResourceGroup", "*account.Name", `""`, `""`, `""`},
					subServiceOverride: "Containers",
					isRelation:         true,
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
