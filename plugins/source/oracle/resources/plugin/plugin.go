package plugin

import (
	"sort"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client/spec"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/blockstorage"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/database"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/filestorage"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/identity"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/objectstorage"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"golang.org/x/exp/maps"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"blockstorage":    "Block Storage",
	"byoip":           "Bring Your Own IP (BYOIP)",
	"cpes":            "Customer Premises Equipment (CPEs)",
	"dhcp":            "Dynamic Host Configuration Protocol (DHCP)",
	"drg":             "Dynamic Routing Gateway (DRG)",
	"drgs":            "Dynamic Routing Gateways (DRGs)",
	"filestorage":     "File Storage",
	"loadbalancer":    "Load Balancer",
	"networkfirewall": "Network Firewall",
	"objectstorage":   "Object Storage",
	"virtualnetwork":  "Virtual Network",
	"vlans":           "Virtual LANs (VLANs)",
	"vm":              "Virtual Machine (VM)",
	"vnic":            "Virtual Network Interface Card (VNIC)",
	"vtaps":           "Virtual Tunnel Access Points (VTAPs)",
}

func titleTransformer(table *schema.Table) error {
	if table.Title != "" {
		return nil
	}

	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	table.Title = csr.ToTitle(table.Name)
	return nil
}

var customTables = []*schema.Table{
	blockstorage.BootVolumeBackups(),
	blockstorage.BootVolumeReplicas(),
	blockstorage.BootVolumes(),
	blockstorage.VolumeBackupPolicies(),
	blockstorage.VolumeBackups(),
	blockstorage.VolumeGroupBackups(),
	blockstorage.VolumeGroups(),
	blockstorage.Volumes(),
	database.AutonomousContainerDatabases(),
	database.AutonomousDatabases(),
	database.AutonomousExadataInfrastructures(),
	database.AutonomousVmClusters(),
	database.BackupDestination(),
	database.CloudAutonomousVmClusters(),
	database.CloudExadataInfrastructures(),
	database.CloudVmClusters(),
	database.ExadataInfrastructures(),
	database.ExternalContainerDatabases(),
	database.ExternalNonContainerDatabases(),
	database.ExternalPluggableDatabases(),
	database.KeyStores(),
	database.VmClusters(),
	filestorage.ExportSets(),
	filestorage.Exports(),
	filestorage.FileSystems(),
	filestorage.MountTargets(),
	filestorage.ReplicationTargets(),
	filestorage.Replications(),
	identity.Compartments(),
	identity.CostTrackingTags(),
	identity.Domains(),
	identity.DynamicGroups(),
	identity.Groups(),
	identity.IamWorkRequests(),
	identity.NetworkSources(),
	identity.Policies(),
	identity.TagNamespaces(),
	identity.TaggingWorkRequests(),
	identity.Users(),
	identity.WorkRequests(),
	objectstorage.Buckets(),
	objectstorage.WorkRequests(),
}

func getTables() schema.Tables {
	tables := append(Tables(), customTables...)

	sort.Slice(tables, func(i, j int) bool {
		return tables[i].Name < tables[j].Name
	})

	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}

	return tables
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"oracle",
		Version,
		Configure,
		plugin.WithJSONSchema(spec.JSONSchema),
	)
}
