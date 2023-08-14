package applicationinsights

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func WebTests() *schema.Table {
	return &schema.Table{
		Name:                 "azure_applicationinsights_web_tests",
		Resolver:             fetchWebTests,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/application-insights/web-tests/list?tabs=HTTP#webtest",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_applicationinsights_web_tests", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&armapplicationinsights.WebTest{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchWebTests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armapplicationinsights.NewWebTestsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
