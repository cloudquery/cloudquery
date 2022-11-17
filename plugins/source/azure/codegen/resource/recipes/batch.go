package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Batch() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armbatch.Account),
			Resolver: &resource.FuncParams{
				Func: batch.AccountClient.NewListPager,
			},
		},
	}
}
