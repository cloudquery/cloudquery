package subscription

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func locations() *schema.Table {
	return &schema.Table{
		Name:      "azure_subscription_locations",
		Resolver:  fetchLocations,
		Transform: transformers.TransformWithStruct(&armsubscription.Location{}),
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

func fetchLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armsubscription.Subscription)
	cl := meta.(*client.Client)

	svc, err := armsubscription.NewSubscriptionsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListLocationsPager(*p.ID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
