package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DiskAccesses() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_disk_accesses",
		Resolver:             fetchDiskAccesses,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/disk-accesses/list?tabs=HTTP#diskaccess",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_disk_accesses", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.DiskAccess{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDiskAccesses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewDiskAccessesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
