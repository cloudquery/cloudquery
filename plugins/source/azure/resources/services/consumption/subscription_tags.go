package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SubscriptionTags() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_subscription_tags",
		Resolver:             fetchSubscriptionTags,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/tags/get?tabs=HTTP#tagsresult",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_consumption_subscription_tags", client.Namespacemicrosoft_consumption),
		Transform:            transformers.TransformWithStruct(&armconsumption.TagsResult{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchSubscriptionTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewTagsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	scope := "subscriptions/" + cl.SubscriptionId
	resp, err := svc.Get(ctx, scope, nil)
	if err != nil {
		return err
	}
	res <- resp.TagsResult
	return nil
}
