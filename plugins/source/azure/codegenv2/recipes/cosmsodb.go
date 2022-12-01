package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

func CosmosDBResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "accounts",
			Struct: &armcosmos.DatabaseAccountGetResults{},
			ResponseStruct: &armcosmos.DatabaseAccountsClientListResponse{},
			Client: &armcosmos.DatabaseAccountsClient{},
			ListFunc: (&armcosmos.DatabaseAccountsClient{}).NewListPager,
			NewFunc: armcosmos.NewDatabaseAccountsClient,
			OutputField: "Value",
		},
		// {
		// 	SubService: "sql_databases",
		// 	Struct: &armcosmos.RestorableSQLDatabaseGetResult{},
		// 	ResponseStruct: &armcosmos.RestorableSQLDatabasesClientListResponse{},
		// 	Client: &armcosmos.RestorableSQLDatabasesClient{},
		// 	ListFunc: (&armcosmos.RestorableSQLDatabasesClient{}).NewListPager,
		// 	NewFunc: armcosmos.NewRestorableSQLDatabasesClient,
		// 	OutputField: "Value",
		// },
	}

	for _, r := range resources {
		r.ImportPath = "cosmos/armcosmos"
		r.Service = "armcosmos"
		r.Template = "list"
	}

	return resources
}