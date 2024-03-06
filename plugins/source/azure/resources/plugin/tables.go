package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/appservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/authorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cosmos"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mysqlflexibleservers"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/peering"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/policy"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresqlflexibleservers"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresqlhsc"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sqlvirtualmachine"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storagecache"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func getTables() schema.Tables {
	list := []*schema.Table{
		appservice.CertificateOrders(),
		appservice.Certificates(),
		appservice.DeletedWebApps(),
		appservice.Domains(),
		appservice.Environments(),
		appservice.Plans(),
		appservice.Recommendations(),
		appservice.ResourceHealthMetadata(),
		appservice.StaticSites(),
		appservice.TopLevelDomains(),
		appservice.WebApps(),
		authorization.ProviderOperationsMetadata(),
		authorization.RoleDefinitions(),
		compute.AvailabilitySets(),
		compute.CapacityReservationGroups(),
		compute.CloudServices(),
		compute.DiskAccesses(),
		compute.DiskEncryptionSets(),
		compute.Disks(),
		compute.Galleries(),
		compute.Images(),
		compute.RestorePointCollections(),
		compute.SSHPublicKeys(),
		compute.SKUs(),
		compute.Snapshots(),
		compute.VirtualMachineScaleSets(),
		compute.VirtualMachines(),
		cosmos.DatabaseAccounts(),
		cosmos.Locations(),
		cosmos.RestorableDatabaseAccounts(),
		mariadb.Servers(),
		mysql.Servers(),
		mysqlflexibleservers.Servers(),
		network.BgpServiceCommunities(),
		peering.ServiceLocations(),
		postgresql.Servers(),
		postgresqlflexibleservers.Servers(),
		postgresqlhsc.ServerGroups(),
		policy.Definitions(),
		security.AssessmentsMetadata(),
		sql.Servers(),
		sql.InstancePools(),
		sql.ManagedInstances(),
		sql.VirtualClusters(),
		sqlvirtualmachine.Groups(),
		sqlvirtualmachine.SqlVirtualMachines(),
		storage.Accounts(),
		storagecache.Caches(),
	}
	for i := range list {
		if list[i].PostResourceResolver == nil {
			panic("no PostResourceResolver in " + list[i].Name)
		}
	}
	if err := transformers.TransformTables(list); err != nil {
		panic(err)
	}
	if err := transformers.Apply(list, titleTransformer); err != nil {
		panic(err)
	}
	for _, table := range list {
		schema.AddCqIDs(table)
	}
	return list
}
