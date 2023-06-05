package blockstorage

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func BootVolumeReplicas() *schema.Table {
	return &schema.Table{
		Name:      "oracle_blockstorage_boot_volume_replicas",
		Resolver:  fetchBootVolumeReplicas,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.BootVolumeReplica{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
