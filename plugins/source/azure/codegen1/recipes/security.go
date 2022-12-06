// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"

func Armsecurity() []*Table {
	tables := []*Table{
		{
			NewFunc: armsecurity.NewApplicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/applications",
		},
		{
			NewFunc: armsecurity.NewWorkspaceSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings",
		},
		{
			NewFunc: armsecurity.NewAutomationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/automations",
		},
		{
			NewFunc: armsecurity.NewSolutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutions",
		},
		{
			NewFunc: armsecurity.NewGovernanceRuleClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules",
		},
		{
			NewFunc: armsecurity.NewSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings",
		},
		{
			NewFunc: armsecurity.NewSecureScoreControlDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/providers/Microsoft.Security/secureScoreControlDefinitions",
		},
		{
			NewFunc: armsecurity.NewSecureScoresClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores",
		},
		{
			NewFunc: armsecurity.NewTasksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/tasks",
		},
		{
			NewFunc: armsecurity.NewSecureScoreControlsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControls",
		},
		{
			NewFunc: armsecurity.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations",
		},
		{
			NewFunc: armsecurity.NewAutoProvisioningSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings",
		},
		{
			NewFunc: armsecurity.NewDiscoveredSecuritySolutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/discoveredSecuritySolutions",
		},
		{
			NewFunc: armsecurity.NewAllowedConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/allowedConnections",
		},
		{
			NewFunc: armsecurity.NewJitNetworkAccessPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/jitNetworkAccessPolicies",
		},
		{
			NewFunc: armsecurity.NewExternalSecuritySolutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/externalSecuritySolutions",
		},
		{
			NewFunc: armsecurity.NewAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alerts",
		},
		{
			NewFunc: armsecurity.NewTopologyClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/topologies",
		},
		{
			NewFunc: armsecurity.NewContactsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityContacts",
		},
		{
			NewFunc: armsecurity.NewAlertsSuppressionRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alertsSuppressionRules",
		},
		{
			NewFunc: armsecurity.NewAccountConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/connectors",
		},
		{
			NewFunc: armsecurity.NewIngestionSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings",
		},
		{
			NewFunc: armsecurity.NewAssessmentsMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/providers/Microsoft.Security/assessmentMetadata",
		},
		{
			NewFunc: armsecurity.NewConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityConnectors",
		},
		{
			NewFunc: armsecurity.NewRegulatoryComplianceStandardsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsecurity())
}
