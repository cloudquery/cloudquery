package privatedns

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/gorilla/mux"
)

func virtualNetworkLinks() *schema.Table {
	return &schema.Table{
		Name:                 "azure_privatedns_private_zone_virtual_network_links",
		Resolver:             fetchVirtualNetworkLinks,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/dns/privatedns/virtual-network-links/list?tabs=HTTP#virtualnetworklink",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_privatedns_private_zone_virtual_network_links", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armprivatedns.VirtualNetworkLink{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualNetworkLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	zone := parent.Item.(*armprivatedns.PrivateZone)
	svc, err := armprivatedns.NewVirtualNetworkLinksClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*zone.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(group, *zone.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

func createMockVirtualNetworkLinks(router *mux.Router) error {
	var item armprivatedns.VirtualNetworkLinksClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}
