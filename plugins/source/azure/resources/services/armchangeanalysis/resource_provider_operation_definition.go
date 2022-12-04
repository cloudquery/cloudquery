// Code generated by codegen; DO NOT EDIT.

package armchangeanalysis

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceProviderOperationDefinition() *schema.Table {
	return &schema.Table{
		Name:      "azure_armchangeanalysis_resource_provider_operation_definition",
		Resolver:  fetchResourceProviderOperationDefinition,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "display",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Display"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
	}
}

func fetchResourceProviderOperationDefinition(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmchangeanalysisResourceProviderOperationDefinition
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
