//go:generate mockgen -destination=./mocks/my_sql.go -package=mocks . MySQLServersClient,MySQLConfigurationsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/Azure/go-autorest/autorest"
)

type MySQLClient struct {
	Servers        MySQLServersClient
	Configurations MySQLConfigurationsClient
}

type MySQLServersClient interface {
	List(ctx context.Context) (result mysql.ServerListResult, err error)
}

type MySQLConfigurationsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ConfigurationListResult, err error)
}

func NewMySQLClient(subscriptionId string, auth autorest.Authorizer) MySQLClient {
	servers := mysql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	conf := mysql.NewConfigurationsClient(subscriptionId)
	conf.Authorizer = auth
	return MySQLClient{
		Servers:        servers,
		Configurations: conf,
	}
}
