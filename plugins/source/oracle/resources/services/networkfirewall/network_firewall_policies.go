package networkfirewall

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/oracle/oci-go-sdk/v65/networkfirewall"
)

func NetworkFirewallPolicies() *schema.Table {
	return &schema.Table{
		Name:      "oracle_networkfirewall_network_firewall_policies",
		Resolver:  fetchNetworkFirewallPolicies,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: transformers.TransformWithStruct(&networkfirewall.NetworkFirewallPolicySummary{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: []schema.Column{
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveOracleRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "compartment_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveCompartmentId,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
