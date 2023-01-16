package eventhub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func networkRuleSets() *schema.Table {
	return &schema.Table{
		Name:      "azure_eventhub_network_rule_sets",
		Resolver:  fetchNetworkRuleSets,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_eventhub_network_rule_sets", client.Namespacemicrosoft_eventhub),
		Transform: transformers.TransformWithStruct(&armeventhub.NetworkRuleSet{}),
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

func fetchNetworkRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armeventhub.EHNamespace)
	cl := meta.(*client.Client)
	svc, err := armeventhub.NewNamespacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}

	resp, err := svc.ListNetworkRuleSet(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.Value

	return nil
}
