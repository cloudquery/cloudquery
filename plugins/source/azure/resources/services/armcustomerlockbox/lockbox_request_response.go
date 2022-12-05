// Code generated by codegen; DO NOT EDIT.

package armcustomerlockbox

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LockboxRequestResponse() *schema.Table {
	return &schema.Table{
		Name:      "azure_armcustomerlockbox_lockbox_request_response",
		Resolver:  fetchLockboxRequestResponse,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchLockboxRequestResponse(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armcustomerlockbox.NewRequestsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(cl.ResourceGroup, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
