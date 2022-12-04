// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"

func Armsql() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armsql.NewManagedInstanceLongTermRetentionPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewRestorableDroppedManagedDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewVirtualNetworkRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewElasticPoolDatabaseActivitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseTablesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobAgentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobStepsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseVulnerabilityAssessmentScansClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewExtendedDatabaseBlobAuditingPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewIPv6FirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewLedgerDigestUploadsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedBackupShortTermRetentionPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedServerSecurityAlertPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerAutomaticTuningClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDataMaskingPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobExecutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseSecurityAlertPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseQueriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseSchemasClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstanceKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstanceVulnerabilityAssessmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewOutboundFirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerConnectionPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewMaintenanceWindowsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewSyncAgentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseSecurityEventsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstanceEncryptionProtectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerTrustCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDataMaskingRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewSensitivityLabelsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseSensitivityLabelsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseExtensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseVulnerabilityAssessmentScansClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseAdvisorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerCommunicationLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewFailoverGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewInstancePoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewRecoverableDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseTablesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewEndpointCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewMaintenanceWindowOptionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseVulnerabilityAssessmentRuleBaselinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewRecoverableManagedDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewRestorableDroppedDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewTransparentDataEncryptionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseColumnsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDistributedAvailabilityGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewFirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewGeoBackupPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseSecurityAlertPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstanceAzureADOnlyAuthenticationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDataWarehouseUserActivitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewElasticPoolActivitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerAdvancedThreatProtectionSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerAzureADAdministratorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerSecurityAlertPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewSyncMembersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewTimeZonesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobStepExecutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstancePrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewCapabilitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobVersionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseRecommendedSensitivityLabelsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstanceOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstancePrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerVulnerabilityAssessmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobCredentialsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseVulnerabilityAssessmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstanceAdministratorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseAutomaticTuningClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerDNSAliasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServiceObjectivesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerBlobAuditingPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseRestoreDetailsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDeletedServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobTargetGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewTdeCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseColumnsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewRestorePointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerDevOpsAuditSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewVirtualClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseRecommendedActionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseVulnerabilityAssessmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewSyncGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewElasticPoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerAzureADOnlyAuthenticationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewLongTermRetentionManagedInstanceBackupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewReplicationLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewSubscriptionUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewExtendedServerBlobAuditingPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewLongTermRetentionBackupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewWorkloadClassifiersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewWorkloadGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseSchemasClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedInstanceTdeCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewInstanceFailoverGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewLongTermRetentionPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseAdvancedThreatProtectionSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseVulnerabilityAssessmentRuleBaselinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewElasticPoolOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerTrustGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewJobTargetExecutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewBackupShortTermRetentionPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewEncryptionProtectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerAdvisorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewServerUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewDatabaseBlobAuditingPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewRecommendedSensitivityLabelsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewAgentClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
		{
			NewFunc: armsql.NewManagedDatabaseTransparentDataEncryptionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armsql())
}