package compute

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func ComputeCapacityReservations() *schema.Table {
	return &schema.Table{
		Name:      "oracle_compute_compute_capacity_reservations",
		Resolver:  fetchComputeCapacityReservations,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.ComputeCapacityReservationSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
