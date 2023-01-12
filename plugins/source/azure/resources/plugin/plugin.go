package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/authorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/azuredata"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cosmos"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/costmanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datafactory"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/healthbot"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/keyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/logic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/monitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/nginx"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/peering"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/redis"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/search"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/subscription"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

// var Tables []*schema.Table

var customTables = []*schema.Table{
	authorization.RoleDefinitions(),
	azuredata.SqlServerRegistrations(),
	cdn.Profiles(),
	compute.VirtualMachines(),
	compute.SKUs(),
	compute.VirtualMachineScaleSets(),
	cosmos.DatabaseAccounts(),
	costmanagement.Views(),
	datafactory.Factories(),
	healthbot.Bots(),
	keyvault.Keyvault(),
	keyvault.KeyvaultManagedHsms(),
	logic.Workflows(),
	mariadb.Servers(),
	monitor.TenantActivityLogAlerts(),
	mysql.Servers(),
	network.ExpressRouteGateways(),
	security.Pricings(),
	storage.Accounts(),
	subscription.Subscriptions(),
	redis.Caches(),
	resources.Resources(),
	subscription.Tenants(),
	search.Services(),
	nginx.Deployments(),
	peering.ServiceProviders(),
}

func Plugin() *source.Plugin {
	allTables := append(generatedTables(), customTables...)
	return source.NewPlugin(
		"azure",
		Version,
		allTables,
		client.New,
	)
}
