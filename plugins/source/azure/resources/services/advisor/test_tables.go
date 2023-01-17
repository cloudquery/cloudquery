package advisor

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TestTables() *schema.Table {
	return &schema.Table{
		Name:      "azure_test_tables",
		Resolver:  fetchSuppressions,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_advisor_suppressions", client.Namespacemicrosoft_advisor),
		Transform: transformers.TransformWithStruct(&armadvisor.SuppressionContract{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
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
