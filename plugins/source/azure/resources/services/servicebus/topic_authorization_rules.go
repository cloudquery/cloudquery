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

func topicAuthorizationRules() *schema.Table {
	return &schema.Table{
		Name:                 "azure_servicebus_namespace_topic_authorization_rules",
		Resolver:             fetchTopicAuthorizationRules,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics%20%E2%80%93%20authorization%20rules/list-authorization-rules?tabs=HTTP#sbauthorizationrule",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_servicebus_namespaces", client.Namespacemicrosoft_servicebus),
		Transform:            transformers.TransformWithStruct(&armservicebus.SBAuthorizationRule{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations:            []*schema.Table{topicRuleAccessKeys()},
	}
}

func fetchTopicAuthorizationRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	namespace := parent.Parent.Item.(*armservicebus.SBNamespace)
	topic := parent.Item.(*armservicebus.SBTopic)
	svc, err := armservicebus.NewTopicsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*namespace.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListAuthorizationRulesPager(group, *namespace.Name, *topic.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

func createTopicAuthorizationRules(router *mux.Router) error {
	var item armservicebus.TopicsClientListAuthorizationRulesResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules", func(w http.ResponseWriter, r *http.Request) {
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
