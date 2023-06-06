package virtualnetwork

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func LocalPeeringGateways() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_local_peering_gateways",
		Resolver:  fetchLocalPeeringGateways,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.LocalPeeringGateway{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
