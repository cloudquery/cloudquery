//go:generate mockgen -destination=./mocks/services.go -package=mocks  . DisksClient,GroupsClient,KeyClient,VaultClient,StorageAccountClient,StorageContainerClient,SqlServerClient,SqlDatabaseClient,MySQLServerClient,MySQLConfigurationClient,PostgresqlConfigurationClient,PostgresqlServerClient,VirtualNetworksClient
//go:generate mockgen -destination=./mocks/ad_users.go -package=mocks . ADUsersClient
package services

import "github.com/Azure/go-autorest/autorest"

type Services struct {
	Compute    ComputeClient
	KeyVault   KeyVaultClient
	MySQL      MySQL
	Network    NetworksClient
	PostgreSQL PostgreSQL
	AD         AD
	Resources  ResourcesClient
	SQL        SQLClient
	Storage    StorageClient
}

func InitServices(subscriptionId string, auth autorest.Authorizer) Services {
	return Services{
		Compute:    NewComputeClient(subscriptionId, auth),
		KeyVault:   NewKeyVaultClient(subscriptionId, auth),
		MySQL:      NewMySQLClient(subscriptionId, auth),
		Network:    NewNetworksClient(subscriptionId, auth),
		PostgreSQL: NewPostgresClient(subscriptionId, auth),
		AD:         NewADClient(subscriptionId, auth),
		Resources:  NewResourcesClient(subscriptionId, auth),
		SQL:        NewSQLClient(subscriptionId, auth),
		Storage:    NewStorageClient(subscriptionId, auth),
	}
}
