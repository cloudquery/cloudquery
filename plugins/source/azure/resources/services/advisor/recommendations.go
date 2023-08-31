package advisor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Recommendations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_advisor_recommendations",
		Description:          "https://learn.microsoft.com/en-us/rest/api/advisor/recommendations/list?tabs=HTTP#resourcerecommendationbase",
		Resolver:             fetchRecommendations,
		PostResourceResolver: client.LowercaseIDResolver,
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_advisor_recommendations", client.Namespacemicrosoft_advisor),
		Transform:            transformers.TransformWithStruct(&armadvisor.ResourceRecommendationBase{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armadvisor.NewRecommendationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
