package virtualnetwork

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func ByoipRanges() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_byoip_ranges",
		Resolver:  fetchByoipRanges,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.ByoipRangeSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
