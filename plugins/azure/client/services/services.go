//go:generate mockgen -destination=./mocks/services.go -package=mocks  . DisksClient,GroupsClient,KeyClient,VaultClient,StorageAccountClient,StorageContainerClient,SqlServerClient,SqlDatabaseClient,MySQLServerClient,MySQLConfigurationClient,PostgresqlConfigurationClient,PostgresqlServerClient,VirtualNetworksClient
package services

import "github.com/Azure/go-autorest/autorest"

type Services struct {
	Compute    ComputeClient
	Resources  ResourcesClient
	KeyVault   KeyVaultClient
	Storage    StorageClient
	SQL        SQLClient
	PostgreSQL PostgreSQL
	MySQL      MySQL
	Network    NetworksClient
}

func InitServices(subscriptionId string, auth autorest.Authorizer) Services {
	return Services{
		Compute:    NewComputeClient(subscriptionId, auth),
		Resources:  NewResourcesClient(subscriptionId, auth),
		KeyVault:   NewKeyVaultClient(subscriptionId, auth),
		Storage:    NewStorageClient(subscriptionId, auth),
		SQL:        NewSQLClient(subscriptionId, auth),
		PostgreSQL: NewPostgresClient(subscriptionId, auth),
		MySQL:      NewMySQLClient(subscriptionId, auth),
		Network:    NewNetworksClient(subscriptionId, auth),
	}
}
