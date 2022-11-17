package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func SQL() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armsql.ManagedInstance),
			Resolver: sql.ManagedInstancesClient.NewListPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armsql.ManagedDatabase),
					Resolver: sql.ManagedDatabasesClient.NewListByInstancePager,
					Children: []*resource.Resource{
						{
							SubService: "managed_database_vulnerability_assessments",
							Struct:     new(armsql.DatabaseVulnerabilityAssessment),
							Resolver:   sql.ManagedDatabaseVulnerabilityAssessmentsClient.NewListByDatabasePager,
						},
						{
							SubService: "managed_database_vulnerability_assessment_scans",
							Struct:     new(armsql.VulnerabilityAssessmentScanRecord),
							Resolver: &resource.FuncParams{
								Func: sql.ManagedDatabaseVulnerabilityAssessmentScansClient.NewListByDatabasePager,
								Params: []string{
									"id.ResourceGroupName",
									"*managedInstance.Name",
									"*managedDatabase.Name",
									"sql.VulnerabilityAssessmentNameDefault",
								},
							},
						},
					},
				},
				{
					Struct:   new(armsql.ManagedInstanceVulnerabilityAssessment),
					Resolver: sql.ManagedInstanceVulnerabilityAssessmentsClient.NewListByInstancePager,
				},
				{
					Struct:   new(armsql.ManagedInstanceEncryptionProtector),
					Resolver: sql.ManagedInstanceEncryptionProtectorsClient.NewListByInstancePager,
				},
			},
		},
		{
			Struct:   new(armsql.Server),
			Resolver: sql.ServersClient.NewListPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armsql.FirewallRule),
					Resolver: sql.FirewallRulesClient.NewListByServerPager,
				},
				{
					Struct:   new(armsql.Database),
					Resolver: sql.DatabasesClient.NewListByServerPager,
					Children: []*resource.Resource{
						{
							Struct:   new(armsql.DatabaseBlobAuditingPolicy),
							Resolver: sql.DatabaseBlobAuditingPoliciesClient.NewListByDatabasePager,
						},
						{
							Struct:   new(armsql.DatabaseVulnerabilityAssessment),
							Resolver: sql.DatabaseVulnerabilityAssessmentsClient.NewListByDatabasePager,
						},
						{
							SubService: "database_vulnerability_assessment_scans",
							Struct:     new(armsql.VulnerabilityAssessmentScanRecord),
							Resolver: &resource.FuncParams{
								Func: sql.DatabaseVulnerabilityAssessmentScansClient.NewListByDatabasePager,
								Params: []string{
									"id.ResourceGroupName",
									"*server.Name",
									"*database.Name",
									"sql.VulnerabilityAssessmentNameDefault",
								},
							},
						},
						{
							SubService: "backup_long_term_retention_policies",
							Struct:     new(armsql.LongTermRetentionPolicy),
							Resolver:   sql.LongTermRetentionPoliciesClient.NewListByDatabasePager,
						},
						{
							SubService: "database_threat_detection_policies",
							Struct:     new(armsql.DatabaseSecurityAlertPolicy),
							Resolver:   sql.DatabaseSecurityAlertPoliciesClient.NewListByDatabasePager,
						},
						{
							SubService: "transparent_data_encryptions",
							Struct:     new(armsql.LogicalDatabaseTransparentDataEncryption),
							Resolver:   sql.TransparentDataEncryptionsClient.NewListByDatabasePager,
						},
					},
				},
				{
					Struct:   new(armsql.EncryptionProtector),
					Resolver: sql.EncryptionProtectorsClient.NewListByServerPager,
				},
				{
					Struct:   new(armsql.VirtualNetworkRule),
					Resolver: sql.VirtualNetworkRulesClient.NewListByServerPager,
				},
				{
					SubService: "server_administrators",
					Struct:     new(armsql.ServerAzureADAdministrator),
					Resolver:   sql.ServerAzureADAdministratorsClient.NewListByServerPager,
				},
				{
					Struct:   new(armsql.ServerBlobAuditingPolicy),
					Resolver: sql.ServerBlobAuditingPoliciesClient.NewListByServerPager,
				},
				{
					Struct:   new(armsql.ServerDevOpsAuditingSettings),
					Resolver: sql.ServerDevOpsAuditSettingsClient.NewListByServerPager,
				},
				{
					Struct:   new(armsql.ServerVulnerabilityAssessment),
					Resolver: sql.ServerVulnerabilityAssessmentsClient.NewListByServerPager,
				},
				{
					Struct:   new(armsql.ServerSecurityAlertPolicy),
					Resolver: sql.ServerSecurityAlertPoliciesClient.NewListByServerPager,
				},
			},
		},
	}
}
