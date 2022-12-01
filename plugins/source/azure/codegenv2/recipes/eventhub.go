package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

func EventHubResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "namespaces",
			Struct: &armeventhub.EHNamespace{},
			ResponseStruct: &armeventhub.NamespacesClientListResponse{},
			Client: &armeventhub.NamespacesClient{},
			ListFunc: (&armeventhub.NamespacesClient{}).NewListPager,
			NewFunc: armeventhub.NewNamespacesClient,
			OutputField: "Value",
		},
		{
			SubService: "network_rule_sets",
			Struct: &armeventhub.NetworkRuleSet{},
			ResponseStruct: &armeventhub.NetworkRuleSetListResult{},
			Client: &armeventhub.NamespacesClient{},
			ListFunc: (&armeventhub.NamespacesClient{}).ListNetworkRuleSet,
			NewFunc: armeventhub.NewNamespacesClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "eventhub/armeventhub"
		r.Service = "armeventhub"
		r.Template = "list"
	}

	return resources
}