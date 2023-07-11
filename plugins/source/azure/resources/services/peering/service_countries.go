package peering

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ServiceCountries() *schema.Table {
	return &schema.Table{
		Name:                 "azure_peering_service_countries",
		Resolver:             fetchServiceCountries,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/peering/peering-service-countries/list?tabs=HTTP#peeringservicecountry",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_peering_service_countries", client.Namespacemicrosoft_peering),
		Transform:            transformers.TransformWithStruct(&armpeering.ServiceCountry{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchServiceCountries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armpeering.NewServiceCountriesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
