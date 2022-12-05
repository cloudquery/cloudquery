// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"

func Armsynapse() []Table {
	tables := []Table{
		{
      Name: "azure_ad_only_authentication",
      Struct: &armsynapse.AzureADOnlyAuthentication{},
      ResponseStruct: &armsynapse.AzureADOnlyAuthenticationsClientListResponse{},
      Client: &armsynapse.AzureADOnlyAuthenticationsClient{},
      ListFunc: (&armsynapse.AzureADOnlyAuthenticationsClient{}).NewListPager,
			NewFunc: armsynapse.NewAzureADOnlyAuthenticationsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/azureADOnlyAuthentications",
		},
		{
      Name: "operation",
      Struct: &armsynapse.Operation{},
      ResponseStruct: &armsynapse.KustoOperationsClientListResponse{},
      Client: &armsynapse.KustoOperationsClient{},
      ListFunc: (&armsynapse.KustoOperationsClient{}).NewListPager,
			NewFunc: armsynapse.NewKustoOperationsClient,
			URL: "/providers/Microsoft.Synapse/kustooperations",
		},
		{
      Name: "database_principal_assignment",
      Struct: &armsynapse.DatabasePrincipalAssignment{},
      ResponseStruct: &armsynapse.KustoPoolDatabasePrincipalAssignmentsClientListResponse{},
      Client: &armsynapse.KustoPoolDatabasePrincipalAssignmentsClient{},
      ListFunc: (&armsynapse.KustoPoolDatabasePrincipalAssignmentsClient{}).NewListPager,
			NewFunc: armsynapse.NewKustoPoolDatabasePrincipalAssignmentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/principalAssignments",
		},
		{
      Name: "cluster_principal_assignment",
      Struct: &armsynapse.ClusterPrincipalAssignment{},
      ResponseStruct: &armsynapse.KustoPoolPrincipalAssignmentsClientListResponse{},
      Client: &armsynapse.KustoPoolPrincipalAssignmentsClient{},
      ListFunc: (&armsynapse.KustoPoolPrincipalAssignmentsClient{}).NewListPager,
			NewFunc: armsynapse.NewKustoPoolPrincipalAssignmentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/principalAssignments",
		},
		{
      Name: "private_endpoint_connection",
      Struct: &armsynapse.PrivateEndpointConnection{},
      ResponseStruct: &armsynapse.PrivateEndpointConnectionsClientListResponse{},
      Client: &armsynapse.PrivateEndpointConnectionsClient{},
      ListFunc: (&armsynapse.PrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armsynapse.NewPrivateEndpointConnectionsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/privateEndpointConnections",
		},
		{
      Name: "private_endpoint_connection_for_private_link_hub",
      Struct: &armsynapse.PrivateEndpointConnectionForPrivateLinkHub{},
      ResponseStruct: &armsynapse.PrivateEndpointConnectionsPrivateLinkHubClientListResponse{},
      Client: &armsynapse.PrivateEndpointConnectionsPrivateLinkHubClient{},
      ListFunc: (&armsynapse.PrivateEndpointConnectionsPrivateLinkHubClient{}).NewListPager,
			NewFunc: armsynapse.NewPrivateEndpointConnectionsPrivateLinkHubClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateEndpointConnections",
		},
		{
      Name: "private_link_resource",
      Struct: &armsynapse.PrivateLinkResource{},
      ResponseStruct: &armsynapse.PrivateLinkHubPrivateLinkResourcesClientListResponse{},
      Client: &armsynapse.PrivateLinkHubPrivateLinkResourcesClient{},
      ListFunc: (&armsynapse.PrivateLinkHubPrivateLinkResourcesClient{}).NewListPager,
			NewFunc: armsynapse.NewPrivateLinkHubPrivateLinkResourcesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateLinkResources",
		},
		{
      Name: "private_link_hub",
      Struct: &armsynapse.PrivateLinkHub{},
      ResponseStruct: &armsynapse.PrivateLinkHubsClientListResponse{},
      Client: &armsynapse.PrivateLinkHubsClient{},
      ListFunc: (&armsynapse.PrivateLinkHubsClient{}).NewListPager,
			NewFunc: armsynapse.NewPrivateLinkHubsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/privateLinkHubs",
		},
		{
      Name: "private_link_resource",
      Struct: &armsynapse.PrivateLinkResource{},
      ResponseStruct: &armsynapse.PrivateLinkResourcesClientListResponse{},
      Client: &armsynapse.PrivateLinkResourcesClient{},
      ListFunc: (&armsynapse.PrivateLinkResourcesClient{}).NewListPager,
			NewFunc: armsynapse.NewPrivateLinkResourcesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/privateLinkResources",
		},
		{
      Name: "geo_backup_policy",
      Struct: &armsynapse.GeoBackupPolicy{},
      ResponseStruct: &armsynapse.SQLPoolGeoBackupPoliciesClientListResponse{},
      Client: &armsynapse.SQLPoolGeoBackupPoliciesClient{},
      ListFunc: (&armsynapse.SQLPoolGeoBackupPoliciesClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolGeoBackupPoliciesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/geoBackupPolicies",
		},
		{
      Name: "sql_pool_operation",
      Struct: &armsynapse.SQLPoolOperation{},
      ResponseStruct: &armsynapse.SQLPoolOperationsClientListResponse{},
      Client: &armsynapse.SQLPoolOperationsClient{},
      ListFunc: (&armsynapse.SQLPoolOperationsClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolOperationsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operations",
		},
		{
      Name: "replication_link",
      Struct: &armsynapse.ReplicationLink{},
      ResponseStruct: &armsynapse.SQLPoolReplicationLinksClientListResponse{},
      Client: &armsynapse.SQLPoolReplicationLinksClient{},
      ListFunc: (&armsynapse.SQLPoolReplicationLinksClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolReplicationLinksClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/replicationLinks",
		},
		{
      Name: "restore_point",
      Struct: &armsynapse.RestorePoint{},
      ResponseStruct: &armsynapse.SQLPoolRestorePointsClientListResponse{},
      Client: &armsynapse.SQLPoolRestorePointsClient{},
      ListFunc: (&armsynapse.SQLPoolRestorePointsClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolRestorePointsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/restorePoints",
		},
		{
      Name: "sql_pool_schema",
      Struct: &armsynapse.SQLPoolSchema{},
      ResponseStruct: &armsynapse.SQLPoolSchemasClientListResponse{},
      Client: &armsynapse.SQLPoolSchemasClient{},
      ListFunc: (&armsynapse.SQLPoolSchemasClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolSchemasClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas",
		},
		{
      Name: "sql_pool_security_alert_policy",
      Struct: &armsynapse.SQLPoolSecurityAlertPolicy{},
      ResponseStruct: &armsynapse.SQLPoolSecurityAlertPoliciesClientListResponse{},
      Client: &armsynapse.SQLPoolSecurityAlertPoliciesClient{},
      ListFunc: (&armsynapse.SQLPoolSecurityAlertPoliciesClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolSecurityAlertPoliciesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/securityAlertPolicies",
		},
		{
      Name: "transparent_data_encryption",
      Struct: &armsynapse.TransparentDataEncryption{},
      ResponseStruct: &armsynapse.SQLPoolTransparentDataEncryptionsClientListResponse{},
      Client: &armsynapse.SQLPoolTransparentDataEncryptionsClient{},
      ListFunc: (&armsynapse.SQLPoolTransparentDataEncryptionsClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolTransparentDataEncryptionsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption",
		},
		{
      Name: "sql_pool_usage",
      Struct: &armsynapse.SQLPoolUsage{},
      ResponseStruct: &armsynapse.SQLPoolUsagesClientListResponse{},
      Client: &armsynapse.SQLPoolUsagesClient{},
      ListFunc: (&armsynapse.SQLPoolUsagesClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolUsagesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/usages",
		},
		{
      Name: "vulnerability_assessment_scan_record",
      Struct: &armsynapse.VulnerabilityAssessmentScanRecord{},
      ResponseStruct: &armsynapse.SQLPoolVulnerabilityAssessmentScansClientListResponse{},
      Client: &armsynapse.SQLPoolVulnerabilityAssessmentScansClient{},
      ListFunc: (&armsynapse.SQLPoolVulnerabilityAssessmentScansClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolVulnerabilityAssessmentScansClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/vulnerabilityAssessments/{vulnerabilityAssessmentName}/scans",
		},
		{
      Name: "sql_pool_vulnerability_assessment",
      Struct: &armsynapse.SQLPoolVulnerabilityAssessment{},
      ResponseStruct: &armsynapse.SQLPoolVulnerabilityAssessmentsClientListResponse{},
      Client: &armsynapse.SQLPoolVulnerabilityAssessmentsClient{},
      ListFunc: (&armsynapse.SQLPoolVulnerabilityAssessmentsClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolVulnerabilityAssessmentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/vulnerabilityAssessments",
		},
		{
      Name: "workload_classifier",
      Struct: &armsynapse.WorkloadClassifier{},
      ResponseStruct: &armsynapse.SQLPoolWorkloadClassifierClientListResponse{},
      Client: &armsynapse.SQLPoolWorkloadClassifierClient{},
      ListFunc: (&armsynapse.SQLPoolWorkloadClassifierClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolWorkloadClassifierClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers",
		},
		{
      Name: "workload_group",
      Struct: &armsynapse.WorkloadGroup{},
      ResponseStruct: &armsynapse.SQLPoolWorkloadGroupClientListResponse{},
      Client: &armsynapse.SQLPoolWorkloadGroupClient{},
      ListFunc: (&armsynapse.SQLPoolWorkloadGroupClient{}).NewListPager,
			NewFunc: armsynapse.NewSQLPoolWorkloadGroupClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups",
		},
		{
      Name: "dedicated_sq_lminimal_tls_settings",
      Struct: &armsynapse.DedicatedSQLminimalTLSSettings{},
      ResponseStruct: &armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse{},
      Client: &armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient{},
      ListFunc: (&armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient{}).NewListPager,
			NewFunc: armsynapse.NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/dedicatedSQLminimalTlsSettings",
		},
		{
      Name: "encryption_protector",
      Struct: &armsynapse.EncryptionProtector{},
      ResponseStruct: &armsynapse.WorkspaceManagedSQLServerEncryptionProtectorClientListResponse{},
      Client: &armsynapse.WorkspaceManagedSQLServerEncryptionProtectorClient{},
      ListFunc: (&armsynapse.WorkspaceManagedSQLServerEncryptionProtectorClient{}).NewListPager,
			NewFunc: armsynapse.NewWorkspaceManagedSQLServerEncryptionProtectorClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector",
		},
		{
      Name: "recoverable_sql_pool",
      Struct: &armsynapse.RecoverableSQLPool{},
      ResponseStruct: &armsynapse.WorkspaceManagedSQLServerRecoverableSQLPoolsClientListResponse{},
      Client: &armsynapse.WorkspaceManagedSQLServerRecoverableSQLPoolsClient{},
      ListFunc: (&armsynapse.WorkspaceManagedSQLServerRecoverableSQLPoolsClient{}).NewListPager,
			NewFunc: armsynapse.NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/recoverableSqlPools",
		},
		{
      Name: "server_security_alert_policy",
      Struct: &armsynapse.ServerSecurityAlertPolicy{},
      ResponseStruct: &armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse{},
      Client: &armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient{},
      ListFunc: (&armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient{}).NewListPager,
			NewFunc: armsynapse.NewWorkspaceManagedSQLServerSecurityAlertPolicyClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/securityAlertPolicies",
		},
		{
      Name: "server_usage",
      Struct: &armsynapse.ServerUsage{},
      ResponseStruct: &armsynapse.WorkspaceManagedSQLServerUsagesClientListResponse{},
      Client: &armsynapse.WorkspaceManagedSQLServerUsagesClient{},
      ListFunc: (&armsynapse.WorkspaceManagedSQLServerUsagesClient{}).NewListPager,
			NewFunc: armsynapse.NewWorkspaceManagedSQLServerUsagesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlUsages",
		},
		{
      Name: "server_vulnerability_assessment",
      Struct: &armsynapse.ServerVulnerabilityAssessment{},
      ResponseStruct: &armsynapse.WorkspaceManagedSQLServerVulnerabilityAssessmentsClientListResponse{},
      Client: &armsynapse.WorkspaceManagedSQLServerVulnerabilityAssessmentsClient{},
      ListFunc: (&armsynapse.WorkspaceManagedSQLServerVulnerabilityAssessmentsClient{}).NewListPager,
			NewFunc: armsynapse.NewWorkspaceManagedSQLServerVulnerabilityAssessmentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/vulnerabilityAssessments",
		},
		{
      Name: "workspace",
      Struct: &armsynapse.Workspace{},
      ResponseStruct: &armsynapse.WorkspacesClientListResponse{},
      Client: &armsynapse.WorkspacesClient{},
      ListFunc: (&armsynapse.WorkspacesClient{}).NewListPager,
			NewFunc: armsynapse.NewWorkspacesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/workspaces",
		},
	}

	for i := range tables {
		tables[i].Service = "armsynapse"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armsynapse()...)
}