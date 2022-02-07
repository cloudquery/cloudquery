package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSQLServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockSQLServerClient(ctrl)
	databaseSvc := mocks.NewMockSQLDatabaseClient(ctrl)
	firewallSvc := mocks.NewMockSQLFirewallClient(ctrl)
	adminsSvc := mocks.NewMockSQLServerAdminClient(ctrl)
	databaseBlobSvc := mocks.NewMockSQLDatabaseBlobAuditingPoliciesClient(ctrl)
	serverBlobSvc := mocks.NewMockSQLServerBlobAuditingPolicies(ctrl)
	devopsAuditSvc := mocks.NewMockSQLServerDevOpsAuditSettingsClient(ctrl)
	databaseThreatsSvc := mocks.NewMockSQLDatabaseThreatDetectionPoliciesClient(ctrl)
	serverVulnsSvc := mocks.NewMockSQLServerVulnerabilityAssessmentsClient(ctrl)
	dbVulnsSvc := mocks.NewMockSQLDatabaseVulnerabilityAssessmentsClient(ctrl)
	encSvc := mocks.NewMockTransparentDataEncryptionsClient(ctrl)
	epSvc := mocks.NewMockEncryptionProtectorsClient(ctrl)
	vnrSvc := mocks.NewMockSQLVirtualNetworkRulesClient(ctrl)
	ssapSvc := mocks.NewMockServerSecurityAlertPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			DatabaseBlobAuditingPolicies:     databaseBlobSvc,
			Databases:                        databaseSvc,
			DatabaseThreatDetectionPolicies:  databaseThreatsSvc,
			DatabaseVulnerabilityAssessments: dbVulnsSvc,
			Firewall:                         firewallSvc,
			ServerAdmins:                     adminsSvc,
			ServerBlobAuditingPolicies:       serverBlobSvc,
			ServerDevOpsAuditSettings:        devopsAuditSvc,
			Servers:                          serverSvc,
			ServerVulnerabilityAssessments:   serverVulnsSvc,
			TransparentDataEncryptions:       encSvc,
			EncryptionProtectors:             epSvc,
			VirtualNetworkRules:              vnrSvc,
			ServerSecurityAlertPolicies:      ssapSvc,
		},
	}
	server := sql.Server{}
	if err := faker.FakeData(&server); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	rg := client.FakeResourceGroup
	server.ID = &rg
	serverSvc.EXPECT().List(gomock.Any()).Return(
		sql.NewServerListResultPage(
			sql.ServerListResult{Value: &[]sql.Server{server}},
			func(context.Context, sql.ServerListResult) (sql.ServerListResult, error) {
				return sql.ServerListResult{}, nil
			},
		),
		nil,
	)

	database := sql.Database{}
	if err := faker.FakeData(&database); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	database.ID = &rg
	databaseSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewDatabaseListResultPage(
			sql.DatabaseListResult{Value: &[]sql.Database{database}},
			func(context.Context, sql.DatabaseListResult) (sql.DatabaseListResult, error) {
				return sql.DatabaseListResult{}, nil
			},
		), nil,
	)

	var databaseBlobPolicy sql.DatabaseBlobAuditingPolicy
	if err := faker.FakeData(&databaseBlobPolicy); err != nil {
		t.Fatal(err)
	}
	databaseBlobSvc.EXPECT().ListByDatabase(gomock.Any(), "test", *server.Name, *database.Name).Return(
		sql.NewDatabaseBlobAuditingPolicyListResultPage(
			sql.DatabaseBlobAuditingPolicyListResult{Value: &[]sql.DatabaseBlobAuditingPolicy{databaseBlobPolicy}},
			func(context.Context, sql.DatabaseBlobAuditingPolicyListResult) (sql.DatabaseBlobAuditingPolicyListResult, error) {
				return sql.DatabaseBlobAuditingPolicyListResult{}, nil
			},
		), nil,
	)

	var rule sql.FirewallRule
	if err := faker.FakeData(&rule); err != nil {
		t.Fatal(err)
	}
	firewallSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.FirewallRuleListResult{Value: &[]sql.FirewallRule{rule}}, nil,
	)

	var admin sql.ServerAzureADAdministrator
	if err := faker.FakeData(&admin); err != nil {
		t.Fatal(err)
	}
	adminPage := sql.NewAdministratorListResultPage(
		sql.AdministratorListResult{Value: &[]sql.ServerAzureADAdministrator{admin}},
		func(context.Context, sql.AdministratorListResult) (sql.AdministratorListResult, error) {
			return sql.AdministratorListResult{}, nil
		},
	)
	adminsSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(adminPage, nil)

	var serverBlobPolicy sql.ServerBlobAuditingPolicy
	if err := faker.FakeData(&serverBlobPolicy); err != nil {
		t.Fatal(err)
	}
	serverBlobSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewServerBlobAuditingPolicyListResultPage(
			sql.ServerBlobAuditingPolicyListResult{Value: &[]sql.ServerBlobAuditingPolicy{serverBlobPolicy}},
			func(context.Context, sql.ServerBlobAuditingPolicyListResult) (sql.ServerBlobAuditingPolicyListResult, error) {
				return sql.ServerBlobAuditingPolicyListResult{}, nil
			},
		), nil,
	)

	var devopsAuditSettings sql.ServerDevOpsAuditingSettings
	if err := faker.FakeData(&devopsAuditSettings); err != nil {
		t.Fatal(err)
	}
	devopsAuditSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewServerDevOpsAuditSettingsListResultPage(
			sql.ServerDevOpsAuditSettingsListResult{Value: &[]sql.ServerDevOpsAuditingSettings{devopsAuditSettings}},
			func(context.Context, sql.ServerDevOpsAuditSettingsListResult) (sql.ServerDevOpsAuditSettingsListResult, error) {
				return sql.ServerDevOpsAuditSettingsListResult{}, nil
			},
		), nil,
	)

	var databaseAlert sql.DatabaseSecurityAlertPolicy
	if err := faker.FakeData(&databaseAlert); err != nil {
		t.Fatal(err)
	}
	databaseThreatsSvc.EXPECT().Get(gomock.Any(), "test", *server.Name, *database.Name).Return(databaseAlert, nil)

	var serverVuln sql.ServerVulnerabilityAssessment
	if err := faker.FakeData(&serverVuln); err != nil {
		t.Fatal(err)
	}
	serverVulnsSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewServerVulnerabilityAssessmentListResultPage(
			sql.ServerVulnerabilityAssessmentListResult{Value: &[]sql.ServerVulnerabilityAssessment{serverVuln}},
			func(context.Context, sql.ServerVulnerabilityAssessmentListResult) (sql.ServerVulnerabilityAssessmentListResult, error) {
				return sql.ServerVulnerabilityAssessmentListResult{}, nil
			},
		), nil,
	)

	var dbVuln sql.DatabaseVulnerabilityAssessment
	if err := faker.FakeData(&dbVuln); err != nil {
		t.Fatal(err)
	}
	dbVulnsSvc.EXPECT().ListByDatabase(gomock.Any(), "test", *server.Name, *database.Name).Return(
		sql.NewDatabaseVulnerabilityAssessmentListResultPage(
			sql.DatabaseVulnerabilityAssessmentListResult{Value: &[]sql.DatabaseVulnerabilityAssessment{dbVuln}},
			func(context.Context, sql.DatabaseVulnerabilityAssessmentListResult) (sql.DatabaseVulnerabilityAssessmentListResult, error) {
				return sql.DatabaseVulnerabilityAssessmentListResult{}, nil
			},
		), nil,
	)

	var e sql.TransparentDataEncryption
	if err := faker.FakeData(&e); err != nil {
		t.Fatal(err)
	}
	encSvc.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		e, nil,
	)

	var ep sql.EncryptionProtector
	if err := faker.FakeData(&ep); err != nil {
		t.Fatal(err)
	}
	epSvc.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		ep, nil,
	)

	var vnr sql.VirtualNetworkRule
	if err := faker.FakeData(&vnr); err != nil {
		t.Fatal(err)
	}
	vnrSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewVirtualNetworkRuleListResultPage(
			sql.VirtualNetworkRuleListResult{Value: &[]sql.VirtualNetworkRule{vnr}},
			func(context.Context, sql.VirtualNetworkRuleListResult) (sql.VirtualNetworkRuleListResult, error) {
				return sql.VirtualNetworkRuleListResult{}, nil
			},
		), nil,
	)

	var ssap sql.ServerSecurityAlertPolicy
	if err := faker.FakeData(&ssap); err != nil {
		t.Fatal(err)
	}
	ssapSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewLogicalServerSecurityAlertPolicyListResultPage(
			sql.LogicalServerSecurityAlertPolicyListResult{Value: &[]sql.ServerSecurityAlertPolicy{ssap}},
			func(ctx context.Context, result sql.LogicalServerSecurityAlertPolicyListResult) (sql.LogicalServerSecurityAlertPolicyListResult, error) {
				return sql.LogicalServerSecurityAlertPolicyListResult{}, nil
			},
		), nil,
	)

	return s
}

func TestSQLServers(t *testing.T) {
	client.AzureMockTestHelper(t, SQLServers(), buildSQLServerMock, client.TestOptions{})
}
