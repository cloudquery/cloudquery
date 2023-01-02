package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"

func init() {
	tables := []Table{
		{
			Service:        "armstorage",
			Name:           "accounts",
			Struct:         &armstorage.Account{},
			ResponseStruct: &armstorage.AccountsClientListResponse{},
			Client:         &armstorage.AccountsClient{},
			ListFunc:       (&armstorage.AccountsClient{}).NewListPager,
			NewFunc:        armstorage.NewAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_storage)`,
			Relations: []*Table{
				{
					Service:        "armstorage",
					Name:           "tables",
					Struct:         &armstorage.Table{},
					ResponseStruct: &armstorage.TableClientListResponse{},
					Client:         &armstorage.AccountsClient{},
					ListFunc:       (&armstorage.TableClient{}).NewListPager,
					NewFunc:        armstorage.NewTableClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables",
					SkipFetch:      true,
				},
				{
					Service:        "armstorage",
					Name:           "containers",
					Struct:         &armstorage.ListContainerItem{},
					ResponseStruct: &armstorage.BlobContainersClientListResponse{},
					Client:         &armstorage.BlobContainersClient{},
					ListFunc:       (&armstorage.BlobContainersClient{}).NewListPager,
					NewFunc:        armstorage.NewBlobContainersClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers",
					SkipFetch:      true,
				},
				{
					Service:        "armstorage",
					Name:           "blob_services",
					Struct:         &armstorage.BlobServiceProperties{},
					ResponseStruct: &armstorage.BlobServicesClientListResponse{},
					Client:         &armstorage.BlobServicesClient{},
					ListFunc:       (&armstorage.BlobServicesClient{}).NewListPager,
					NewFunc:        armstorage.NewBlobServicesClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices",
					SkipFetch:      true,
				},
			},
		},
	}
	Tables = append(Tables, tables...)
}
