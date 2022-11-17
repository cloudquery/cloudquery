package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Storage() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armstorage.Account),
			Resolver: storage.AccountsClient.NewListPager,
			Children: []*resource.Resource{
				{
					SubService: "containers",
					Struct:     new(armstorage.ListContainerItem),
					Resolver:   storage.BlobContainersClient.NewListPager,
				},
				{
					SubService:         "blob_services",
					Struct:             new(armstorage.BlobServiceProperties),
					UnwrapStructFields: []string{"BlobServiceProperties"},
					Resolver:           storage.BlobServicesClient.NewListPager,
				},
			},
		},
	}
}
