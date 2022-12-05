// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"

func Armdnsresolver() []Table {
	tables := []Table{
		{
      Name: "dns_forwarding_ruleset",
      Struct: &armdnsresolver.DNSForwardingRuleset{},
      ResponseStruct: &armdnsresolver.DNSForwardingRulesetsClientListResponse{},
      Client: &armdnsresolver.DNSForwardingRulesetsClient{},
      ListFunc: (&armdnsresolver.DNSForwardingRulesetsClient{}).NewListPager,
			NewFunc: armdnsresolver.NewDNSForwardingRulesetsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsForwardingRulesets",
		},
		{
      Name: "dns_resolver",
      Struct: &armdnsresolver.DNSResolver{},
      ResponseStruct: &armdnsresolver.DNSResolversClientListResponse{},
      Client: &armdnsresolver.DNSResolversClient{},
      ListFunc: (&armdnsresolver.DNSResolversClient{}).NewListPager,
			NewFunc: armdnsresolver.NewDNSResolversClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsResolvers",
		},
		{
      Name: "forwarding_rule",
      Struct: &armdnsresolver.ForwardingRule{},
      ResponseStruct: &armdnsresolver.ForwardingRulesClientListResponse{},
      Client: &armdnsresolver.ForwardingRulesClient{},
      ListFunc: (&armdnsresolver.ForwardingRulesClient{}).NewListPager,
			NewFunc: armdnsresolver.NewForwardingRulesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules",
		},
		{
      Name: "inbound_endpoint",
      Struct: &armdnsresolver.InboundEndpoint{},
      ResponseStruct: &armdnsresolver.InboundEndpointsClientListResponse{},
      Client: &armdnsresolver.InboundEndpointsClient{},
      ListFunc: (&armdnsresolver.InboundEndpointsClient{}).NewListPager,
			NewFunc: armdnsresolver.NewInboundEndpointsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/inboundEndpoints",
		},
		{
      Name: "outbound_endpoint",
      Struct: &armdnsresolver.OutboundEndpoint{},
      ResponseStruct: &armdnsresolver.OutboundEndpointsClientListResponse{},
      Client: &armdnsresolver.OutboundEndpointsClient{},
      ListFunc: (&armdnsresolver.OutboundEndpointsClient{}).NewListPager,
			NewFunc: armdnsresolver.NewOutboundEndpointsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints",
		},
		{
      Name: "virtual_network_link",
      Struct: &armdnsresolver.VirtualNetworkLink{},
      ResponseStruct: &armdnsresolver.VirtualNetworkLinksClientListResponse{},
      Client: &armdnsresolver.VirtualNetworkLinksClient{},
      ListFunc: (&armdnsresolver.VirtualNetworkLinksClient{}).NewListPager,
			NewFunc: armdnsresolver.NewVirtualNetworkLinksClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks",
		},
	}

	for i := range tables {
		tables[i].Service = "armdnsresolver"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdnsresolver()...)
}