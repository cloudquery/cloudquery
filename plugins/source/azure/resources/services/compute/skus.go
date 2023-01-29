package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SKUs() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_skus",
		Resolver:    fetchResourceSKUs,
		Description: "https://learn.microsoft.com/en-us/rest/api/compute/resource-skus/list?tabs=HTTP#resourceskusresult",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_compute_skus", client.Namespacemicrosoft_compute),
		Transform: transformers.TransformWithStruct(&armcompute.ResourceSKU{},
			transformers.WithPrimaryKeys("Family", "Kind", "Name"),
		),
		Columns: schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchResourceSKUs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewResourceSKUsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(&armcompute.ResourceSKUsClientListOptions{IncludeExtendedLocations: to.Ptr("true")})
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
