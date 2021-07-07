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
			"compute.disks":                       ComputeDisks(),
			"keyvault.vaults":                     KeyVaultVaults(),
			"mysql.servers":                       MySQLServers(),
			"network.virtual_networks":            NetworkVirtualNetworks(),
			"postgresql.servers":                  PostgresqlServers(),
			"resources.groups":                    ResourcesGroups(),
			"security.auto_provisioning_settings": SecurityAutoProvisioningSettings(),
			"security.contacts":                   SecurityContacts(),
			"security.pricings":                   SecurityPricings(),
			"security.settings":                   SecuritySettings(),
			"sql.servers":                         SQLServers(),
			"storage.accounts":                    StorageAccounts(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
