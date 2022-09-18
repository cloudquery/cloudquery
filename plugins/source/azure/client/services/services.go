package services

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/go-autorest/autorest"
)

type Services struct {
	Authorization   AuthorizationClient
	Batch           BatchClient
	CDN             CDNClient
	Compute         ComputeClient
	Container       ContainerClient
	CosmosDB        CosmosDBClient
	DataLake        DataLakeClient
	EventHub        EventHubClient
	FrontDoor       FrontDoorClient
	IotHub          IotHubClient
	KeyVault        KeyVaultClient
	Logic           LogicClient
	MariaDB         MariaDBClient
	Monitor         MonitorClient
	MySQL           MySQLClient
	Network         NetworkClient
	PostgreSQL      PostgreSQLClient
	Redis           RedisClient
	Resources       ResourcesClient
	Search          SearchClient
	Servicebus      ServicebusClient
	Security        SecurityClient
	SQL             SQLClient
	Storage         StorageClient
	StreamAnalytics StreamAnalyticsClient
	Subscriptions   SubscriptionsClient
	Web             WebClient
}

func InitServices(subscriptionId string, auth autorest.Authorizer, azCred azcore.TokenCredential) (Services, error) {
	keyVault, err := NewKeyVaultClient(subscriptionId, auth)
	if err != nil {
		return Services{}, err
	}

	subscriptionsClient, err := NewSubscriptionsClient(subscriptionId, auth, azCred)
	if err != nil {
		return Services{}, err
	}

	return Services{
		Authorization:   NewAuthorizationClient(subscriptionId, auth),
		Batch:           NewBatchClient(subscriptionId, auth),
		CDN:             NewCDNClient(subscriptionId, auth),
		Compute:         NewComputeClient(subscriptionId, auth),
		Container:       NewContainerClient(subscriptionId, auth),
		CosmosDB:        NewCosmosDbClient(subscriptionId, auth),
		DataLake:        NewDataLakeClient(subscriptionId, auth),
		EventHub:        NewEventHubClient(subscriptionId, auth),
		FrontDoor:       NewFrontDoorClient(subscriptionId, auth),
		IotHub:          NewIotHubClient(subscriptionId, auth),
		Logic:           NewLogicClient(subscriptionId, auth),
		KeyVault:        keyVault,
		MariaDB:         NewMariaDBClient(subscriptionId, auth),
		Monitor:         NewMonitorClient(subscriptionId, auth),
		MySQL:           NewMySQLClient(subscriptionId, auth),
		Network:         NewNetworksClient(subscriptionId, auth),
		PostgreSQL:      NewPostgresClient(subscriptionId, auth),
		Redis:           NewRedisClient(subscriptionId, auth),
		Resources:       NewResourcesClient(subscriptionId, auth),
		Search:          NewSearchClient(subscriptionId, auth),
		Security:        NewSecurityClient(subscriptionId, auth),
		Servicebus:      NewServicebusClient(subscriptionId, auth),
		SQL:             NewSQLClient(subscriptionId, auth),
		Storage:         NewStorageClient(subscriptionId, auth),
		StreamAnalytics: NewStreamAnalyticsClient(subscriptionId, auth),
		Subscriptions:   subscriptionsClient,
		Web:             NewWebClient(subscriptionId, auth),
	}, nil
}
