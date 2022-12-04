// Code generated by codegen; DO NOT EDIT.

package armstoragepool

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RpOperation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armstoragepool_rp_operation",
		Resolver:  fetchRpOperation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "display",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Display"),
			},
			{
				Name:     "is_data_action",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDataAction"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "action_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActionType"),
			},
			{
				Name:     "origin",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Origin"),
			},
		},
	}
}

func fetchRpOperation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmstoragepoolRpOperation
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
