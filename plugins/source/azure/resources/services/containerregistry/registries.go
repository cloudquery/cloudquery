package containerregistry

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:                 "azure_containerregistry_registries",
		Resolver:             fetchRegistries,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/containerregistry/registries/list?tabs=HTTP#registry",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_containerregistry_registries", client.Namespacemicrosoft_containerregistry),
		Transform:            transformers.TransformWithStruct(&armcontainerregistry.Registry{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcontainerregistry.NewRegistriesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
