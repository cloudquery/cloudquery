package loadbalancer

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/loadbalancer"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:      "oracle_loadbalancer_load_balancers",
		Resolver:  fetchLoadBalancers,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&loadbalancer.LoadBalancer{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
