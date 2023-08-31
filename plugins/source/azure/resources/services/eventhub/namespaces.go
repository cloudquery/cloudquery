package eventhub

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:                 "azure_eventhub_namespaces",
		Resolver:             fetchNamespaces,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/eventhub/stable/namespaces/list?tabs=HTTP#ehnamespace",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_eventhub_namespaces", client.Namespacemicrosoft_eventhub),
		Transform:            transformers.TransformWithStruct(&armeventhub.EHNamespace{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			namespaceNetworkRuleSets(),
		},
	}
}

func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armeventhub.NewNamespacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
