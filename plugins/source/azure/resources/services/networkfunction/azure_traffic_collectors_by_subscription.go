package networkfunction

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AzureTrafficCollectorsBySubscription() *schema.Table {
	return &schema.Table{
		Name:                 "azure_networkfunction_azure_traffic_collectors_by_subscription",
		Resolver:             fetchAzureTrafficCollectorsBySubscription,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction@v1.0.0#AzureTrafficCollector",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_networkfunction_azure_traffic_collectors_by_subscription", client.Namespacemicrosoft_networkfunction),
		Transform:            transformers.TransformWithStruct(&armnetworkfunction.AzureTrafficCollector{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAzureTrafficCollectorsBySubscription(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetworkfunction.NewAzureTrafficCollectorsBySubscriptionClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
