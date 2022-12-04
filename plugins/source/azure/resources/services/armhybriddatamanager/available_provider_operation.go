// Code generated by codegen; DO NOT EDIT.

package armhybriddatamanager

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AvailableProviderOperation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armhybriddatamanager_available_provider_operation",
		Resolver:  fetchAvailableProviderOperation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "display",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Display"),
			},
			{
				Name:     "origin",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Origin"),
			},
		},
	}
}

func fetchAvailableProviderOperation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmhybriddatamanagerAvailableProviderOperation
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
