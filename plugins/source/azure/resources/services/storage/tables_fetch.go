package storage

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armstorage.NewTableClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

func getTable(ctx context.Context, meta schema.ClientMeta, res *schema.Resource) error {
	cl := meta.(*client.Client)
	svc, err := armstorage.NewTableClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	item := res.Parent.Item.(*armstorage.Account)
	group, err := client.ParseResourceGroup(*item.ID)
	if err != nil {
		return err
	}

	t := res.Item.(*armstorage.Table)

	v, err := svc.Get(ctx, group, *item.Name, *t.Name, nil)
	if err != nil {
		return err
	}
	res.Item = v.Table
	return nil
}
