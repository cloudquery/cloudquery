package compute

import (
	"context"
	"crypto/sha256"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"golang.org/x/exp/slices"
)

func SKUs() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_skus",
		Resolver:    fetchResourceSKUs,
		Description: "https://learn.microsoft.com/en-us/rest/api/compute/resource-skus/list?tabs=HTTP#resourceskusresult",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_compute_skus", client.Namespacemicrosoft_compute),
		Transform: transformers.TransformWithStruct(&armcompute.ResourceSKU{},
			transformers.WithPrimaryKeys("Name", "ResourceType"),
		),
		Columns: schema.ColumnList{
			client.SubscriptionIDPK,
			schema.Column{
				Name:            "locations_hash",
				Type:            schema.TypeString,
				Resolver:        calcLocationsHash,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			schema.Column{
				Name:            "capabilities_hash",
				Type:            schema.TypeString,
				Resolver:        calcCapabilitiesHash,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

func fetchResourceSKUs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
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

func calcLocationsHash(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sku := resource.Item.(*armcompute.ResourceSKU)
	slices.SortFunc(sku.Locations, func(a, b *string) bool {
		return *a < *b
	})

	h := sha256.New()
	for _, loc := range sku.Locations {
		h.Write([]byte(*loc))
	}

	return resource.Set(c.Name, string(h.Sum(nil)))
}

func calcCapabilitiesHash(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sku := resource.Item.(*armcompute.ResourceSKU)
	slices.SortFunc(sku.Capabilities, func(a, b *armcompute.ResourceSKUCapabilities) bool {
		return *a.Name < *b.Name
	})

	h := sha256.New()
	for _, capability := range sku.Capabilities {
		h.Write([]byte(*capability.Name + ":" + *capability.Value))
	}

	return resource.Set(c.Name, string(h.Sum(nil)))
}
