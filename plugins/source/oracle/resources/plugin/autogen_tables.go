package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/loadbalancer"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/networkfirewall"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/networkloadbalancer"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/virtualnetwork"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AutogenTables() []*schema.Table {
	return []*schema.Table{
		compute.ComputeCapacityReservations(),
		compute.ConsoleHistories(),
		compute.DedicatedVmHosts(),
		compute.Images(),
		compute.InstanceConsoleConnections(),
		compute.Instances(),
		compute.VnicAttachments(),
		loadbalancer.LoadBalancers(),
		networkfirewall.NetworkFirewallPolicies(),
		networkfirewall.NetworkFirewalls(),
		networkfirewall.WorkRequests(),
		networkloadbalancer.NetworkLoadBalancers(),
		networkloadbalancer.WorkRequests(),
		virtualnetwork.ByoipRanges(),
		virtualnetwork.CaptureFilters(),
		virtualnetwork.Cpes(),
		virtualnetwork.CrossConnectGroups(),
		virtualnetwork.CrossConnects(),
		virtualnetwork.DhcpOptions(),
		virtualnetwork.DrgAttachments(),
		virtualnetwork.Drgs(),
		virtualnetwork.FastConnectProviderServices(),
		virtualnetwork.InternetGateways(),
		virtualnetwork.IpSecConnections(),
		virtualnetwork.LocalPeeringGateways(),
		virtualnetwork.NatGateways(),
		virtualnetwork.PublicIpPools(),
		virtualnetwork.RemotePeeringConnections(),
		virtualnetwork.RouteTables(),
		virtualnetwork.SecurityLists(),
		virtualnetwork.ServiceGateways(),
		virtualnetwork.Subnets(),
		virtualnetwork.Vcns(),
		virtualnetwork.VirtualCircuits(),
		virtualnetwork.Vlans(),
		virtualnetwork.Vtaps(),
	}
}
