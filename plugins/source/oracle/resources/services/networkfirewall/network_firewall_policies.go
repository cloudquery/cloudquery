package networkfirewall

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/networkfirewall"
)

func NetworkFirewallPolicies() *schema.Table {
	return &schema.Table{
		Name:      "oracle_networkfirewall_network_firewall_policies",
		Resolver:  fetchNetworkFirewallPolicies,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&networkfirewall.NetworkFirewallPolicySummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
