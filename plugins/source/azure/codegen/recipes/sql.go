package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func SQL() []Resource {
	var sqlDatabaseRelations = []resourceDefinition{
		{
			azureStruct:      &sql.DatabaseBlobAuditingPolicy{},
			listFunction:     "ListByDatabase",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Parent.Item.(sql.Server)",
				"database := parent.Item.(sql.Database)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.DatabaseVulnerabilityAssessment{},
			listFunction:     "ListByDatabase",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Parent.Item.(sql.Server)",
				"database := parent.Item.(sql.Database)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.VulnerabilityAssessmentScanRecord{},
			listFunction:     "ListByDatabase",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Parent.Item.(sql.Server)",
				"database := parent.Item.(sql.Database)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			subServiceOverride:       "DatabaseVulnerabilityAssessmentScans",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.BackupLongTermRetentionPolicy{},
			listFunction:     "ListByDatabase",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Parent.Item.(sql.Server)",
				"database := parent.Item.(sql.Database)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			listHandler: `if err != nil {
				return err
			}
			res <- response`,
			subServiceOverride:       "BackupLongTermRetentionPolicies",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
			mockListResult:           mockDirectResponse,
		},
		{
			azureStruct:      &sql.DatabaseSecurityAlertPolicy{},
			listFunction:     "Get",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Parent.Item.(sql.Server)",
				"database := parent.Item.(sql.Database)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			listHandler: `if err != nil {
				return err
			}
			res <- response`,
			subServiceOverride:       "DatabaseThreatDetectionPolicies",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
			mockListResult:           mockDirectResponse,
		},
		{
			azureStruct:      &sql.TransparentDataEncryption{},
			listFunction:     "Get",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Parent.Item.(sql.Server)",
				"database := parent.Item.(sql.Database)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			listHandler: `if err != nil {
				return err
			}
			res <- response`,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
			mockListResult:           mockDirectResponse,
		},
	}

	var firewallRuleResource = resourceDefinition{
		azureStruct:      &sql.FirewallRule{},
		listFunction:     "ListByServer",
		listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
		listFunctionArgsInit: []string{
			"server := parent.Item.(sql.Server)",
			`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
		listHandler:              valueHandler,
		mockListFunctionArgsInit: []string{""},
		mockListFunctionArgs:     []string{`"test"`, `"test"`},
	}
	var sqlServerRelations = []resourceDefinition{
		firewallRuleResource,
		{
			azureStruct:      &sql.Database{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			relations:                sqlDatabaseRelations,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.EncryptionProtector{},
			listFunction:     "Get",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			listHandler: `if err != nil {
				return err
			}
			res <- response`,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           mockDirectResponse,
		},
		{
			azureStruct:      &sql.VirtualNetworkRule{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.ServerAzureADAdministrator{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			subServiceOverride:       "ServerAdmins",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           "AdministratorListResult",
		},
		{
			azureStruct:      &sql.ServerBlobAuditingPolicy{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.ServerDevOpsAuditingSettings{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           "ServerDevOpsAuditSettingsListResult",
		},
		{
			azureStruct:      &sql.ServerVulnerabilityAssessment{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.ServerSecurityAlertPolicy{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{
				"server := parent.Item.(sql.Server)",
				`resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           "LogicalServerSecurityAlertPolicyListResult",
		},
	}

	var managedDatabaseRelations = []resourceDefinition{
		{
			azureStruct:      &sql.DatabaseVulnerabilityAssessment{},
			listFunction:     "ListByDatabase",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*instance.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"instance := parent.Parent.Item.(sql.ManagedInstance)",
				"database := parent.Item.(sql.ManagedDatabase)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			subServiceOverride:       "ManagedDatabaseVulnerabilityAssessments",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.VulnerabilityAssessmentScanRecord{},
			listFunction:     "ListByDatabase",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*instance.Name", "*database.Name"},
			listFunctionArgsInit: []string{
				"instance := parent.Parent.Item.(sql.ManagedInstance)",
				"database := parent.Item.(sql.ManagedDatabase)",
				`resourceDetails, err := client.ParseResourceID(*database.ID)
			if err != nil {
				return err
			}`},
			subServiceOverride:       "ManagedDatabaseVulnerabilityAssessmentScans",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
	}

	var managedInstanceRelations = []resourceDefinition{
		{
			azureStruct:      &sql.ManagedDatabase{},
			listFunction:     "ListByInstance",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*instance.Name"},
			listFunctionArgsInit: []string{"instance := parent.Item.(sql.ManagedInstance)", `resourceDetails, err := client.ParseResourceID(*instance.ID)
			if err != nil {
				return err
			}`},
			relations:                managedDatabaseRelations,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.ManagedInstanceVulnerabilityAssessment{},
			listFunction:     "ListByInstance",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*instance.Name"},
			listFunctionArgsInit: []string{
				"instance := parent.Item.(sql.ManagedInstance)",
				`resourceDetails, err := client.ParseResourceID(*instance.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:      &sql.ManagedInstanceEncryptionProtector{},
			listFunction:     "ListByInstance",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*instance.Name"},
			listFunctionArgsInit: []string{
				"instance := parent.Item.(sql.ManagedInstance)",
				`resourceDetails, err := client.ParseResourceID(*instance.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
	}

	var topLevelResources = []resourceDefinition{
		{
			azureStruct:  &sql.Server{},
			listFunction: "List",
			relations:    sqlServerRelations,
		},
		{
			azureStruct:  &sql.ManagedInstance{},
			listFunction: "List",
			relations:    managedInstanceRelations,
		}}

	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"},
				},
			},
			definitions:         topLevelResources,
			serviceNameOverride: "SQL",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"},
				},
			},
			serviceNameOverride: "SQL",
		},
	}

	initParents(resourcesByTemplates)

	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, sqlServerRelations...)
	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, sqlDatabaseRelations...)
	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, managedInstanceRelations...)
	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, managedDatabaseRelations...)
	resourcesByTemplates[1].definitions = []resourceDefinition{sqlServerRelations[0]}

	return generateResources(resourcesByTemplates)
}
