package virtualnetwork

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func FastConnectProviderServices() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_fast_connect_provider_services",
		Resolver:  fetchFastConnectProviderServices,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.FastConnectProviderService{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
