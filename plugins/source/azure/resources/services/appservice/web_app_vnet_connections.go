package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func webAppVnetConnections() *schema.Table {
	return &schema.Table{
		Name:                 "azure_appservice_web_app_vnet_connections",
		Resolver:             fetchWebAppVnetConnections,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list-vnet-connections#vnetinforesource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_web_app_vnet_connections", client.Namespacemicrosoft_web),
		Transform:            transformers.TransformWithStruct(&armappservice.VnetInfoResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchWebAppVnetConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
	resp, err := svc.ListVnetConnections(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.VnetInfoResourceArray
	return nil
}
