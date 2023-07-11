package privatedns

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PrivateZones() *schema.Table {
	return &schema.Table{
		Name:                 "azure_privatedns_private_zones",
		Resolver:             fetchPrivateZones,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/dns/privatedns/private-zones/list?tabs=HTTP#privatezone",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_privatedns_private_zones", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armprivatedns.PrivateZone{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPrivateZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armprivatedns.NewPrivateZonesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
