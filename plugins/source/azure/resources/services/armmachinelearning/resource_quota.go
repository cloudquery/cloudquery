// Code generated by codegen; DO NOT EDIT.

package armmachinelearning

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceQuota() *schema.Table {
	return &schema.Table{
		Name:      "azure_armmachinelearning_resource_quota",
		Resolver:  fetchResourceQuota,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "aml_workspace_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AmlWorkspaceLocation"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Limit"),
			},
			{
				Name:     "name",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Unit"),
			},
		},
	}
}

func fetchResourceQuota(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmmachinelearningResourceQuota
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
