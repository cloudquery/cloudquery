package eventhub

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func namespaceNetworkRuleSets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_eventhub_namespace_network_rule_sets",
		Resolver:             fetchNetworkRuleSets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/eventhub/stable/network-rule-sets/list-network-rule-set?tabs=HTTP#networkruleset",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_eventhub_network_rule_sets", client.Namespacemicrosoft_eventhub),
		Transform:            transformers.TransformWithStruct(&armeventhub.NetworkRuleSet{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
