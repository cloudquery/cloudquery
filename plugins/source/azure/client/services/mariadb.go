//go:generate mockgen -destination=./mocks/mariadb.go -package=mocks . MariaDBConfigurationsClient,MariaDBServersClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
	"github.com/Azure/go-autorest/autorest"
)

type MariaDB struct {
	Configurations MariaDBConfigurationsClient
	Servers        MariaDBServersClient
}

type MariaDBServersClient interface {
	List(ctx context.Context) (result mariadb.ServerListResult, err error)
}

type MariaDBConfigurationsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ConfigurationListResult, err error)
}

func NewMariaDBClient(subscriptionId string, auth autorest.Authorizer) MariaDB {
	configs := mariadb.NewConfigurationsClient(subscriptionId)
	configs.Authorizer = auth
	servers := mariadb.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	return MariaDB{
		Configurations: configs,
		Servers:        servers,
	}
}
