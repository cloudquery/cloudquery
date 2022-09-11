package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func SQL() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &sql.Server{},
					listFunction: "List",
					relations: []string{
						"databases()",
						"encryptionProtectors()",
						"virtualNetworkRules()",
						"firewallRules()",
						"serverAdmins()",
						"serverBlobAuditingPolicies()",
						"serverDevOpsAuditingSettings()",
						"serverVulnerabilityAssessments()",
						"serverSecurityAlertPolicies()",
					},
				},
				{
					azureStruct:  &sql.ManagedInstance{},
					listFunction: "List",
					relations:    []string{"managedDatabases()", "managedInstanceVulnerabilityAssessments()", "managedInstanceEncryptionProtectors()"},
				},
				// relations of sql.ManagedInstance
				{
					azureStruct:      &sql.ManagedDatabase{},
					listFunction:     "ListByInstance",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*instance.Name"},
					listFunctionArgsInit: []string{"instance := parent.Item.(sql.ManagedInstance)", `resourceDetails, err := client.ParseResourceID(*instance.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					relations:                []string{"managedDatabaseVulnerabilityAssessments()", "managedDatabaseVulnerabilityAssessmentScans()"},
					isRelation:               true,
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
						return errors.WithStack(err)
					}`},
					isRelation:               true,
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
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				// relations of sql.ManagedDatabase
				{
					azureStruct:      &sql.DatabaseVulnerabilityAssessment{},
					listFunction:     "ListByDatabase",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*instance.Name", "*database.Name"},
					listFunctionArgsInit: []string{
						"instance := parent.Parent.Item.(sql.ManagedInstance)",
						"database := parent.Item.(sql.ManagedDatabase)",
						`resourceDetails, err := client.ParseResourceID(*database.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					subServiceOverride:       "ManagedDatabaseVulnerabilityAssessments",
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					subServiceOverride:       "ManagedDatabaseVulnerabilityAssessmentScans",
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				// relations of sql.Server
				{
					azureStruct:      &sql.Database{},
					listFunction:     "ListByServer",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{
						"server := parent.Item.(sql.Server)",
						`resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					relations: []string{
						"databaseBlobAuditingPolicies()",
						"databaseVulnerabilityAssessments()",
						"databaseVulnerabilityAssessmentScans()",
						"backupLongTermRetentionPolicies()",
						"databaseThreatDetectionPolicies()",
						"transparentDataEncryptions()",
					},
					isRelation:               true,
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
						return errors.WithStack(err)
					}`},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:      &sql.VirtualNetworkRule{},
					listFunction:     "ListByServer",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{
						"server := parent.Item.(sql.Server)",
						`resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:      &sql.FirewallRule{},
					listFunction:     "ListByServer",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{
						"server := parent.Item.(sql.Server)",
						`resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listHandler:              valueHandler,
					isRelation:               true,
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
						return errors.WithStack(err)
					}`},
					subServiceOverride:       "ServerAdmins",
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:      &sql.ServerBlobAuditingPolicy{},
					listFunction:     "ListByServer",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{
						"server := parent.Item.(sql.Server)",
						`resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					isRelation:               true,
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
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:      &sql.ServerVulnerabilityAssessment{},
					listFunction:     "ListByServer",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{
						"server := parent.Item.(sql.Server)",
						`resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					isRelation:               true,
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
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				// relations of sql.Database
				{
					azureStruct:      &sql.DatabaseBlobAuditingPolicy{},
					listFunction:     "ListByDatabase",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name", "*database.Name"},
					listFunctionArgsInit: []string{
						"server := parent.Parent.Item.(sql.Server)",
						"database := parent.Item.(sql.Database)",
						`resourceDetails, err := client.ParseResourceID(*database.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
						return errors.WithStack(err)
					}`},
					subServiceOverride:       "DatabaseVulnerabilityAssessmentScans",
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
						return errors.WithStack(err)
					}`},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					subServiceOverride:       "BackupLongTermRetentionPolicies",
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
						return errors.WithStack(err)
					}`},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					subServiceOverride:       "DatabaseThreatDetectionPolicies",
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
						return errors.WithStack(err)
					}`},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
			},
			serviceNameOverride: "SQL",
		},
	}

	return generateResources(resourcesByTemplates)
}
