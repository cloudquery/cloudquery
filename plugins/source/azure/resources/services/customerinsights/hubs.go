package customerinsights

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Hubs() *schema.Table {
	return &schema.Table{
		Name:                 "azure_customerinsights_hubs",
		Resolver:             fetchHubs,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights@v1.0.0#Hub",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_customerinsights_hubs", client.Namespacemicrosoft_customerinsights),
		Transform:            transformers.TransformWithStruct(&armcustomerinsights.Hub{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchHubs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcustomerinsights.NewHubsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
