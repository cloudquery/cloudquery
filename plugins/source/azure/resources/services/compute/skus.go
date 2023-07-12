package compute

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/mitchellh/hashstructure/v2"
)

func SKUs() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_skus",
		Resolver:             fetchResourceSKUs,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/resource-skus/list?tabs=HTTP#resourceskusresult",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_skus", client.Namespacemicrosoft_compute),
		Transform: transformers.TransformWithStruct(&armcompute.ResourceSKU{},
			transformers.WithPrimaryKeys("Name"),
		),
		Columns: schema.ColumnList{
			client.SubscriptionIDPK,
			schema.Column{
				Name:       "_sku_hash",
				Type:       arrow.BinaryTypes.String,
				Resolver:   calcEntryHash,
				PrimaryKey: true,
			},
		},
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

func calcEntryHash(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	hash, err := hashstructure.Hash(resource.Item, hashstructure.FormatV2, nil)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, fmt.Sprint(hash))
}
