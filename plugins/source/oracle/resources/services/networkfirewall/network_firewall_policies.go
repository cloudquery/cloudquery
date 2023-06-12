package networkfirewall

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/networkfirewall"
)

func FirewallPolicies() *schema.Table {
	return &schema.Table{
		Name:      "oracle_networkfirewall_network_firewall_policies",
		Resolver:  fetchFirewallPolicies,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&networkfirewall.NetworkFirewallPolicySummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchFirewallPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := networkfirewall.ListNetworkFirewallPoliciesRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].NetworkfirewallNetworkfirewallClient.ListNetworkFirewallPolicies(ctx, request)

		if err != nil {
			return err
		}

		res <- response.Items

		if response.OpcNextPage == nil {
			break
		}

		page = response.OpcNextPage
	}

	return nil
}
