// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"

func init() {
	tables := []Table{
		{
			Service:        "armsecurity",
			Name:           "alerts",
			Struct:         &armsecurity.Alert{},
			ResponseStruct: &armsecurity.AlertsClientListResponse{},
			Client:         &armsecurity.AlertsClient{},
			ListFunc:       (&armsecurity.AlertsClient{}).NewListPager,
			NewFunc:        armsecurity.NewAlertsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alerts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "alerts_suppression_rules",
			Struct:         &armsecurity.AlertsSuppressionRule{},
			ResponseStruct: &armsecurity.AlertsSuppressionRulesClientListResponse{},
			Client:         &armsecurity.AlertsSuppressionRulesClient{},
			ListFunc:       (&armsecurity.AlertsSuppressionRulesClient{}).NewListPager,
			NewFunc:        armsecurity.NewAlertsSuppressionRulesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alertsSuppressionRules",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "allowed_connections",
			Struct:         &armsecurity.AllowedConnectionsResource{},
			ResponseStruct: &armsecurity.AllowedConnectionsClientListResponse{},
			Client:         &armsecurity.AllowedConnectionsClient{},
			ListFunc:       (&armsecurity.AllowedConnectionsClient{}).NewListPager,
			NewFunc:        armsecurity.NewAllowedConnectionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/allowedConnections",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "applications",
			Struct:         &armsecurity.Application{},
			ResponseStruct: &armsecurity.ApplicationsClientListResponse{},
			Client:         &armsecurity.ApplicationsClient{},
			ListFunc:       (&armsecurity.ApplicationsClient{}).NewListPager,
			NewFunc:        armsecurity.NewApplicationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/applications",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "assessments_metadata",
			Struct:         &armsecurity.AssessmentMetadataResponse{},
			ResponseStruct: &armsecurity.AssessmentsMetadataClientListResponse{},
			Client:         &armsecurity.AssessmentsMetadataClient{},
			ListFunc:       (&armsecurity.AssessmentsMetadataClient{}).NewListPager,
			NewFunc:        armsecurity.NewAssessmentsMetadataClient,
			URL:            "/providers/Microsoft.Security/assessmentMetadata",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "auto_provisioning_settings",
			Struct:         &armsecurity.AutoProvisioningSetting{},
			ResponseStruct: &armsecurity.AutoProvisioningSettingsClientListResponse{},
			Client:         &armsecurity.AutoProvisioningSettingsClient{},
			ListFunc:       (&armsecurity.AutoProvisioningSettingsClient{}).NewListPager,
			NewFunc:        armsecurity.NewAutoProvisioningSettingsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "automations",
			Struct:         &armsecurity.Automation{},
			ResponseStruct: &armsecurity.AutomationsClientListResponse{},
			Client:         &armsecurity.AutomationsClient{},
			ListFunc:       (&armsecurity.AutomationsClient{}).NewListPager,
			NewFunc:        armsecurity.NewAutomationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/automations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "connectors",
			Struct:         &armsecurity.Connector{},
			ResponseStruct: &armsecurity.ConnectorsClientListResponse{},
			Client:         &armsecurity.ConnectorsClient{},
			ListFunc:       (&armsecurity.ConnectorsClient{}).NewListPager,
			NewFunc:        armsecurity.NewConnectorsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityConnectors",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "discovered_security_solutions",
			Struct:         &armsecurity.DiscoveredSecuritySolution{},
			ResponseStruct: &armsecurity.DiscoveredSecuritySolutionsClientListResponse{},
			Client:         &armsecurity.DiscoveredSecuritySolutionsClient{},
			ListFunc:       (&armsecurity.DiscoveredSecuritySolutionsClient{}).NewListPager,
			NewFunc:        armsecurity.NewDiscoveredSecuritySolutionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/discoveredSecuritySolutions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "external_security_solutions",
			Struct:         &armsecurity.ExternalSecuritySolution{},
			ResponseStruct: &armsecurity.ExternalSecuritySolutionsClientListResponse{},
			Client:         &armsecurity.ExternalSecuritySolutionsClient{},
			ListFunc:       (&armsecurity.ExternalSecuritySolutionsClient{}).NewListPager,
			NewFunc:        armsecurity.NewExternalSecuritySolutionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/externalSecuritySolutions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "governance_rule",
			Struct:         &armsecurity.GovernanceRule{},
			ResponseStruct: &armsecurity.GovernanceRuleClientListResponse{},
			Client:         &armsecurity.GovernanceRuleClient{},
			ListFunc:       (&armsecurity.GovernanceRuleClient{}).NewListPager,
			NewFunc:        armsecurity.NewGovernanceRuleClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "jit_network_access_policies",
			Struct:         &armsecurity.JitNetworkAccessPolicy{},
			ResponseStruct: &armsecurity.JitNetworkAccessPoliciesClientListResponse{},
			Client:         &armsecurity.JitNetworkAccessPoliciesClient{},
			ListFunc:       (&armsecurity.JitNetworkAccessPoliciesClient{}).NewListPager,
			NewFunc:        armsecurity.NewJitNetworkAccessPoliciesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/jitNetworkAccessPolicies",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "locations",
			Struct:         &armsecurity.AscLocation{},
			ResponseStruct: &armsecurity.LocationsClientListResponse{},
			Client:         &armsecurity.LocationsClient{},
			ListFunc:       (&armsecurity.LocationsClient{}).NewListPager,
			NewFunc:        armsecurity.NewLocationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "regulatory_compliance_standards",
			Struct:         &armsecurity.RegulatoryComplianceStandard{},
			ResponseStruct: &armsecurity.RegulatoryComplianceStandardsClientListResponse{},
			Client:         &armsecurity.RegulatoryComplianceStandardsClient{},
			ListFunc:       (&armsecurity.RegulatoryComplianceStandardsClient{}).NewListPager,
			NewFunc:        armsecurity.NewRegulatoryComplianceStandardsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "secure_score_control_definitions",
			Struct:         &armsecurity.SecureScoreControlDefinitionItem{},
			ResponseStruct: &armsecurity.SecureScoreControlDefinitionsClientListResponse{},
			Client:         &armsecurity.SecureScoreControlDefinitionsClient{},
			ListFunc:       (&armsecurity.SecureScoreControlDefinitionsClient{}).NewListPager,
			NewFunc:        armsecurity.NewSecureScoreControlDefinitionsClient,
			URL:            "/providers/Microsoft.Security/secureScoreControlDefinitions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "secure_score_controls",
			Struct:         &armsecurity.SecureScoreControlDetails{},
			ResponseStruct: &armsecurity.SecureScoreControlsClientListResponse{},
			Client:         &armsecurity.SecureScoreControlsClient{},
			ListFunc:       (&armsecurity.SecureScoreControlsClient{}).NewListPager,
			NewFunc:        armsecurity.NewSecureScoreControlsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControls",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "secure_scores",
			Struct:         &armsecurity.SecureScoreItem{},
			ResponseStruct: &armsecurity.SecureScoresClientListResponse{},
			Client:         &armsecurity.SecureScoresClient{},
			ListFunc:       (&armsecurity.SecureScoresClient{}).NewListPager,
			NewFunc:        armsecurity.NewSecureScoresClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "solutions",
			Struct:         &armsecurity.Solution{},
			ResponseStruct: &armsecurity.SolutionsClientListResponse{},
			Client:         &armsecurity.SolutionsClient{},
			ListFunc:       (&armsecurity.SolutionsClient{}).NewListPager,
			NewFunc:        armsecurity.NewSolutionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "tasks",
			Struct:         &armsecurity.Task{},
			ResponseStruct: &armsecurity.TasksClientListResponse{},
			Client:         &armsecurity.TasksClient{},
			ListFunc:       (&armsecurity.TasksClient{}).NewListPager,
			NewFunc:        armsecurity.NewTasksClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/tasks",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "topology",
			Struct:         &armsecurity.TopologyResource{},
			ResponseStruct: &armsecurity.TopologyClientListResponse{},
			Client:         &armsecurity.TopologyClient{},
			ListFunc:       (&armsecurity.TopologyClient{}).NewListPager,
			NewFunc:        armsecurity.NewTopologyClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/topologies",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsecurity",
			Name:           "workspace_settings",
			Struct:         &armsecurity.WorkspaceSetting{},
			ResponseStruct: &armsecurity.WorkspaceSettingsClientListResponse{},
			Client:         &armsecurity.WorkspaceSettingsClient{},
			ListFunc:       (&armsecurity.WorkspaceSettingsClient{}).NewListPager,
			NewFunc:        armsecurity.NewWorkspaceSettingsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_security)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
