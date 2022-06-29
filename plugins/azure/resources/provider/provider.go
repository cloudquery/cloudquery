package provider

import (
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/resources/services/account"
	"github.com/cloudquery/cq-provider-azure/resources/services/authorization"
	"github.com/cloudquery/cq-provider-azure/resources/services/batch"
	"github.com/cloudquery/cq-provider-azure/resources/services/compute"
	"github.com/cloudquery/cq-provider-azure/resources/services/container"
	"github.com/cloudquery/cq-provider-azure/resources/services/cosmosdb"
	"github.com/cloudquery/cq-provider-azure/resources/services/datalake"
	"github.com/cloudquery/cq-provider-azure/resources/services/eventhub"
	"github.com/cloudquery/cq-provider-azure/resources/services/iothub"
	"github.com/cloudquery/cq-provider-azure/resources/services/keyvault"
	"github.com/cloudquery/cq-provider-azure/resources/services/logic"
	"github.com/cloudquery/cq-provider-azure/resources/services/mariadb"
	"github.com/cloudquery/cq-provider-azure/resources/services/monitor"
	"github.com/cloudquery/cq-provider-azure/resources/services/mysql"
	"github.com/cloudquery/cq-provider-azure/resources/services/network"
	"github.com/cloudquery/cq-provider-azure/resources/services/postgresql"
	"github.com/cloudquery/cq-provider-azure/resources/services/redis"
	resources2 "github.com/cloudquery/cq-provider-azure/resources/services/resources"
	"github.com/cloudquery/cq-provider-azure/resources/services/search"
	"github.com/cloudquery/cq-provider-azure/resources/services/security"
	"github.com/cloudquery/cq-provider-azure/resources/services/servicebus"
	"github.com/cloudquery/cq-provider-azure/resources/services/sql"
	"github.com/cloudquery/cq-provider-azure/resources/services/storage"
	"github.com/cloudquery/cq-provider-azure/resources/services/streamanalytics"
	"github.com/cloudquery/cq-provider-azure/resources/services/subscription"
	"github.com/cloudquery/cq-provider-azure/resources/services/web"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:         Version,
		Name:            "azure",
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
			"account.locations":                  account.AccountLocations(),
			"authorization.role_assignments":     authorization.AuthorizationRoleAssignments(),
			"authorization.role_definitions":     authorization.AuthorizationRoleDefinitions(),
			"batch.accounts":                     batch.BatchAccounts(),
			"compute.disks":                      compute.ComputeDisks(),
			"compute.virtual_machines":           compute.ComputeVirtualMachines(),
			"compute.virtual_machine_scale_sets": compute.VirtualMachineScaleSets(),
			"container.managed_clusters":         container.ContainerManagedClusters(),
			"container.registries":               container.ContainerRegistries(),
			"cosmosdb.accounts":                  cosmosdb.CosmosDBAccounts(),
			"cosmosdb.sql_databases":             cosmosdb.CosmosDBSqlDatabases(),
			"cosmosdb.mongodb_databases":         cosmosdb.CosmosDBMongoDBDatabases(),
			"datalake.storage_accounts":          datalake.StorageAccounts(),
			"datalake.analytics_accounts":        datalake.AnalyticsAccounts(),
			"eventhub.namespaces":                eventhub.EventHubNamespaces(),
			"iothub.hubs":                        iothub.IothubHubs(),
			// This resource is currently not working
			// https://github.com/cloudquery/cq-provider-azure/issues/107
			"keyvault.vaults":      keyvault.KeyvaultVaults(),
			"keyvault.managed_hsm": keyvault.KeyvaultManagedHSM(),
			"logic.app_workflows":  logic.LogicAppWorkflows(),
			"mariadb.servers":      mariadb.MariadbServers(),
			"monitor.log_profiles": monitor.MonitorLogProfiles(),
			// This resource is currently not working
			"monitor.diagnostic_settings":          monitor.MonitorDiagnosticSettings(),
			"monitor.activity_logs":                monitor.MonitorActivityLogs(),
			"monitor.activity_log_alerts":          monitor.MonitorActivityLogAlerts(),
			"mysql.servers":                        mysql.MySQLServers(),
			"network.express_route_circuits":       network.NetworkExpressRouteCircuits(),
			"network.express_route_gateways":       network.NetworkExpressRouteGateways(),
			"network.express_route_ports":          network.NetworkExpressRoutePorts(),
			"network.interfaces":                   network.NetworkInterfaces(),
			"network.public_ip_addresses":          network.NetworkPublicIPAddresses(),
			"network.route_filters":                network.NetworkRouteFilters(),
			"network.route_tables":                 network.NetworkRouteTables(),
			"network.security_groups":              network.NetworkSecurityGroups(),
			"network.virtual_networks":             network.NetworkVirtualNetworks(),
			"network.watchers":                     network.NetworkWatchers(),
			"postgresql.servers":                   postgresql.PostgresqlServers(),
			"redis.services":                       redis.RedisServices(),
			"resources.groups":                     resources2.ResourcesGroups(),
			"resources.policy_assignments":         resources2.ResourcesPolicyAssignments(),
			"resources.links":                      resources2.ResourcesLinks(),
			"security.assessments":                 security.SecurityAssessments(),
			"search.services":                      search.SearchServices(),
			"security.auto_provisioning_settings":  security.SecurityAutoProvisioningSettings(),
			"security.contacts":                    security.SecurityContacts(),
			"security.pricings":                    security.SecurityPricings(),
			"security.settings":                    security.SecuritySettings(),
			"security.jit_network_access_policies": security.SecurityJitNetworkAccessPolicies(),
			"servicebus.namespaces":                servicebus.ServicebusNamespaces(),
			"sql.servers":                          sql.SQLServers(),
			"sql.managed_instances":                sql.SqlManagedInstances(),
			"storage.accounts":                     storage.StorageAccounts(),
			"streamanalytics.jobs":                 streamanalytics.StreamanalyticsJobs(),
			"subscription.subscriptions":           subscription.SubscriptionSubscriptions(),
			"web.apps":                             web.WebApps(),
		},
		Config: func(f cqproto.ConfigFormat) provider.Config {
			return client.NewConfig(f)
		},
	}
}
