package compute

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func InstanceConsoleConnections() *schema.Table {
	return &schema.Table{
		Name:      "oracle_compute_instance_console_connections",
		Resolver:  fetchInstanceConsoleConnections,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.InstanceConsoleConnection{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
