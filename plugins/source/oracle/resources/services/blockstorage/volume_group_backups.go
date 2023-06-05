package blockstorage

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func VolumeGroupBackups() *schema.Table {
	return &schema.Table{
		Name:      "oracle_blockstorage_volume_group_backups",
		Resolver:  fetchVolumeGroupBackups,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.VolumeGroupBackup{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
