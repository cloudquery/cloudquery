//go:generate mockgen -destination=./mocks/ad_applications.go -package=mocks . ADApplicationsClient
//go:generate mockgen -destination=./mocks/ad_groups.go -package=mocks . ADGroupsClient
//go:generate mockgen -destination=./mocks/ad_service_principals.go -package=mocks . ADServicePrinicpals
//go:generate mockgen -destination=./mocks/ad_users.go -package=mocks . ADUsersClient
//go:generate mockgen -destination=./mocks/authorization.go -package=mocks . RoleAssignmentsClient,RoleDefinitionsClient
//go:generate mockgen -destination=./mocks/containerservice.go -package=mocks . ManagedClustersClient
//go:generate mockgen -destination=./mocks/eventhub.go -package=mocks . EventHubClient
//go:generate mockgen -destination=./mocks/keyvault.go -package=mocks . KeyVault71Client,VaultClient,KeyVaultManagedHSMClient
//go:generate mockgen -destination=./mocks/monitor.go -package=mocks . ActivityLogAlertsClient,LogProfilesClient,DiagnosticSettingsClient,ActivityLogClient
//go:generate mockgen -destination=./mocks/mariadb.go -package=mocks . MariaDBConfigurationsClient,MariaDBServersClient
//go:generate mockgen -destination=./mocks/my_sql.go -package=mocks . MySQLServerClient,MySQLConfigurationClient
//go:generate mockgen -destination=./mocks/network.go -package=mocks . InterfacesClient,VirtualNetworksClient,SecurityGroupsClient,WatchersClient,PublicIPAddressesClient
//go:generate mockgen -destination=./mocks/postgresql.go -package=mocks . PostgresqlConfigurationClient,PostgresqlServerClient,PostgresqlFirewallRuleClient
//go:generate mockgen -destination=./mocks/redis.go -package=mocks . RedisClient
//go:generate mockgen -destination=./mocks/resources.go -package=mocks . ResClient,GroupsClient,AssignmentsClient,LinksClient
//go:generate mockgen -destination=./mocks/servicebus.go -package=mocks . NamespacesClient
//go:generate mockgen -destination=./mocks/security.go -package=mocks . SecurityAutoProvisioningSettingsClient,SecurityContactsClient,SecurityPricingsClient,SecuritySettingsClient,JitNetworkAccessPoliciesClient
//go:generate mockgen -destination=./mocks/storage.go -package=mocks . StorageAccountClient,StorageBlobServicePropertiesClient,StorageBlobServicesClient,StorageContainerClient,StorageQueueServicePropertiesClient
//go:generate mockgen -destination=./mocks/subscriptions.go -package=mocks . SubscriptionGetter
//go:generate mockgen -destination=./mocks/web.go -package=mocks . AppsClient
//go:generate mockgen -destination=./mocks/cosmosdb.go -package=mocks . CosmosDBAccountClient,CosmosDBSQLClient,CosmosDBMongoDBClient
//go:generate mockgen -destination=./mocks/iothub.go -package=mocks . IotHubClient
package services

import "github.com/Azure/go-autorest/autorest"

type Services struct {
	AD                AD
	Authorization     AuthorizationClient
	Compute           ComputeClient
	ContainerService  ContainerServiceClient
	ContainerRegistry ContainerRegistryClient
	CosmosDb          CosmosDbClient
	DataLake          DataLakeClient
	EventHub          EventHubClient
	IotHub            IotHubClient
	KeyVault          KeyVaultClient
	MariaDB           MariaDB
	Monitor           MonitorClient
	MySQL             MySQL
	Network           NetworksClient
	PostgreSQL        PostgreSQL
	Redis             RedisClient
	Resources         ResourcesClient
	Servicebus        ServicebusClient
	Security          SecurityClient
	SQL               SQLClient
	Storage           StorageClient
	Subscriptions     SubscriptionsClient
	Web               WebClient
}

func InitServices(subscriptionId string, auth autorest.Authorizer) Services {
	return Services{
		AD:                NewADClient(subscriptionId, auth),
		Authorization:     NewAuthorizationClient(subscriptionId, auth),
		Compute:           NewComputeClient(subscriptionId, auth),
		ContainerService:  NewContainerServiceClient(subscriptionId, auth),
		ContainerRegistry: NewContainerRegistryClient(subscriptionId, auth),
		CosmosDb:          NewCosmosDbClient(subscriptionId, auth),
		DataLake:          NewDataLakeClient(subscriptionId, auth),
		EventHub:          NewEventHubClient(subscriptionId, auth),
		IotHub:            NewIotHubClient(subscriptionId, auth),
		KeyVault:          NewKeyVaultClient(subscriptionId, auth),
		MariaDB:           NewMariaDBClient(subscriptionId, auth),
		Monitor:           NewMonitorClient(subscriptionId, auth),
		MySQL:             NewMySQLClient(subscriptionId, auth),
		Network:           NewNetworksClient(subscriptionId, auth),
		PostgreSQL:        NewPostgresClient(subscriptionId, auth),
		Redis:             NewRedisClient(subscriptionId, auth),
		Resources:         NewResourcesClient(subscriptionId, auth),
		Security:          NewSecurityClient(subscriptionId, auth),
		Servicebus:        NewServicebusClient(subscriptionId, auth),
		SQL:               NewSQLClient(subscriptionId, auth),
		Storage:           NewStorageClient(subscriptionId, auth),
		Subscriptions:     NewSubscriptionsClient(subscriptionId, auth),
		Web:               NewWebClient(subscriptionId, auth),
	}
}
