package account

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AccountLocations() *schema.Table {
	return &schema.Table{
		Name:         "azure_account_locations",
		Description:  "Azure location information",
		Resolver:     fetchAccountLocations,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "id",
				Description: "The fully qualified ID of the location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The location name",
				Type:        schema.TypeString,
			},
			{
				Name:        "display_name",
				Description: "The display name of the location",
				Type:        schema.TypeString,
			},
			{
				Name:        "latitude",
				Description: "The latitude of the location",
				Type:        schema.TypeString,
			},
			{
				Name:        "longitude",
				Description: "The longitude of the location",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccountLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	response, err := svc.Subscriptions.ListLocations(ctx, svc.SubscriptionID)
	if err != nil {
		return err
	}
	res <- *response.Value
	return nil
}
