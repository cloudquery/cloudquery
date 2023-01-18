package search

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:      "azure_search_services",
		Resolver:  fetchServices,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_search_services", client.Namespacemicrosoft_search),
		Transform: transformers.TransformWithStruct(&armsearch.Service{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
