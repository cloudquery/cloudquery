package plugin

import (
	_ "embed"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/authorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/container"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cosmosdb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datalake"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/eventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/frontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/iothub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/keyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/logic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/monitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/redis"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/search"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/servicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/streamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/subscriptions"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/web"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
	//go:embed example.yml
	exampleConfig string
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"azure",
		Version,
		[]*schema.Table{
			authorization.RoleAssignments(),
			authorization.RoleDefinitions(),
			batch.Accounts(),
			cdn.Profiles(),
			compute.Disks(),
			compute.VirtualMachineScaleSets(),
			compute.VirtualMachines(),
			container.ManagedClusters(),
			container.Registries(),
			cosmosdb.Accounts(),
			datalake.AnalyticsAccounts(),
			datalake.StoreAccounts(),
			eventhub.Namespaces(),
			frontdoor.Doors(),
			iothub.Devices(),
			keyvault.ManagedHsms(),
			keyvault.Vaults(),
			logic.Workflows(),
			mariadb.Servers(),
			monitor.ActivityLogAlerts(),
			monitor.ActivityLogs(),
			monitor.LogProfiles(),
			monitor.Resources(),
			mysql.Servers(),
			network.ExpressRouteCircuits(),
			network.ExpressRouteGateways(),
			network.ExpressRoutePorts(),
			network.Interfaces(),
			network.PublicIPAddresses(),
			network.RouteFilters(),
			network.RouteTables(),
			network.SecurityGroups(),
			network.VirtualNetworks(),
			network.Watchers(),
			postgresql.Servers(),
			redis.Caches(),
			resources.Groups(),
			resources.Links(),
			resources.PolicyAssignments(),
			search.Services(),
			security.Assessments(),
			security.AutoProvisioningSettings(),
			security.Contacts(),
			security.JitNetworkAccessPolicies(),
			security.Pricings(),
			security.Settings(),
			servicebus.Namespaces(),
			sql.ManagedInstances(),
			sql.Servers(),
			storage.Accounts(),
			streamanalytics.StreamingJobs(),
			subscriptions.Locations(),
			subscriptions.Subscriptions(),
			subscriptions.Tenants(),
			web.Apps(),
		},
		client.New,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
