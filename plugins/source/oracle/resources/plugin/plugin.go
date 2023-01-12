package plugin

import (
	"sort"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/blockstorage"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/database"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/filestorage"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/identity"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/objectstorage"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

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

func Plugin() *source.Plugin {
	allTables := append(AutogenTables(), customTables...)

	sort.Slice(allTables, func(i, j int) bool {
		return allTables[i].Name < allTables[j].Name
	})

	// here you can append custom non-generated tables
	return source.NewPlugin(
		"oracle",
		Version,
		allTables,
		client.Configure,
	)
}
