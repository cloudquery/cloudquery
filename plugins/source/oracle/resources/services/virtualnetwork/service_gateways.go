package virtualnetwork

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func ServiceGateways() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_service_gateways",
		Resolver:  fetchServiceGateways,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.ServiceGateway{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
