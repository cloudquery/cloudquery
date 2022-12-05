// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"

func Armsecurity() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armsecurity.NewAccountConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/connectors",
		},
		{
			NewFunc: armsecurity.NewAdaptiveApplicationControlsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/applicationWhitelistings",
		},
		{
			NewFunc: armsecurity.NewAdaptiveNetworkHardeningsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewAdvancedThreatProtectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alerts",
		},
		{
			NewFunc: armsecurity.NewAlertsSuppressionRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alertsSuppressionRules",
		},
		{
			NewFunc: armsecurity.NewAllowedConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/allowedConnections",
		},
		{
			NewFunc: armsecurity.NewApplicationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewApplicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/applications",
		},
		{
			NewFunc: armsecurity.NewAssessmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{scope}/providers/Microsoft.Security/assessments",
		},
		{
			NewFunc: armsecurity.NewAssessmentsMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/providers/Microsoft.Security/assessmentMetadata",
		},
		{
			NewFunc: armsecurity.NewAutoProvisioningSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings",
		},
		{
			NewFunc: armsecurity.NewAutomationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/automations",
		},
		{
			NewFunc: armsecurity.NewComplianceResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{scope}/providers/Microsoft.Security/complianceResults",
		},
		{
			NewFunc: armsecurity.NewCompliancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{scope}/providers/Microsoft.Security/compliances",
		},
		{
			NewFunc: armsecurity.NewConnectorApplicationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewConnectorApplicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/applications",
		},
		{
			NewFunc: armsecurity.NewConnectorGovernanceRuleClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules",
		},
		{
			NewFunc: armsecurity.NewConnectorGovernanceRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewConnectorGovernanceRulesExecuteStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityConnectors",
		},
		{
			NewFunc: armsecurity.NewContactsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityContacts",
		},
		{
			NewFunc: armsecurity.NewCustomAssessmentAutomationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewCustomEntityStoreAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewDeviceSecurityGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups",
		},
		{
			NewFunc: armsecurity.NewDiscoveredSecuritySolutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/discoveredSecuritySolutions",
		},
		{
			NewFunc: armsecurity.NewExternalSecuritySolutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/externalSecuritySolutions",
		},
		{
			NewFunc: armsecurity.NewGovernanceAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments",
		},
		{
			NewFunc: armsecurity.NewGovernanceRuleClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules",
		},
		{
			NewFunc: armsecurity.NewGovernanceRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewInformationProtectionPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{scope}/providers/Microsoft.Security/informationProtectionPolicies",
		},
		{
			NewFunc: armsecurity.NewIngestionSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings",
		},
		{
			NewFunc: armsecurity.NewIotSecuritySolutionAnalyticsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels",
		},
		{
			NewFunc: armsecurity.NewIotSecuritySolutionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewIotSecuritySolutionsAnalyticsAggregatedAlertClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts",
		},
		{
			NewFunc: armsecurity.NewIotSecuritySolutionsAnalyticsRecommendationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations",
		},
		{
			NewFunc: armsecurity.NewJitNetworkAccessPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/jitNetworkAccessPolicies",
		},
		{
			NewFunc: armsecurity.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations",
		},
		{
			NewFunc: armsecurity.NewMdeOnboardingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/mdeOnboardings",
		},
		{
			NewFunc: armsecurity.NewPricingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/pricings",
		},
		{
			NewFunc: armsecurity.NewRegulatoryComplianceAssessmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments",
		},
		{
			NewFunc: armsecurity.NewRegulatoryComplianceControlsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls",
		},
		{
			NewFunc: armsecurity.NewRegulatoryComplianceStandardsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards",
		},
		{
			NewFunc: armsecurity.NewSQLVulnerabilityAssessmentBaselineRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{resourceId}/providers/Microsoft.Security/sqlVulnerabilityAssessments/default/baselineRules",
		},
		{
			NewFunc: armsecurity.NewSQLVulnerabilityAssessmentScanResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{resourceId}/providers/Microsoft.Security/sqlVulnerabilityAssessments/default/scans/{scanId}/scanResults",
		},
		{
			NewFunc: armsecurity.NewSQLVulnerabilityAssessmentScansClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{resourceId}/providers/Microsoft.Security/sqlVulnerabilityAssessments/default/scans",
		},
		{
			NewFunc: armsecurity.NewSecureScoreControlDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/providers/Microsoft.Security/secureScoreControlDefinitions",
		},
		{
			NewFunc: armsecurity.NewSecureScoreControlsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControls",
		},
		{
			NewFunc: armsecurity.NewSecureScoresClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores",
		},
		{
			NewFunc: armsecurity.NewServerVulnerabilityAssessmentClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings",
		},
		{
			NewFunc: armsecurity.NewSoftwareInventoriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewSolutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutions",
		},
		{
			NewFunc: armsecurity.NewSolutionsReferenceDataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutionsReferenceData",
		},
		{
			NewFunc: armsecurity.NewSubAssessmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments",
		},
		{
			NewFunc: armsecurity.NewSubscriptionGovernanceRulesExecuteStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "",
		},
		{
			NewFunc: armsecurity.NewTasksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/tasks",
		},
		{
			NewFunc: armsecurity.NewTopologyClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/topologies",
		},
		{
			NewFunc: armsecurity.NewWorkspaceSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armsecurity())
}