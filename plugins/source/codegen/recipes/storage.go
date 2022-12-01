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

	storageClient := meta.(*client.Client).Services().Storage.Accounts
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return err
	}
	result, err := storageClient.GetBlobServiceProperties(ctx, details.ResourceGroup, *acc.Name)
	if err != nil {
		return err
	}
	if result.StorageServiceProperties != nil {
		return resource.Set(c.Name, result.StorageServiceProperties.Logging)
	}
	return nil
}`

const fetchStorageAccountQueueLoggingSettings = `func fetchStorageAccountQueueLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	acc := resource.Item.(storage.Account)
	if !isQueueSupported(&acc) {
		return nil
	}

	storageClient := meta.(*client.Client).Services().Storage.Accounts
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return err
	}
	result, err := storageClient.GetQueueServiceProperties(ctx, details.ResourceGroup, *acc.Name)
	if err != nil {
		return err
	}
	return  resource.Set(c.Name, result.Logging)
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
	var listContainerResource = resourceDefinition{
		azureStruct:  &storage.ListContainerItem{},
		listFunction: "List",
		listFunctionArgsInit: []string{`account := parent.Item.(storage.Account)
		if !isBlobSupported(&account) {
			return nil
		}
	
		resource, err := client.ParseResourceID(*account.ID)
		if err != nil {
			return err
		}`},
		listFunctionArgs:         []string{"resource.ResourceGroup", "*account.Name", `""`, `""`, `""`},
		subServiceOverride:       "Containers",
		mockListFunctionArgsInit: []string{""},
		mockListFunctionArgs:     []string{`"test"`, `"test"`, `""`, `""`, `gomock.Any()`},
		mockListResult:           "ListContainerItems",
	}
	var blobPropertiesResource = resourceDefinition{
		azureStruct:  &storage.BlobServiceProperties{},
		listFunction: "List",
		listFunctionArgsInit: []string{`account := parent.Item.(storage.Account)
		if !isBlobSupported(&account) {
			return nil
		}
	
		resource, err := client.ParseResourceID(*account.ID)
		if err != nil {
			return err
		}`},
		listFunctionArgs:         []string{"resource.ResourceGroup", "*account.Name"},
		listHandler:              valueHandler,
		subServiceOverride:       "BlobServices",
		mockListFunctionArgsInit: []string{""},
		mockListFunctionArgs:     []string{`"test"`, `"test"`},
		mockListResult:           "BlobServiceItems",
	}

	var accountRelations = []resourceDefinition{listContainerResource, blobPropertiesResource}

	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage",
						"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts",
						"github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues",
					},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage",
						"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts",
						"github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues",
					},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &storage.Account{},
					listFunction: "List",
					helpers:      []string{fetchStorageAccountBlobLoggingSettings, fetchStorageAccountQueueLoggingSettings, isQueueSupported, isBlobSupported},
					skipFields:   []string{"EnableNfsV3"},
					customColumns: codegen.ColumnDefinitions{
						codegen.ColumnDefinition{Name: "blob_logging_settings", Type: schema.TypeJSON, Resolver: "fetchStorageAccountBlobLoggingSettings"},
						codegen.ColumnDefinition{Name: "queue_logging_settings", Type: schema.TypeJSON, Resolver: "fetchStorageAccountQueueLoggingSettings"},
						codegen.ColumnDefinition{Name: "is_nfs_v3_enabled", Type: schema.TypeBool, Resolver: `schema.PathResolver("EnableNfsV3")`},
					},
					relations: accountRelations,
					mockListFunctionArgsInit: []string{
						`result.Values()[0].Sku.Tier = storage.Standard`,
						`result.Values()[0].Kind = storage.StorageV2`,

						`blobProperties := accounts.StorageServiceProperties{}`,
						`require.Nil(t, faker.FakeObject(&blobProperties))`,
						`blobResult := accounts.GetServicePropertiesResult{StorageServiceProperties: &blobProperties}`,
						`mockClient.EXPECT().GetBlobServiceProperties(gomock.Any(), "test", "test").Return(blobResult, nil)`,

						`queueProperties := queues.StorageServiceProperties{}`,
						`require.Nil(t, faker.FakeObject(&queueProperties))`,
						`queueResult :=  queues.StorageServicePropertiesResponse{StorageServiceProperties: queueProperties}`,
						`mockClient.EXPECT().GetQueueServiceProperties(gomock.Any(), "test", "test").Return(queueResult, nil)`,
					},
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage", "github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage", "github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"},
				},
			},
		},
	}

	initParents(resourcesByTemplates)

	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, accountRelations[0])
	resourcesByTemplates[1].definitions = []resourceDefinition{accountRelations[1]}

	return generateResources(resourcesByTemplates)
}
