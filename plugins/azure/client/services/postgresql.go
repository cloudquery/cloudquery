package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/Azure/go-autorest/autorest"
)

type PostgreSQL struct {
	Servers       PostgresqlServerClient
	Configuration PostgresqlConfigurationClient
}

func NewPostgresClient(subscriptionId string, auth autorest.Authorizer) PostgreSQL {
	servers := postgresql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	return PostgreSQL{
		Servers: servers,
	}
}

type PostgresqlServerClient interface {
	List(ctx context.Context) (result postgresql.ServerListResult, err error)
}

type PostgresqlConfigurationClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ConfigurationListResult, err error)
}
