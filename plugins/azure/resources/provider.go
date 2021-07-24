package resources

import (
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "azure",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"ad.applications":                     AdApplications(),
			"ad.groups":                           AdGroups(),
			"ad.service_principals":               AdServicePrincipals(),
			"ad.users":                            AdUsers(),
			"authorization.role_assignments":      AuthorizationRoleAssignments(),
			"authorization.role_definitions":      AuthorizationRoleDefinitions(),
			"compute.disks":                       ComputeDisks(),
			"compute.virtual_machines":            ComputeVirtualMachines(),
			"container.managed_clusters":          ContainerManagedClusters(),
			"keyvault.vaults":                     KeyvaultVaults(),
			"monitor.log_profiles":                MonitorLogProfiles(),
			"monitor.diagnostic_settings":         MonitorDiagnosticSettings(),
			"monitor.activity_logs":               MonitorActivityLogs(),
			"monitor.activity_log_alerts":         MonitorActivityLogAlerts(),
			"mysql.servers":                       MySQLServers(),
			"network.virtual_networks":            NetworkVirtualNetworks(),
			"network.security_groups":             NetworkSecurityGroups(),
			"network.public_ip_addresses":         NetworkPublicIPAddresses(),
			"network.watchers":                    NetworkWatchers(),
			"postgresql.servers":                  PostgresqlServers(),
			"resources.groups":                    ResourcesGroups(),
			"security.auto_provisioning_settings": SecurityAutoProvisioningSettings(),
			"security.contacts":                   SecurityContacts(),
			"security.pricings":                   SecurityPricings(),
			"security.settings":                   SecuritySettings(),
			"sql.servers":                         SQLServers(),
			"storage.accounts":                    StorageAccounts(),
			"subscription.subscriptions":          SubscriptionSubscriptions(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
