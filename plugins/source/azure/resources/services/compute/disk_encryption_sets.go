package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DiskEncryptionSets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_disk_encryption_sets",
		Resolver:             fetchDiskEncryptionSets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/disk-encryption-sets/list?tabs=HTTP#diskencryptionset",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_disk_encryption_sets", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.DiskEncryptionSet{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDiskEncryptionSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewDiskEncryptionSetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
