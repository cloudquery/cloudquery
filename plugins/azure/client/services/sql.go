package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/Azure/go-autorest/autorest"
)

type SQLClient struct {
	DatabaseBlobAuditingPolicies     SQLDatabaseBlobAuditingPoliciesClient
	Databases                        SQLDatabaseClient
	DatabaseThreatDetectionPolicies  SQLDatabaseThreatDetectionPoliciesClient
	DatabaseVulnerabilityAssessments SQLDatabaseVulnerabilityAssessmentsClient
	Firewall                         SQLFirewallClient
	ServerAdmins                     SQLServerAdminClient
	ServerBlobAuditingPolicies       SQLServerBlobAuditingPolicies
	ServerDevOpsAuditSettings        SQLServerDevOpsAuditSettingsClient
	Servers                          SQLServerClient
	ServerVulnerabilityAssessments   SQLServerVulnerabilityAssessmentsClient
	TransparentDataEncryptions       TransparentDataEncryptionsClient
	EncryptionProtectors             EncryptionProtectorsClient
	VirtualNetworkRules              SQLVirtualNetworkRulesClient
	ServerSecurityAlertPolicies      ServerSecurityAlertPoliciesClient
}

func NewSQLClient(subscriptionId string, auth autorest.Authorizer) SQLClient {
	databases := sql.NewDatabasesClient(subscriptionId)
	databases.Authorizer = auth
	dbap := sql.NewDatabaseBlobAuditingPoliciesClient(subscriptionId)
	dbap.Authorizer = auth
	dtdp := sql.NewDatabaseThreatDetectionPoliciesClient(subscriptionId)
	dtdp.Authorizer = auth
	dva := sql.NewDatabaseVulnerabilityAssessmentsClient(subscriptionId)
	dva.Authorizer = auth
	firewall := sql.NewFirewallRulesClient(subscriptionId)
	firewall.Authorizer = auth
	sbap := sql.NewServerBlobAuditingPoliciesClient(subscriptionId)
	sbap.Authorizer = auth
	sdas := sql.NewServerDevOpsAuditSettingsClient(subscriptionId)
	sdas.Authorizer = auth
	serverAdmins := sql.NewServerAzureADAdministratorsClient(subscriptionId)
	serverAdmins.Authorizer = auth
	servers := sql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	sva := sql.NewServerVulnerabilityAssessmentsClient(subscriptionId)
	sva.Authorizer = auth
	enc := sql.NewTransparentDataEncryptionsClient(subscriptionId)
	enc.Authorizer = auth
	ep := sql.NewEncryptionProtectorsClient(subscriptionId)
	ep.Authorizer = auth
	vnr := sql.NewVirtualNetworkRulesClient(subscriptionId)
	vnr.Authorizer = auth
	ssap := sql.NewServerSecurityAlertPoliciesClient(subscriptionId)
	ssap.Authorizer = auth
	return SQLClient{
		DatabaseBlobAuditingPolicies:     dbap,
		Databases:                        databases,
		DatabaseThreatDetectionPolicies:  dtdp,
		DatabaseVulnerabilityAssessments: dva,
		Firewall:                         firewall,
		ServerAdmins:                     serverAdmins,
		ServerBlobAuditingPolicies:       sbap,
		ServerDevOpsAuditSettings:        sdas,
		Servers:                          servers,
		ServerVulnerabilityAssessments:   sva,
		TransparentDataEncryptions:       enc,
		EncryptionProtectors:             ep,
		VirtualNetworkRules:              vnr,
		ServerSecurityAlertPolicies:      ssap,
	}
}

type SQLServerClient interface {
	List(ctx context.Context) (result sql.ServerListResultPage, err error)
}

type SQLFirewallClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.FirewallRuleListResult, err error)
}

type SQLServerAdminClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.AdministratorListResultPage, err error)
}

type SQLServerBlobAuditingPolicies interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerBlobAuditingPolicyListResultPage, err error)
}

type SQLServerDevOpsAuditSettingsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerDevOpsAuditSettingsListResultPage, err error)
}

type SQLServerVulnerabilityAssessmentsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerVulnerabilityAssessmentListResultPage, err error)
}
type SQLDatabaseClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.DatabaseListResultPage, err error)
}

type SQLDatabaseBlobAuditingPoliciesClient interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseBlobAuditingPolicyListResultPage, err error)
}

type SQLDatabaseThreatDetectionPoliciesClient interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseSecurityAlertPolicy, err error)
}

type SQLDatabaseVulnerabilityAssessmentsClient interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseVulnerabilityAssessmentListResultPage, err error)
}

type TransparentDataEncryptionsClient interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.TransparentDataEncryption, err error)
}

type EncryptionProtectorsClient interface {
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.EncryptionProtector, err error)
}

type SQLVirtualNetworkRulesClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.VirtualNetworkRuleListResultPage, err error)
}

type ServerSecurityAlertPoliciesClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.LogicalServerSecurityAlertPolicyListResultPage, err error)
}
