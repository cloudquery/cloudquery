package appservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func webAppAuthSettings() *schema.Table {
	return &schema.Table{
		Name:      "azure_appservice_web_app_auth_settings",
		Resolver:  fetchWebAppAuthSettings,
		Transform: transformers.TransformWithStruct(&armappservice.SiteAuthSettings{}),
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

func fetchWebAppAuthSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
	a, err := svc.GetAuthSettings(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	res <- a
	return nil
}
