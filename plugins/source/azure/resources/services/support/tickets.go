package support

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Tickets() *schema.Table {
	return &schema.Table{
		Name:        "azure_support_tickets",
		Resolver:    fetchTickets,
		Description: "https://learn.microsoft.com/en-us/rest/api/support/support-tickets/list?tabs=HTTP#supportticketdetails",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_support_tickets", client.Namespacemicrosoft_support),
		Transform:   transformers.TransformWithStruct(&armsupport.TicketDetails{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchTickets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsupport.NewTicketsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
