package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	"github.com/Azure/go-autorest/autorest"
)

type CosmosDBAccountClient interface {
	List(ctx context.Context) (result documentdb.DatabaseAccountsListResult, err error)
}

type CosmosDBSQLClient interface {
	ListSQLDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.SQLDatabaseListResult, err error)
}

type CosmosDBMongoDBClient interface {
	ListMongoDBDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.MongoDBDatabaseListResult, err error)
}

type CosmosDbClient struct {
	Accounts CosmosDBAccountClient
	SQL      CosmosDBSQLClient
	MongoDB  CosmosDBMongoDBClient
}

func NewCosmosDbClient(subscriptionId string, auth autorest.Authorizer) CosmosDbClient {
	accounts := documentdb.NewDatabaseAccountsClient(subscriptionId)
	accounts.Authorizer = auth

	sql := documentdb.NewSQLResourcesClient(subscriptionId)
	sql.Authorizer = auth

	mongo := documentdb.NewMongoDBResourcesClient(subscriptionId)
	mongo.Authorizer = auth

	return CosmosDbClient{
		Accounts: accounts,
		SQL:      sql,
		MongoDB:  mongo,
	}
}
