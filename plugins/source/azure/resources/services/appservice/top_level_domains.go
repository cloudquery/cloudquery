package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TopLevelDomains() *schema.Table {
	return &schema.Table{
		Name:        "azure_appservice_top_level_domains",
		Resolver:    fetchTopLevelDomains,
		Description: "https://learn.microsoft.com/en-us/rest/api/appservice/top-level-domains/list?tabs=HTTP#topleveldomain",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_top_level_domains", client.Namespacemicrosoft_domainregistration),
		Transform:   transformers.TransformWithStruct(&armappservice.TopLevelDomain{}),
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

func fetchTopLevelDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappservice.NewTopLevelDomainsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
