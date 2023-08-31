package saas

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:                 "azure_saas_resources",
		Resolver:             fetchResources,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas@v0.5.0#Resource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_saas_resources", client.Namespacemicrosoft_saas),
		Transform:            transformers.TransformWithStruct(&armsaas.Resource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsaas.NewResourcesClient(cl.Creds, cl.Options)
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
