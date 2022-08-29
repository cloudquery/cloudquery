//go:generate mockgen -destination=./mocks/cosmosdb.go -package=mocks . CosmosDBAccountsClient,CosmosDBSQLClient,CosmosDBMongoDBClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	"github.com/Azure/go-autorest/autorest"
)

type CosmosDBAccountsClient interface {
	List(ctx context.Context) (result documentdb.DatabaseAccountsListResult, err error)
}

type CosmosDBSQLClient interface {
	ListSQLDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.SQLDatabaseListResult, err error)
}

type CosmosDBMongoDBClient interface {
	ListMongoDBDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.MongoDBDatabaseListResult, err error)
}

type CosmosDBClient struct {
	Accounts CosmosDBAccountsClient
	SQL      CosmosDBSQLClient
	MongoDB  CosmosDBMongoDBClient
}

func NewCosmosDbClient(subscriptionId string, auth autorest.Authorizer) CosmosDBClient {
	accounts := documentdb.NewDatabaseAccountsClient(subscriptionId)
	accounts.Authorizer = auth

	sql := documentdb.NewSQLResourcesClient(subscriptionId)
	sql.Authorizer = auth

	mongo := documentdb.NewMongoDBResourcesClient(subscriptionId)
	mongo.Authorizer = auth

	return CosmosDBClient{
		Accounts: accounts,
		SQL:      sql,
		MongoDB:  mongo,
	}
}
