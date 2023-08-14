package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func TopLevelDomains() *schema.Table {
	return &schema.Table{
		Name:                 "azure_appservice_top_level_domains",
		Resolver:             fetchTopLevelDomains,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/appservice/top-level-domains/list?tabs=HTTP#topleveldomain",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_top_level_domains", client.Namespacemicrosoft_domainregistration),
		Transform:            transformers.TransformWithStruct(&armappservice.TopLevelDomain{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
