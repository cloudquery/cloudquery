//go:generate mockgen -destination=./mocks/my_sql.go -package=mocks . MySQLServerClient,MySQLConfigurationClient
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

type MySQLServerClient interface {
	List(ctx context.Context) (result mysql.ServerListResult, err error)
}

type MySQLConfigurationClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ConfigurationListResult, err error)
}

func NewMySQLClient(subscriptionId string, auth autorest.Authorizer) MySQL {
	servers := mysql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	conf := mysql.NewConfigurationsClient(subscriptionId)
	conf.Authorizer = auth
	return MySQL{
		Servers:       servers,
		Configuration: conf,
	}
}
