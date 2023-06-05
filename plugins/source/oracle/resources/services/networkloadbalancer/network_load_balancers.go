package networkloadbalancer

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func NetworkLoadBalancers() *schema.Table {
	return &schema.Table{
		Name:      "oracle_networkloadbalancer_network_load_balancers",
		Resolver:  fetchNetworkLoadBalancers,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&networkloadbalancer.NetworkLoadBalancerSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
