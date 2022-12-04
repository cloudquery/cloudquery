// Code generated by codegen; DO NOT EDIT.

package armoperationalinsights

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ManagementGroup() *schema.Table {
	return &schema.Table{
		Name:      "azure_armoperationalinsights_management_group",
		Resolver:  fetchManagementGroup,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
		},
	}
}

func fetchManagementGroup(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmoperationalinsightsManagementGroup
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
