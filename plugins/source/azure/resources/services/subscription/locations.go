package subscription

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func locations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_subscription_subscription_locations",
		Resolver:             fetchLocations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/resources/subscriptions/list-locations?tabs=HTTP#location",
		Transform:            transformers.TransformWithStruct(&armsubscriptions.Location{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armsubscriptions.Subscription)
	cl := meta.(*client.Client)

	svc, err := armsubscriptions.NewClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListLocationsPager(*p.SubscriptionID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
