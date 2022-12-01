package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

func DataLakeStoreResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "accounts",
			Struct: &armdatalakestore.AccountBasic{},
			ResponseStruct: &armdatalakestore.AccountsClientListResponse{},
			Client: &armdatalakestore.AccountsClient{},
			ListFunc: (&armdatalakestore.AccountsClient{}).NewListPager,
			NewFunc: armdatalakestore.NewAccountsClient,
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
		r.ImportPath = "datalake-store/armdatalakestore"
		r.Service = "armdatalakestore"
		r.Template = "list"
	}

	return resources
}