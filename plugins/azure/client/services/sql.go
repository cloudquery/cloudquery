package services

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/Azure/go-autorest/autorest"
)

type SQLClient struct {
	Servers  SqlServerClient
	Database SqlDatabaseClient
}

func NewSQLClient(subscriptionId string, auth autorest.Authorizer) SQLClient {
	servers := sql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	database := sql.NewDatabasesClient(subscriptionId)
	database.Authorizer = auth
	return SQLClient{
		Servers:  servers,
		Database: database,
	}
}

type SqlServerClient interface {
	List(ctx context.Context) (result sql.ServerListResult, err error)
}

type SqlDatabaseClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, expand string, filter string) (result sql.DatabaseListResult, err error)
}
