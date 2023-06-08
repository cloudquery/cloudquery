package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:                 "azure_monitor_resources",
		Resolver:             fetchResources,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/resources/resources/list#genericresourceexpanded",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_diagnostic_settings", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&armresources.GenericResourceExpanded{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations:            []*schema.Table{diagnosticSettings()},
	}
}

func fetchResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armresources.NewClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
