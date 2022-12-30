package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

func SqlServersE() []Table {
	tables := []Table{
		{
			Service:        "armsql",
			Name:           "servers",
			Struct:         &armsql.Server{},
			ResponseStruct: &armsql.ServersClientListResponse{},
			Client:         &armsql.ServersClient{},
			ListFunc:       (&armsql.ServersClient{}).NewListPager,
			NewFunc:        armsql.NewServersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/servers",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_sql)`,
			ExtraColumns:   DefaultExtraColumns,
			Relations: []*Table{
				{
					Service:        "armsql",
					Name:           "server_databases",
					Struct:         &armsql.Database{},
					ResponseStruct: &armsql.DatabasesClientListByServerResponse{},
					Client:         &armsql.DatabasesClient{},
					ListFunc:       (&armsql.DatabasesClient{}).NewListByServerPager,
					NewFunc:        armsql.NewDatabasesClient,
					URL:            `/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Sql/servers/test string/databases`,
					SkipFetch:      true,
					Relations: []*Table{
						{
							Service:        "armsql",
							Name:           "server_database_blob_auditing_policies",
							Struct:         &armsql.DatabaseBlobAuditingPolicy{},
							ResponseStruct: &armsql.DatabaseBlobAuditingPolicyListResult{},
							Client:         &armsql.DatabaseBlobAuditingPoliciesClient{},
							ListFunc:       (&armsql.DatabaseBlobAuditingPoliciesClient{}).NewListByDatabasePager,
							NewFunc:        armsql.NewDatabaseBlobAuditingPoliciesClient,
							URL:            "/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Sql/servers/test string/databases/test string/auditingSettings",
							SkipFetch:      true,
						},
					},
				},
			},
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, SqlServersE()...)
}
