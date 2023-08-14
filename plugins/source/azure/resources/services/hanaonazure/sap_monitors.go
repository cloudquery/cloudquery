package hanaonazure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SapMonitors() *schema.Table {
	return &schema.Table{
		Name:                 "azure_hanaonazure_sap_monitors",
		Resolver:             fetchSapMonitors,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure@v0.5.0#SapMonitor",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_hanaonazure_sap_monitors", client.Namespacemicrosoft_hanaonazure),
		Transform:            transformers.TransformWithStruct(&armhanaonazure.SapMonitor{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSapMonitors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armhanaonazure.NewSapMonitorsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
