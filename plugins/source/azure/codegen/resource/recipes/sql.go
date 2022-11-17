package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func SQL() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armsql.ManagedInstance),
			Resolver: &resource.FuncParams{
				Func: sql.ManagedInstancesClient.NewListPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armsql.ManagedDatabase),
					Resolver: &resource.FuncParams{
						Func:   sql.ManagedDatabasesClient.NewListByInstancePager,
						Params: []string{"id.ResourceGroupName", "*managedInstance.Name"},
					},
					Children: []*resource.Resource{
						{
							SubService: "managed_database_vulnerability_assessments",
							Struct:     new(armsql.DatabaseVulnerabilityAssessment),
							Resolver: &resource.FuncParams{
								Func:   sql.ManagedDatabaseVulnerabilityAssessmentsClient.NewListByDatabasePager,
								Params: []string{"id.ResourceGroupName", "*managedInstance.Name", "*managedDatabase.Name"},
							},
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
					Struct: new(armsql.ManagedInstanceVulnerabilityAssessment),
					Resolver: &resource.FuncParams{
						Func:   sql.ManagedInstanceVulnerabilityAssessmentsClient.NewListByInstancePager,
						Params: []string{"id.ResourceGroupName", "*managedInstance.Name"},
					},
				},
				{
					Struct: new(armsql.ManagedInstanceEncryptionProtector),
					Resolver: &resource.FuncParams{
						Func:   sql.ManagedInstanceEncryptionProtectorsClient.NewListByInstancePager,
						Params: []string{"id.ResourceGroupName", "*managedInstance.Name"},
					},
				},
			},
		},
		{
			Struct: new(armsql.Server),
			Resolver: &resource.FuncParams{
				Func: sql.ServersClient.NewListPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armsql.FirewallRule),
					Resolver: &resource.FuncParams{
						Func:   sql.FirewallRulesClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
				{
					Struct: new(armsql.Database),
					Resolver: &resource.FuncParams{
						Func:   sql.DatabasesClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
					Children: []*resource.Resource{
						{
							Struct: new(armsql.DatabaseBlobAuditingPolicy),
							Resolver: &resource.FuncParams{
								Func:   sql.DatabaseBlobAuditingPoliciesClient.NewListByDatabasePager,
								Params: []string{"id.ResourceGroupName", "*server.Name", "*database.Name"},
							},
						},
						{
							Struct: new(armsql.DatabaseVulnerabilityAssessment),
							Resolver: &resource.FuncParams{
								Func:   sql.DatabaseVulnerabilityAssessmentsClient.NewListByDatabasePager,
								Params: []string{"id.ResourceGroupName", "*server.Name", "*database.Name"},
							},
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
							Resolver: &resource.FuncParams{
								Func:   sql.LongTermRetentionPoliciesClient.NewListByDatabasePager,
								Params: []string{"id.ResourceGroupName", "*server.Name", "*database.Name"},
							},
						},
						{
							SubService: "database_threat_detection_policies",
							Struct:     new(armsql.DatabaseSecurityAlertPolicy),
							Resolver: &resource.FuncParams{
								Func:   sql.DatabaseSecurityAlertPoliciesClient.NewListByDatabasePager,
								Params: []string{"id.ResourceGroupName", "*server.Name", "*database.Name"},
							},
						},
						{
							SubService: "transparent_data_encryptions",
							Struct:     new(armsql.LogicalDatabaseTransparentDataEncryption),
							Resolver: &resource.FuncParams{
								Func:   sql.TransparentDataEncryptionsClient.NewListByDatabasePager,
								Params: []string{"id.ResourceGroupName", "*server.Name", "*database.Name"},
							},
						},
					},
				},
				{
					Struct: new(armsql.EncryptionProtector),
					Resolver: &resource.FuncParams{
						Func:   sql.EncryptionProtectorsClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
				{
					Struct: new(armsql.VirtualNetworkRule),
					Resolver: &resource.FuncParams{
						Func:   sql.VirtualNetworkRulesClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
				{
					SubService: "server_administrators",
					Struct:     new(armsql.ServerAzureADAdministrator),
					Resolver: &resource.FuncParams{
						Func:   sql.ServerAzureADAdministratorsClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
				{
					Struct: new(armsql.ServerBlobAuditingPolicy),
					Resolver: &resource.FuncParams{
						Func:   sql.ServerBlobAuditingPoliciesClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
				{
					Struct: new(armsql.ServerDevOpsAuditingSettings),
					Resolver: &resource.FuncParams{
						Func:   sql.ServerDevOpsAuditSettingsClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
				{
					Struct: new(armsql.ServerVulnerabilityAssessment),
					Resolver: &resource.FuncParams{
						Func:   sql.ServerVulnerabilityAssessmentsClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
				{
					Struct: new(armsql.ServerSecurityAlertPolicy),
					Resolver: &resource.FuncParams{
						Func:   sql.ServerSecurityAlertPoliciesClient.NewListByServerPager,
						Params: []string{"id.ResourceGroupName", "*server.Name"},
					},
				},
			},
		},
	}
}
