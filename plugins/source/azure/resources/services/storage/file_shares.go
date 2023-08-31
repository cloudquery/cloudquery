package storage

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func file_shares() *schema.Table {
	return &schema.Table{
		Name:                 "azure_storage_file_shares",
		Resolver:             fetchFileShares,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/storagerp/file-shares/list?tabs=HTTP#fileshareitem",
		Transform:            transformers.TransformWithStruct(&armstorage.FileShareItem{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchFileShares(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armstorage.NewFileSharesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	item := parent.Item.(*armstorage.Account)
	group, err := client.ParseResourceGroup(*item.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(group, *item.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
