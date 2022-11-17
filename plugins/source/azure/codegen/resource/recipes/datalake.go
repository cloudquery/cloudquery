package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/datalakeanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/datalakestore"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func DataLake() []*resource.Resource {
	return []*resource.Resource{
		{
			Service: "datalake",
			Struct:  new(armdatalakeanalytics.Account),
			Resolver: &resource.FuncParams{
				Func: datalakeanalytics.AccountsClient.NewListPager,
			},
			PreResolver: &resource.FuncParams{
				Func:       datalakeanalytics.AccountsClient.Get,
				Params:     []string{"ctx", "id.ResourceGroupName", "*accountBasic.Name"},
				BasicValue: new(armdatalakeanalytics.AccountBasic),
			},
		},
		{
			Service:    "datalake",
			SubService: "store_accounts",
			Struct:     new(armdatalakestore.Account),
			Resolver: &resource.FuncParams{
				Func: datalakestore.AccountsClient.NewListPager,
			},
			PreResolver: &resource.FuncParams{
				Func:       datalakestore.AccountsClient.Get,
				Params:     []string{"ctx", "id.ResourceGroupName", "*accountBasic.Name"},
				BasicValue: new(armdatalakestore.AccountBasic),
			},
		},
	}
}
