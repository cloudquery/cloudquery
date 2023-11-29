package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Providers() *schema.Table {
	return &schema.Table{
		Name:                 "azure_resources_providers",
		Resolver:             fetchProviders,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://docs.microsoft.com/en-us/rest/api/resources/providers/list",
		Multiplex:            client.SubscriptionMultiplex,
		Transform:            transformers.TransformWithStruct(&armresources.Provider{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armresources.NewProvidersClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(&armresources.ProvidersClientListOptions{
		Expand: to.Ptr("resourceTypes/aliases"),
	})
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
