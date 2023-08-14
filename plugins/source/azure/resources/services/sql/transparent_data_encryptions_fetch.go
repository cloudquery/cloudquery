package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchTransparentDataEncryptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pp := parent.Parent.Item.(*armsql.Server)
	p := parent.Item.(*armsql.Database)
	cl := meta.(*client.Client)
	svc, err := armsql.NewTransparentDataEncryptionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByDatabasePager(group, *pp.Name, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
