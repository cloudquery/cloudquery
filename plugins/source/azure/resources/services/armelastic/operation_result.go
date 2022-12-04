// Code generated by codegen; DO NOT EDIT.

package armelastic

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func OperationResult() *schema.Table {
	return &schema.Table{
		Name:      "azure_armelastic_operation_result",
		Resolver:  fetchOperationResult,
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
				Name:     "origin",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Origin"),
			},
		},
	}
}

func fetchOperationResult(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armelastic.NewOperationsClient(cl.Creds, cl.Options)
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
