package provider

import (
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
			"authorization.role_assignments":       authorization.RoleAssignments(),
			"authorization.role_definitions":       authorization.RoleDefinitions(),
			"batch.accounts":                       batch.Accounts(),
			"cdn.profiles":                         cdn.Profiles(),
			"compute.disks":                        compute.Disks(),
			"compute.virtual_machine_scale_sets":   compute.VirtualMachineScaleSets(),
			"compute.virtual_machines":             compute.VirtualMachines(),
			"container.managed_clusters":           container.ManagedClusters(),
			"container.registries":                 container.Registries(),
			"cosmosdb.accounts":                    cosmosdb.Accounts(),
			"datalake.analytics_accounts":          datalake.AnalyticsAccounts(),
			"datalake.store_accounts":              datalake.StoreAccounts(),
			"eventhub.namespaces":                  eventhub.Namespaces(),
			"frontdoor.doors":                      frontdoor.Doors(),
			"iothub.devices":                       iothub.Devices(),
			"keyvault.managed_hs_ms":               keyvault.ManagedHSMs(),
			"keyvault.vaults":                      keyvault.Vaults(),
			"logic.workflows":                      logic.Workflows(),
			"mariadb.servers":                      mariadb.Servers(),
			"monitor.activity_log_alerts":          monitor.ActivityLogAlerts(),
			"monitor.activity_logs":                monitor.ActivityLogs(),
			"monitor.log_profiles":                 monitor.LogProfiles(),
			"monitor.resources":                    monitor.Resources(),
			"mysql.servers":                        mysql.Servers(),
			"network.express_route_circuits":       network.ExpressRouteCircuits(),
			"network.express_route_gateways":       network.ExpressRouteGateways(),
			"network.express_route_ports":          network.ExpressRoutePorts(),
			"network.interfaces":                   network.Interfaces(),
			"network.public_ip_addresses":          network.PublicIPAddresses(),
			"network.route_filters":                network.RouteFilters(),
			"network.route_tables":                 network.RouteTables(),
			"network.security_groups":              network.SecurityGroups(),
			"network.virtual_networks":             network.VirtualNetworks(),
			"network.watchers":                     network.Watchers(),
			"postgresql.servers":                   postgresql.Servers(),
			"redis.resource_types":                 redis.ResourceTypes(),
			"resources.groups":                     resources.Groups(),
			"resources.links":                      resources.Links(),
			"resources.policy_assignments":         resources.PolicyAssignments(),
			"search.services":                      search.Services(),
			"security.assessments":                 security.Assessments(),
			"security.auto_provisioning_settings":  security.AutoProvisioningSettings(),
			"security.contacts":                    security.Contacts(),
			"security.jit_network_access_policies": security.JitNetworkAccessPolicies(),
			"security.pricings":                    security.Pricings(),
			"security.settings":                    security.Settings(),
			"servicebus.namespaces":                servicebus.Namespaces(),
			"sql.managed_instances":                sql.ManagedInstances(),
			"sql.servers":                          sql.Servers(),
			"storage.accounts":                     storage.Accounts(),
			"streamanalytics.streaming_jobs":       streamanalytics.StreamingJobs(),
			"subscriptions.locations":              subscriptions.Locations(),
			"subscriptions.subscriptions":          subscriptions.Subscriptions(),
			"subscriptions.tenants":                subscriptions.Tenants(),
			"web.apps":                             web.Apps(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
