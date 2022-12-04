// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"

func Armdnsresolver() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdnsresolver.NewDNSForwardingRulesetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
		},
		{
			NewFunc: armdnsresolver.NewInboundEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
		},
		{
			NewFunc: armdnsresolver.NewOutboundEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
		},
		{
			NewFunc: armdnsresolver.NewVirtualNetworkLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
		},
		{
			NewFunc: armdnsresolver.NewDNSResolversClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
		},
		{
			NewFunc: armdnsresolver.NewForwardingRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdnsresolver())
}