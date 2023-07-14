package appconfiguration

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ConfigurationStores() *schema.Table {
	return &schema.Table{
		Name:                 "azure_appconfiguration_configuration_stores",
		Resolver:             fetchConfigurationStores,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/appconfiguration/stable/configuration-stores/list?tabs=HTTP#configurationstore",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_appconfiguration_configuration_stores", client.Namespacemicrosoft_appconfiguration),
		Transform:            transformers.TransformWithStruct(&armappconfiguration.ConfigurationStore{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchConfigurationStores(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappconfiguration.NewConfigurationStoresClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
