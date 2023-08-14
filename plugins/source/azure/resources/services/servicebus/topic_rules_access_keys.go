package servicebus

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/gorilla/mux"
)

func topicRuleAccessKeys() *schema.Table {
	return &schema.Table{
		Name:                 "azure_servicebus_namespace_topic_rule_access_keys",
		Resolver:             fetchTopicRuleAccessKeys,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics%20%E2%80%93%20authorization%20rules/list-keys?tabs=HTTP#accesskeys",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_servicebus_namespaces", client.Namespacemicrosoft_servicebus),
		Transform:            transformers.TransformWithStruct(&armservicebus.AccessKeys{}, transformers.WithPrimaryKeys("KeyName")),
		Columns: schema.ColumnList{
			client.SubscriptionID,
			schema.Column{
				Name:       "rule_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchTopicRuleAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	namespace := parent.Parent.Parent.Item.(*armservicebus.SBNamespace)
	topic := parent.Parent.Item.(*armservicebus.SBTopic)
	rule := parent.Item.(*armservicebus.SBAuthorizationRule)
	svc, err := armservicebus.NewTopicsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*rule.ID)
	if err != nil {
		return err
	}
	keys, err := svc.ListKeys(ctx, group, *namespace.Name, *topic.Name, *rule.Name, nil)
	if err != nil {
		return err
	}
	res <- keys.AccessKeys
	return nil
}

func createTopicRuleAccessKeys(router *mux.Router) error {
	var item armservicebus.TopicsClientListKeysResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}/ListKeys", func(w http.ResponseWriter, r *http.Request) {
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
