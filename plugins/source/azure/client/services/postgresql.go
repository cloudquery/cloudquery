//go:generate mockgen -destination=./mocks/postgresql.go -package=mocks . PostgreSQLConfigurationsClient,PostgreSQLServersClient,PostgreSQLFirewallRulesClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/Azure/go-autorest/autorest"
)

type PostgreSQLClient struct {
	Servers        PostgreSQLServersClient
	Configurations PostgreSQLConfigurationsClient
	FirewallRules  PostgreSQLFirewallRulesClient
}

type PostgreSQLServersClient interface {
	List(ctx context.Context) (result postgresql.ServerListResult, err error)
}

type PostgreSQLConfigurationsClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ConfigurationListResult, err error)
}

type PostgreSQLFirewallRulesClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.FirewallRuleListResult, err error)
}

func NewPostgresClient(subscriptionId string, auth autorest.Authorizer) PostgreSQLClient {
	servers := postgresql.NewServersClient(subscriptionId)
	servers.Authorizer = auth

	confSvc := postgresql.NewConfigurationsClient(subscriptionId)
	confSvc.Authorizer = auth

	firewallSvc := postgresql.NewFirewallRulesClient(subscriptionId)
	firewallSvc.Authorizer = auth
	return PostgreSQLClient{
		Servers:        servers,
		Configurations: confSvc,
		FirewallRules:  firewallSvc,
	}
}
