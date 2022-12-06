// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"

func Armdnsresolver() []Table {
	tables := []Table{
		{
			Name:           "dns_forwarding_ruleset",
			Struct:         &armdnsresolver.DNSForwardingRuleset{},
			ResponseStruct: &armdnsresolver.DNSForwardingRulesetsClientListResponse{},
			Client:         &armdnsresolver.DNSForwardingRulesetsClient{},
			ListFunc:       (&armdnsresolver.DNSForwardingRulesetsClient{}).NewListPager,
			NewFunc:        armdnsresolver.NewDNSForwardingRulesetsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsForwardingRulesets",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Network")`,
		},
		{
			Name:           "dns_resolver",
			Struct:         &armdnsresolver.DNSResolver{},
			ResponseStruct: &armdnsresolver.DNSResolversClientListResponse{},
			Client:         &armdnsresolver.DNSResolversClient{},
			ListFunc:       (&armdnsresolver.DNSResolversClient{}).NewListPager,
			NewFunc:        armdnsresolver.NewDNSResolversClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsResolvers",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Network")`,
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
