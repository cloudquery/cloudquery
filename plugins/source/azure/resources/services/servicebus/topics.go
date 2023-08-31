package servicebus

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/gorilla/mux"
)

func topics() *schema.Table {
	return &schema.Table{
		Name:                 "azure_servicebus_namespace_topics",
		Resolver:             fetchTopics,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics/list-by-namespace?tabs=HTTP#sbtopic",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_servicebus_namespaces", client.Namespacemicrosoft_servicebus),
		Transform:            transformers.TransformWithStruct(&armservicebus.SBTopic{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations:            []*schema.Table{topicAuthorizationRules()},
	}
}

func fetchTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	namespace := parent.Item.(*armservicebus.SBNamespace)
	svc, err := armservicebus.NewTopicsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*namespace.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByNamespacePager(group, *namespace.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

func createTopics(router *mux.Router) error {
	var item armservicebus.TopicsClientListByNamespaceResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics", func(w http.ResponseWriter, r *http.Request) {
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
