package hybriddatamanager

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DataManagers() *schema.Table {
	return &schema.Table{
		Name:                 "azure_hybriddatamanager_data_managers",
		Resolver:             fetchDataManagers,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager@v1.0.0#DataManager",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_hybriddatamanager_data_managers", client.Namespacemicrosoft_hybriddata),
		Transform:            transformers.TransformWithStruct(&armhybriddatamanager.DataManager{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDataManagers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armhybriddatamanager.NewDataManagersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
