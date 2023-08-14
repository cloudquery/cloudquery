package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func WebApps() *schema.Table {
	return &schema.Table{
		Name:                 "azure_appservice_web_apps",
		Resolver:             fetchWebApps,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list#site",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_web_apps", client.Namespacemicrosoft_web),
		Transform:            transformers.TransformWithStruct(&armappservice.Site{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			webAppAuthSettings(),
			webAppConfigurations(),
			webAppVnetConnections(),
		},
	}
}

func fetchWebApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappservice.NewWebAppsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
