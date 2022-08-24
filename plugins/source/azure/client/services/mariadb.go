//go:generate mockgen -destination=./mocks/mariadb.go -package=mocks . ConfigurationsClient,ServersClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
	"github.com/Azure/go-autorest/autorest"
)

type MariaDBClient struct {
	Configurations ConfigurationsClient
	Servers        ServersClient
}

type ServersClient interface {
	List(ctx context.Context) (result mariadb.ServerListResult, err error)
}

type ConfigurationsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ConfigurationListResult, err error)
}

func NewMariaDBClient(subscriptionId string, auth autorest.Authorizer) MariaDBClient {
	configs := mariadb.NewConfigurationsClient(subscriptionId)
	configs.Authorizer = auth
	servers := mariadb.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	return MariaDBClient{
		Configurations: configs,
		Servers:        servers,
	}
}
