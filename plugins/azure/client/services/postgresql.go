package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/Azure/go-autorest/autorest"
)

type PostgreSQL struct {
	Servers       PostgresqlServerClient
	Configuration PostgresqlConfigurationClient
	FirewallRule  PostgresqlFirewallRuleClient
}

func NewPostgresClient(subscriptionId string, auth autorest.Authorizer) PostgreSQL {
	servers := postgresql.NewServersClient(subscriptionId)
	servers.Authorizer = auth

	confSvc := postgresql.NewConfigurationsClient(subscriptionId)
	confSvc.Authorizer = auth

	firewallSvc := postgresql.NewFirewallRulesClient(subscriptionId)
	firewallSvc.Authorizer = auth
	return PostgreSQL{
		Servers:       servers,
		Configuration: confSvc,
		FirewallRule:  firewallSvc,
	}
}

type PostgresqlServerClient interface {
	List(ctx context.Context) (result postgresql.ServerListResult, err error)
}

type PostgresqlConfigurationClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ConfigurationListResult, err error)
}

type PostgresqlFirewallRuleClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.FirewallRuleListResult, err error)
}
