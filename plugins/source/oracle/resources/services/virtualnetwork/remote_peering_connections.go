package virtualnetwork

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func RemotePeeringConnections() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_remote_peering_connections",
		Resolver:  fetchRemotePeeringConnections,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.RemotePeeringConnection{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
