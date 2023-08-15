package costmanagement

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Views() *schema.Table {
	return &schema.Table{
		Name:                 "azure_costmanagement_views",
		Resolver:             fetchViews,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cost-management/views/list?tabs=HTTP#view",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_costmanagement_views", client.Namespacemicrosoft_costmanagement),
		Transform: transformers.TransformWithStruct(&armcostmanagement.View{},
			transformers.WithNameTransformer(client.ETagNameTransformer),
			transformers.WithPrimaryKeys("ID"),
		),
		Columns: schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			view_queries(),
		},
	}
}

func fetchViews(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcostmanagement.NewViewsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListByScopePager("subscriptions/"+cl.SubscriptionId, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
