package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func webAppConfigurations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_appservice_web_app_configurations",
		Resolver:             fetchWebAppConfigurations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list-configurations#siteconfigresource",
		Transform:            transformers.TransformWithStruct(&armappservice.SiteConfigResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchWebAppConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armappservice.Site)
	cl := meta.(*client.Client)
	svc, err := armappservice.NewWebAppsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListConfigurationsPager(group, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
