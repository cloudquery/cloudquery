package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/Azure/go-autorest/autorest"
)

type SQLClient struct {
	BackupLongTermRetentionPolicies             SQLBackupLongTermRetentionPoliciesClient
	DatabaseBlobAuditingPolicies                SQLDatabaseBlobAuditingPoliciesClient
	Databases                                   SQLDatabasesClient
	DatabaseVulnerabilityAssessmentScans        SQLDatabaseVulnerabilityAssessmentScansClient
	DatabaseThreatDetectionPolicies             SQLDatabaseThreatDetectionPoliciesClient
	DatabaseVulnerabilityAssessments            SQLDatabaseVulnerabilityAssessmentsClient
	FirewallRules                               SQLFirewallRulesClient
	ServerAdmins                                SQLServerAdminsClient
	ServerBlobAuditingPolicies                  SQLServerBlobAuditingPoliciesClient
	ServerDevOpsAuditingSettings                SQLServerDevOpsAuditingSettingsClient
	Servers                                     SQLServersClient
	ServerVulnerabilityAssessments              SQLServerVulnerabilityAssessmentsClient
	TransparentDataEncryptions                  SQLTransparentDataEncryptionsClient
	EncryptionProtectors                        SQLEncryptionProtectorsClient
	ManagedInstances                            SQLManagedInstancesClient
	ManagedInstanceVulnerabilityAssessments     SQLManagedInstanceVulnerabilityAssessmentsClient
	ManagedDatabases                            SQLManagedDatabasesClient
	ManagedDatabaseVulnerabilityAssessments     SQLManagedDatabaseVulnerabilityAssessmentsClient
	ManagedDatabaseVulnerabilityAssessmentScans SQLManagedDatabaseVulnerabilityAssessmentScansClient
	ManagedInstanceEncryptionProtectors         SQLManagedInstanceEncryptionProtectorsClient
	VirtualNetworkRules                         SQLVirtualNetworkRulesClient
	ServerSecurityAlertPolicies                 SQLServerSecurityAlertPoliciesClient
}

//go:generate mockgen -destination=./mocks/sql.go -package=mocks . SQLServersClient,SQLFirewallRulesClient,SQLServerAdminsClient,SQLServerBlobAuditingPoliciesClient,SQLServerDevOpsAuditingSettingsClient,SQLServerVulnerabilityAssessmentsClient,SQLEncryptionProtectorsClient,SQLVirtualNetworkRulesClient,SQLServerSecurityAlertPoliciesClient,SQLDatabasesClient,SQLDatabaseBlobAuditingPoliciesClient,SQLDatabaseThreatDetectionPoliciesClient,SQLDatabaseVulnerabilityAssessmentsClient,SQLDatabaseVulnerabilityAssessmentScansClient,SQLTransparentDataEncryptionsClient,SQLBackupLongTermRetentionPoliciesClient,SQLManagedInstancesClient,SQLManagedInstanceVulnerabilityAssessmentsClient,SQLManagedInstanceEncryptionProtectorsClient,SQLManagedDatabasesClient,SQLManagedDatabaseVulnerabilityAssessmentsClient,SQLManagedDatabaseVulnerabilityAssessmentScansClient
type SQLServersClient interface {
	List(ctx context.Context) (result sql.ServerListResultPage, err error)
}

type SQLFirewallRulesClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.FirewallRuleListResult, err error)
}

type SQLServerAdminsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.AdministratorListResultPage, err error)
}

type SQLServerBlobAuditingPoliciesClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerBlobAuditingPolicyListResultPage, err error)
}

type SQLServerDevOpsAuditingSettingsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerDevOpsAuditSettingsListResultPage, err error)
}

type SQLServerVulnerabilityAssessmentsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerVulnerabilityAssessmentListResultPage, err error)
}

type SQLEncryptionProtectorsClient interface {
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.EncryptionProtector, err error)
}

type SQLVirtualNetworkRulesClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.VirtualNetworkRuleListResultPage, err error)
}

type SQLServerSecurityAlertPoliciesClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.LogicalServerSecurityAlertPolicyListResultPage, err error)
}

type SQLDatabasesClient interface {
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

type SQLDatabaseVulnerabilityAssessmentScansClient interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.VulnerabilityAssessmentScanRecordListResultPage, err error)
}

type SQLTransparentDataEncryptionsClient interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.TransparentDataEncryption, err error)
}

type SQLBackupLongTermRetentionPoliciesClient interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.BackupLongTermRetentionPolicy, err error)
}

type SQLManagedInstancesClient interface {
	List(ctx context.Context) (result sql.ManagedInstanceListResultPage, err error)
}

type SQLManagedInstanceVulnerabilityAssessmentsClient interface {
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceVulnerabilityAssessmentListResultPage, err error)
}

type SQLManagedInstanceEncryptionProtectorsClient interface {
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceEncryptionProtectorListResultPage, err error)
}

type SQLManagedDatabasesClient interface {
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedDatabaseListResultPage, err error)
}

type SQLManagedDatabaseVulnerabilityAssessmentsClient interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.DatabaseVulnerabilityAssessmentListResultPage, err error)
}

type SQLManagedDatabaseVulnerabilityAssessmentScansClient interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.VulnerabilityAssessmentScanRecordListResultPage, err error)
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
	dvas := sql.NewDatabaseVulnerabilityAssessmentScansClient(subscriptionId)
	dvas.Authorizer = auth
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
	mi := sql.NewManagedInstancesClient(subscriptionId)
	mi.Authorizer = auth
	miva := sql.NewManagedInstanceVulnerabilityAssessmentsClient(subscriptionId)
	miva.Authorizer = auth
	miep := sql.NewManagedInstanceEncryptionProtectorsClient(subscriptionId)
	miep.Authorizer = auth
	md := sql.NewManagedDatabasesClient(subscriptionId)
	md.Authorizer = auth
	mdva := sql.NewManagedDatabaseVulnerabilityAssessmentsClient(subscriptionId)
	mdva.Authorizer = auth
	mdvas := sql.NewManagedDatabaseVulnerabilityAssessmentScansClient(subscriptionId)
	mdvas.Authorizer = auth
	vnr := sql.NewVirtualNetworkRulesClient(subscriptionId)
	vnr.Authorizer = auth
	ssap := sql.NewServerSecurityAlertPoliciesClient(subscriptionId)
	ssap.Authorizer = auth
	bltrp := sql.NewBackupLongTermRetentionPoliciesClient(subscriptionId)
	bltrp.Authorizer = auth
	return SQLClient{
		BackupLongTermRetentionPolicies:             bltrp,
		DatabaseBlobAuditingPolicies:                dbap,
		Databases:                                   databases,
		DatabaseThreatDetectionPolicies:             dtdp,
		DatabaseVulnerabilityAssessments:            dva,
		DatabaseVulnerabilityAssessmentScans:        dvas,
		FirewallRules:                               firewall,
		ServerAdmins:                                serverAdmins,
		ServerBlobAuditingPolicies:                  sbap,
		ServerDevOpsAuditingSettings:                sdas,
		Servers:                                     servers,
		ServerVulnerabilityAssessments:              sva,
		TransparentDataEncryptions:                  enc,
		EncryptionProtectors:                        ep,
		ManagedInstances:                            mi,
		ManagedInstanceVulnerabilityAssessments:     miva,
		ManagedInstanceEncryptionProtectors:         miep,
		ManagedDatabases:                            md,
		ManagedDatabaseVulnerabilityAssessments:     mdva,
		ManagedDatabaseVulnerabilityAssessmentScans: mdvas,
		VirtualNetworkRules:                         vnr,
		ServerSecurityAlertPolicies:                 ssap,
	}
}
