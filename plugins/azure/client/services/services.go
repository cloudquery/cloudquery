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
