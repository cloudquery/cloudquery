package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/account"
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
	resources2 "github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/search"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/servicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/streamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/subscriptions"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/web"
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
			"account.locations":                    account.Locations(),
			"authorization.role_assignments":       authorization.AuthorizationRoleAssignments(),
			"authorization.role_definitions":       authorization.AuthorizationRoleDefinitions(),
			"batch.accounts":                       batch.BatchAccounts(),
			"cdn.profiles":                         cdn.Profiles(),
			"compute.disks":                        compute.ComputeDisks(),
			"compute.virtual_machines":             compute.ComputeVirtualMachines(),
			"compute.virtual_machine_scale_sets":   compute.VirtualMachineScaleSets(),
			"container.managed_clusters":           container.ContainerManagedClusters(),
			"container.registries":                 container.ContainerRegistries(),
			"cosmosdb.accounts":                    cosmosdb.CosmosDBAccounts(),
			"cosmosdb.sql_databases":               cosmosdb.CosmosDBSqlDatabases(),
			"cosmosdb.mongodb_databases":           cosmosdb.CosmosDBMongoDBDatabases(),
			"datalake.storage_accounts":            datalake.StorageAccounts(),
			"datalake.analytics_accounts":          datalake.AnalyticsAccounts(),
			"eventhub.namespaces":                  eventhub.EventHubNamespaces(),
			"frontdoor.front_doors":                frontdoor.FrontDoors(),
			"iothub.hubs":                          iothub.IothubHubs(),
			"keyvault.vaults":                      keyvault.KeyvaultVaults(),
			"keyvault.managed_hsm":                 keyvault.KeyvaultManagedHSM(),
			"logic.app_workflows":                  logic.LogicAppWorkflows(),
			"mariadb.servers":                      mariadb.MariadbServers(),
			"monitor.log_profiles":                 monitor.MonitorLogProfiles(),
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
			"servicebus.namespaces":                servicebus.Namespaces(),
			"sql.servers":                          sql.Servers(),
			"sql.managed_instances":                sql.ManagedInstances(),
			"storage.accounts":                     storage.StorageAccounts(),
			"streamanalytics.jobs":                 streamanalytics.StreamanalyticsJobs(),
			"subscriptions.subscriptions":          subscriptions.Subscriptions(),
			"subscriptions.tenants":                subscriptions.Tenants(),
			"web.apps":                             web.WebApps(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
