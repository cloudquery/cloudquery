//go:generate mockgen -destination=./mocks/ad_applications.go -package=mocks . ADApplicationsClient
//go:generate mockgen -destination=./mocks/ad_groups.go -package=mocks . ADGroupsClient
//go:generate mockgen -destination=./mocks/ad_service_principals.go -package=mocks . ADServicePrinicpals
//go:generate mockgen -destination=./mocks/ad_users.go -package=mocks . ADUsersClient
//go:generate mockgen -destination=./mocks/authorization.go -package=mocks . RoleAssignmentsClient,RoleDefinitionsClient
//go:generate mockgen -destination=./mocks/compute.go -package=mocks . DisksClient,VirtualMachinesClient
//go:generate mockgen -destination=./mocks/containerservice.go -package=mocks . ManagedClustersClient
//go:generate mockgen -destination=./mocks/keyvault.go -package=mocks . KeyClient,SecretsClient,VaultClient
//go:generate mockgen -destination=./mocks/monitor.go -package=mocks . ActivityLogAlertsClient,LogProfilesClient,DiagnosticSettingsClient,ActivityLogClient
//go:generate mockgen -destination=./mocks/my_sql.go -package=mocks . MySQLServerClient,MySQLConfigurationClient
//go:generate mockgen -destination=./mocks/network.go -package=mocks . VirtualNetworksClient,SecurityGroupsClient,WatchersClient,PublicIPAddressesClient
//go:generate mockgen -destination=./mocks/postgresql.go -package=mocks . PostgresqlConfigurationClient,PostgresqlServerClient,PostgresqlFirewallRuleClient
//go:generate mockgen -destination=./mocks/resources.go -package=mocks . ResClient,GroupsClient
//go:generate mockgen -destination=./mocks/security.go -package=mocks . SecurityAutoProvisioningSettingsClient,SecurityContactsClient,SecurityPricingsClient,SecuritySettingsClient
//go:generate mockgen -destination=./mocks/sql.go -package=mocks . SQLDatabaseBlobAuditingPoliciesClient,SQLFirewallClient,SQLServerAdminClient,SQLServerBlobAuditingPolicies,SqlDatabaseClient,SqlServerClient
//go:generate mockgen -destination=./mocks/storage.go -package=mocks . StorageAccountClient,StorageContainerClient
//go:generate mockgen -destination=./mocks/subscriptions.go -package=mocks . SubscriptionGetter
package services

import "github.com/Azure/go-autorest/autorest"

type Services struct {
	AD            AD
	Authorization AuthorizationClient
	Compute       ComputeClient
	Container     ContainerServiceClient
	KeyVault      KeyVaultClient
	Monitor       MonitorClient
	MySQL         MySQL
	Network       NetworksClient
	PostgreSQL    PostgreSQL
	Resources     ResourcesClient
	Security      SecurityClient
	SQL           SQLClient
	Storage       StorageClient
	Subscriptions SubscriptionsClient
}

func InitServices(subscriptionId string, auth autorest.Authorizer) Services {
	return Services{
		AD:            NewADClient(subscriptionId, auth),
		Authorization: NewAuthorizationClient(subscriptionId, auth),
		Compute:       NewComputeClient(subscriptionId, auth),
		Container:     NewContainerServiceClient(subscriptionId, auth),
		KeyVault:      NewKeyVaultClient(subscriptionId, auth),
		Monitor:       NewMonitorClient(subscriptionId, auth),
		MySQL:         NewMySQLClient(subscriptionId, auth),
		Network:       NewNetworksClient(subscriptionId, auth),
		PostgreSQL:    NewPostgresClient(subscriptionId, auth),
		Resources:     NewResourcesClient(subscriptionId, auth),
		Security:      NewSecurityClient(subscriptionId, auth),
		SQL:           NewSQLClient(subscriptionId, auth),
		Storage:       NewStorageClient(subscriptionId, auth),
		Subscriptions: NewSubscriptionsClient(subscriptionId, auth),
	}
}
