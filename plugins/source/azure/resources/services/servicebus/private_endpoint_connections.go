package servicebus

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/gorilla/mux"
)

func privateEndpointConnections() *schema.Table {
	return &schema.Table{
		Name:        "azure_servicebus_namespace_private_endpoint_connections",
		Resolver:    fetchPrivateEndpointConnections,
		Description: "https://learn.microsoft.com/en-us/rest/api/servicebus/stable/private-endpoint-connections/list?tabs=HTTP",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_servicebus_namespaces", client.Namespacemicrosoft_servicebus),
		Transform:   transformers.TransformWithStruct(&armservicebus.PrivateEndpointConnection{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	namespace := parent.Item.(*armservicebus.SBNamespace)
	svc, err := armservicebus.NewPrivateEndpointConnectionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*namespace.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(group, *namespace.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

func createPrivateEndpointConnections(router *mux.Router) error {
	var item armservicebus.PrivateEndpointConnectionsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/privateEndpointConnections", func(w http.ResponseWriter, r *http.Request) {
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
