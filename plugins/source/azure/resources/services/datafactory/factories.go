package datafactory

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Factories() *schema.Table {
	return &schema.Table{
		Name:                 "azure_datafactory_factories",
		Resolver:             fetchFactories,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/datafactory/factories/list?tabs=HTTP#factory",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_datafactory_factories", client.Namespacemicrosoft_datafactory),
		Transform: transformers.TransformWithStruct(&armdatafactory.Factory{},
			transformers.WithNameTransformer(client.ETagNameTransformer),
			transformers.WithPrimaryKeys("ID"),
		),
		Columns: schema.ColumnList{client.SubscriptionID},
	}
}

func fetchFactories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatafactory.NewFactoriesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
