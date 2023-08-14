package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PrivateLinkScopes() *schema.Table {
	return &schema.Table{
		Name:                 "azure_monitor_private_link_scopes",
		Resolver:             fetchPrivateLinkScopes,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/monitor/privatelinkscopes(preview)/private%20link%20scopes%20(preview)/list?tabs=HTTP#azuremonitorprivatelinkscope",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_private_link_scopes", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&armmonitor.AzureMonitorPrivateLinkScope{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPrivateLinkScopes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmonitor.NewPrivateLinkScopesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
