package search

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "azure_search_services",
		Resolver:    fetchServices,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch@v1.0.0#Service",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_search_services", client.Namespacemicrosoft_search),
		Transform:   transformers.TransformWithStruct(&armsearch.Service{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			client.IDColumn,
		},
	}
}
