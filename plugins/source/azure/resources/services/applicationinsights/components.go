package applicationinsights

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Components() *schema.Table {
	return &schema.Table{
		Name:                 "azure_applicationinsights_components",
		Resolver:             fetchComponents,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/application-insights/components/list?tabs=HTTP#applicationinsightscomponent",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_applicationinsights_components", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&armapplicationinsights.Component{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchComponents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armapplicationinsights.NewComponentsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
