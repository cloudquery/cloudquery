package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DscpConfiguration() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_dscp_configuration",
		Resolver:             fetchDscpConfiguration,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/dscp-configuration/list?tabs=HTTP#dscpconfiguration",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_dscp_configuration", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.DscpConfiguration{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDscpConfiguration(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewDscpConfigurationClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAllPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
