package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/Azure/go-autorest/autorest"
)

type MySQL struct {
	Servers       MySQLServerClient
	Configuration MySQLConfigurationClient
}

func NewMySQLClient(subscriptionId string, auth autorest.Authorizer) MySQL {
	servers := mysql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	return MySQL{
		Servers: servers,
	}
}

type MySQLServerClient interface {
	List(ctx context.Context) (result mysql.ServerListResult, err error)
}

type MySQLConfigurationClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ConfigurationListResult, err error)
}
