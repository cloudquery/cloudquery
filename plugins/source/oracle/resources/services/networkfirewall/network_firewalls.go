package networkfirewall

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/networkfirewall"
)

func NetworkFirewalls() *schema.Table {
	return &schema.Table{
		Name:      "oracle_networkfirewall_network_firewalls",
		Resolver:  fetchNetworkFirewalls,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&networkfirewall.NetworkFirewallSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
