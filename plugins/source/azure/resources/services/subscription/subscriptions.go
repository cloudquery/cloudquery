package subscription

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_subscription_subscriptions",
		Resolver:             fetchSubscriptions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/resources/subscriptions/list?tabs=HTTP#subscription",
		Transform:            transformers.TransformWithStruct(&armsubscriptions.Subscription{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
		Relations: []*schema.Table{
			locations(),
		},
	}
}
