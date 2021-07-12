package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/Azure/go-autorest/autorest"
)

type SQLClient struct {
	Databases    SqlDatabaseClient
	Firewall     SQLFirewallClient
	ServerAdmins SQLServerAdminClient
	Servers      SqlServerClient
}

func NewSQLClient(subscriptionId string, auth autorest.Authorizer) SQLClient {
	servers := sql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	databases := sql.NewDatabasesClient(subscriptionId)
	databases.Authorizer = auth
	firewall := sql.NewFirewallRulesClient(subscriptionId)
	firewall.Authorizer = auth
	serverAdmins := sql.NewServerAzureADAdministratorsClient(subscriptionId)
	serverAdmins.Authorizer = auth
	return SQLClient{
		Databases:    databases,
		Firewall:     firewall,
		ServerAdmins: serverAdmins,
		Servers:      servers,
	}
}

type SqlServerClient interface {
	List(ctx context.Context) (result sql.ServerListResultPage, err error)
}

type SQLFirewallClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.FirewallRuleListResult, err error)
}

type SQLServerAdminClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.AdministratorListResultPage, err error)
}

type SqlDatabaseClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.DatabaseListResultPage, err error)
}
