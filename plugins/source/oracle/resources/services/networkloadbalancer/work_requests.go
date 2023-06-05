package networkloadbalancer

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func WorkRequests() *schema.Table {
	return &schema.Table{
		Name:      "oracle_networkloadbalancer_work_requests",
		Resolver:  fetchWorkRequests,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&networkloadbalancer.WorkRequestSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
