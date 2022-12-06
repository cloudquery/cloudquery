// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"

func Armdnsresolver() []*Table {
	tables := []*Table{
		{
			NewFunc: armdnsresolver.NewDNSForwardingRulesetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsForwardingRulesets",
		},
		{
			NewFunc: armdnsresolver.NewDNSResolversClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsResolvers",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdnsresolver())
}