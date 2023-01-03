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
					Name:           "server_vulnerability_assessments",
					Struct:         &armsql.ServerVulnerabilityAssessment{},
					ResponseStruct: &armsql.ServerVulnerabilityAssessmentsClientListByServerResponse{},
					Client:         &armsql.ServerVulnerabilityAssessmentsClient{},
					ListFunc:       (&armsql.ServerVulnerabilityAssessmentsClient{}).NewListByServerPager,
					NewFunc:        armsql.NewServerVulnerabilityAssessmentsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Sql/servers/test string/vulnerabilityAssessments",
					SkipFetch:      true,
				},
				{
					Service:        "armsql",
					Name:           "server_admins",
					Struct:         &armsql.ServerAzureADAdministrator{},
					ResponseStruct: &armsql.ServerAzureADAdministratorsClientListByServerResponse{},
					Client:         &armsql.ServerAzureADAdministratorsClient{},
					ListFunc:       (&armsql.ServerAzureADAdministratorsClient{}).NewListByServerPager,
					NewFunc:        armsql.NewServerAzureADAdministratorsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Sql/servers/test string/administrators",
					SkipFetch:      true,
				},
				{
					Service:        "armsql",
					Name:           "encryption_protectors",
					Struct:         &armsql.EncryptionProtector{},
					ResponseStruct: &armsql.EncryptionProtectorsClientListByServerResponse{},
					Client:         &armsql.EncryptionProtectorsClient{},
					ListFunc:       (&armsql.EncryptionProtectorsClient{}).NewListByServerPager,
					NewFunc:        armsql.NewEncryptionProtectorsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Sql/servers/test string/encryptionProtector",
					SkipFetch:      true,
				},
				{
					Service:        "armsql",
					Name:           "databases",
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
							Name:           "database_blob_auditing_policies",
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
