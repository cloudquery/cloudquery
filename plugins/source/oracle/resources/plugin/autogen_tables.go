package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/loadbalancer"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/networkfirewall"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/networkloadbalancer"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/virtualnetwork"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		compute.CapacityReservations(),
		compute.ConsoleHistories(),
		compute.DedicatedVMHosts(),
		compute.Images(),
		compute.InstanceConsoleConnections(),
		compute.Instances(),
		compute.VNICAttachments(),
		compute.VolumeAttachments(),
		loadbalancer.LoadBalancers(),
		networkfirewall.FirewallPolicies(),
		networkfirewall.NetworkFirewalls(),
		networkfirewall.WorkRequests(),
		networkloadbalancer.NetworkLoadBalancers(),
		networkloadbalancer.WorkRequests(),
		virtualnetwork.BYOIPRanges(),
		virtualnetwork.CaptureFilters(),
		virtualnetwork.CPEs(),
		virtualnetwork.CrossConnectGroups(),
		virtualnetwork.CrossConnects(),
		virtualnetwork.DHCPOptions(),
		virtualnetwork.DRGAttachments(),
		virtualnetwork.DRGs(),
		virtualnetwork.FastConnectProviderServices(),
		virtualnetwork.InternetGateways(),
		virtualnetwork.IPSecConnections(),
		virtualnetwork.LocalPeeringGateways(),
		virtualnetwork.NATGateways(),
		virtualnetwork.PrivateIPs(),
		virtualnetwork.PublicIpPools(),
		virtualnetwork.AssignedPublicIPs(),
		virtualnetwork.EphemeralPublicIPs(),
		virtualnetwork.ReservedPublicIPs(),
		virtualnetwork.RemotePeeringConnections(),
		virtualnetwork.RouteTables(),
		virtualnetwork.SecurityLists(),
		virtualnetwork.ServiceGateways(),
		virtualnetwork.Subnets(),
		virtualnetwork.VCNs(),
		virtualnetwork.VirtualCircuits(),
		virtualnetwork.VLANs(),
		virtualnetwork.VTAPs(),
	}
}
