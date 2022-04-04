//go:generate mockgen -destination=./mocks/ad_applications.go -package=mocks . ADApplicationsClient
//go:generate mockgen -destination=./mocks/ad_groups.go -package=mocks . ADGroupsClient
//go:generate mockgen -destination=./mocks/ad_service_principals.go -package=mocks . ADServicePrinicpals
//go:generate mockgen -destination=./mocks/ad_users.go -package=mocks . ADUsersClient
//go:generate mockgen -destination=./mocks/authorization.go -package=mocks . RoleAssignmentsClient,RoleDefinitionsClient
//go:generate mockgen -destination=./mocks/containerservice.go -package=mocks . ManagedClustersClient
//go:generate mockgen -destination=./mocks/eventhub.go -package=mocks . EventHubClient
//go:generate mockgen -destination=./mocks/keyvault.go -package=mocks . KeyVault71Client,VaultClient,KeyVaultManagedHSMClient,KeysClient
//go:generate mockgen -destination=./mocks/monitor.go -package=mocks . ActivityLogAlertsClient,LogProfilesClient,DiagnosticSettingsClient,ActivityLogClient
//go:generate mockgen -destination=./mocks/logic.go -package=mocks . MonitorDiagnosticSettingsClient,WorkflowsClient
//go:generate mockgen -destination=./mocks/mariadb.go -package=mocks . MariaDBConfigurationsClient,MariaDBServersClient
//go:generate mockgen -destination=./mocks/my_sql.go -package=mocks . MySQLServerClient,MySQLConfigurationClient
//go:generate mockgen -destination=./mocks/network.go -package=mocks . ExpressRouteCircuitsClient,ExpressRouteGatewaysClient,ExpressRoutePortsClient,InterfacesClient,PublicIPAddressesClient,RouteFiltersClient,SecurityGroupsClient,VirtualNetworksClient,WatchersClient
//go:generate mockgen -destination=./mocks/postgresql.go -package=mocks . PostgresqlConfigurationClient,PostgresqlServerClient,PostgresqlFirewallRuleClient
//go:generate mockgen -destination=./mocks/redis.go -package=mocks . RedisClient
//go:generate mockgen -destination=./mocks/resources.go -package=mocks . ResClient,GroupsClient,AssignmentsClient,LinksClient
//go:generate mockgen -destination=./mocks/servicebus.go -package=mocks . NamespacesClient
//go:generate mockgen -destination=./mocks/security.go -package=mocks . SecurityAutoProvisioningSettingsClient,SecurityContactsClient,SecurityPricingsClient,SecuritySettingsClient,JitNetworkAccessPoliciesClient,AssessmentsClient
//go:generate mockgen -destination=./mocks/streamanalytics.go -package=mocks . JobsClient
//go:generate mockgen -destination=./mocks/storage.go -package=mocks . StorageAccountClient,StorageBlobServicePropertiesClient,StorageBlobServicesClient,StorageContainerClient,StorageQueueServicePropertiesClient
//go:generate mockgen -destination=./mocks/subscriptions.go -package=mocks . SubscriptionGetter
//go:generate mockgen -destination=./mocks/web.go -package=mocks . AppsClient
//go:generate mockgen -destination=./mocks/cosmosdb.go -package=mocks . CosmosDBAccountClient,CosmosDBSQLClient,CosmosDBMongoDBClient
//go:generate mockgen -destination=./mocks/iothub.go -package=mocks . IotHubClient
//go:generate mockgen -destination=./mocks/batch.go -package=mocks . BatchAccountClient
//go:generate mockgen -destination=./mocks/search.go -package=mocks . SearchServiceClient
package services

import "github.com/Azure/go-autorest/autorest"

type Services struct {
	AD                AD
	Authorization     AuthorizationClient
	Batch             BatchClient
	Compute           ComputeClient
	ContainerService  ContainerServiceClient
	ContainerRegistry ContainerRegistryClient
	CosmosDb          CosmosDbClient
	DataLake          DataLakeClient
	EventHub          EventHubClient
	IotHub            IotHubClient
	KeyVault          KeyVaultClient
	Logic             LogicClient
	MariaDB           MariaDB
	Monitor           MonitorClient
	MySQL             MySQL
	Network           NetworksClient
	PostgreSQL        PostgreSQL
	Redis             RedisClient
	Resources         ResourcesClient
	Search            SearchClient
	Servicebus        ServicebusClient
	Security          SecurityClient
	SQL               SQLClient
	Storage           StorageClient
	StreamAnalytics   StreamAnalyticsClient
	Subscriptions     Subscriptions
	Web               WebClient
}

func InitServices(subscriptionId string, auth autorest.Authorizer) (Services, error) {
	keyVault, err := NewKeyVaultClient(subscriptionId, auth)
	if err != nil {
		return Services{}, err
	}
	return Services{
		AD:                NewADClient(subscriptionId, auth),
		Authorization:     NewAuthorizationClient(subscriptionId, auth),
		Batch:             NewBatchClient(subscriptionId, auth),
		Compute:           NewComputeClient(subscriptionId, auth),
		ContainerService:  NewContainerServiceClient(subscriptionId, auth),
		ContainerRegistry: NewContainerRegistryClient(subscriptionId, auth),
		CosmosDb:          NewCosmosDbClient(subscriptionId, auth),
		DataLake:          NewDataLakeClient(subscriptionId, auth),
		EventHub:          NewEventHubClient(subscriptionId, auth),
		IotHub:            NewIotHubClient(subscriptionId, auth),
		Logic:             NewLogicClient(subscriptionId, auth),
		KeyVault:          keyVault,
		MariaDB:           NewMariaDBClient(subscriptionId, auth),
		Monitor:           NewMonitorClient(subscriptionId, auth),
		MySQL:             NewMySQLClient(subscriptionId, auth),
		Network:           NewNetworksClient(subscriptionId, auth),
		PostgreSQL:        NewPostgresClient(subscriptionId, auth),
		Redis:             NewRedisClient(subscriptionId, auth),
		Resources:         NewResourcesClient(subscriptionId, auth),
		Search:            NewSearchClient(subscriptionId, auth),
		Security:          NewSecurityClient(subscriptionId, auth),
		Servicebus:        NewServicebusClient(subscriptionId, auth),
		SQL:               NewSQLClient(subscriptionId, auth),
		Storage:           NewStorageClient(subscriptionId, auth),
		StreamAnalytics:   NewStreamAnalyticsClient(subscriptionId, auth),
		Subscriptions:     NewSubscriptionsClient(subscriptionId, auth),
		Web:               NewWebClient(subscriptionId, auth),
	}, nil
}
