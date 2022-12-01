package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

func DataLakeAnalyticsResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "accounts",
			Struct: &armdatalakeanalytics.AccountBasic{},
			ResponseStruct: &armdatalakeanalytics.AccountsClientListResponse{},
			Client: &armdatalakeanalytics.AccountsClient{},
			ListFunc: (&armdatalakeanalytics.AccountsClient{}).NewListPager,
			NewFunc: armdatalakeanalytics.NewAccountsClient,
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
		r.ImportPath = "datalake-analytics/armdatalakeanalytics"
		r.Service = "armdatalakeanalytics"
		r.Template = "list"
	}

	return resources
}