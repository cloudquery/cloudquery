package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DeletedWebApps() *schema.Table {
	return &schema.Table{
		Name:                 "azure_appservice_deleted_web_apps",
		Resolver:             fetchDeletedWebApps,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/appservice/deleted-web-apps/list#deletedsite",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_deleted_web_apps", client.Namespacemicrosoft_web),
		Transform:            transformers.TransformWithStruct(&armappservice.DeletedSite{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDeletedWebApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappservice.NewDeletedWebAppsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
