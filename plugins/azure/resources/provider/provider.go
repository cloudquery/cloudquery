package provider

import (
	"embed"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/resources/services/authorization"
	"github.com/cloudquery/cq-provider-azure/resources/services/compute"
	"github.com/cloudquery/cq-provider-azure/resources/services/container"
	"github.com/cloudquery/cq-provider-azure/resources/services/cosmosdb"
	"github.com/cloudquery/cq-provider-azure/resources/services/eventhub"
	"github.com/cloudquery/cq-provider-azure/resources/services/keyvault"
	"github.com/cloudquery/cq-provider-azure/resources/services/monitor"
	"github.com/cloudquery/cq-provider-azure/resources/services/mysql"
	"github.com/cloudquery/cq-provider-azure/resources/services/network"
	"github.com/cloudquery/cq-provider-azure/resources/services/postgresql"
	"github.com/cloudquery/cq-provider-azure/resources/services/redis"
	resources2 "github.com/cloudquery/cq-provider-azure/resources/services/resources"
	"github.com/cloudquery/cq-provider-azure/resources/services/security"
	"github.com/cloudquery/cq-provider-azure/resources/services/sql"
	"github.com/cloudquery/cq-provider-azure/resources/services/storage"
	"github.com/cloudquery/cq-provider-azure/resources/services/subscription"
	"github.com/cloudquery/cq-provider-azure/resources/services/web"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*/*.sql
	azureMigrations embed.FS
	Version         = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:         Version,
		Name:            "azure",
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		Migrations:      azureMigrations,
		ResourceMap: map[string]*schema.Table{
			"authorization.role_assignments":     authorization.AuthorizationRoleAssignments(),
			"authorization.role_definitions":     authorization.AuthorizationRoleDefinitions(),
			"compute.disks":                      compute.ComputeDisks(),
			"compute.virtual_machines":           compute.ComputeVirtualMachines(),
			"compute.virtual_machine_scale_sets": compute.VirtualMachineScaleSets(),
			"container.managed_clusters":         container.ContainerManagedClusters(),
			"cosmosdb.accounts":                  cosmosdb.CosmosDBAccounts(),
			"cosmosdb.sql_databases":             cosmosdb.CosmosDBSqlDatabases(),
			"cosmosdb.mongodb_databases":         cosmosdb.CosmosDBMongoDBDatabases(),
			"eventhub.namespaces":                eventhub.EventHubNamespaces(),
			// This resource is currently not working
			// https://github.com/cloudquery/cq-provider-azure/issues/107
			"keyvault.vaults":      keyvault.KeyvaultVaults(),
			"keyvault.managed_hsm": keyvault.KeyvaultManagedHSM(),
			"monitor.log_profiles": monitor.MonitorLogProfiles(),
			// This resource is currently not working
			"monitor.diagnostic_settings":          monitor.MonitorDiagnosticSettings(),
			"monitor.activity_logs":                monitor.MonitorActivityLogs(),
			"monitor.activity_log_alerts":          monitor.MonitorActivityLogAlerts(),
			"mysql.servers":                        mysql.MySQLServers(),
			"network.virtual_networks":             network.NetworkVirtualNetworks(),
			"network.security_groups":              network.NetworkSecurityGroups(),
			"network.public_ip_addresses":          network.NetworkPublicIPAddresses(),
			"network.watchers":                     network.NetworkWatchers(),
			"postgresql.servers":                   postgresql.PostgresqlServers(),
			"redis.services":                       redis.RedisServices(),
			"resources.groups":                     resources2.ResourcesGroups(),
			"resources.policy_assignments":         resources2.ResourcesPolicyAssignments(),
			"resources.links":                      resources2.ResourcesLinks(),
			"security.auto_provisioning_settings":  security.SecurityAutoProvisioningSettings(),
			"security.contacts":                    security.SecurityContacts(),
			"security.pricings":                    security.SecurityPricings(),
			"security.settings":                    security.SecuritySettings(),
			"security.jit_network_access_policies": security.SecurityJitNetworkAccessPolicies(),
			"sql.servers":                          sql.SQLServers(),
			"sql.managed_instances":                sql.SqlManagedInstances(),
			"storage.accounts":                     storage.StorageAccounts(),
			"subscription.subscriptions":           subscription.SubscriptionSubscriptions(),
			"web.apps":                             web.WebApps(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
