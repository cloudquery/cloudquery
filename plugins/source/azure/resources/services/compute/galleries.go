package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Galleries() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_galleries",
		Resolver:             fetchGalleries,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/galleries/list?tabs=HTTP#gallery",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_galleries", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.Gallery{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations:            []*schema.Table{galleryImages()},
	}
}

func fetchGalleries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewGalleriesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
