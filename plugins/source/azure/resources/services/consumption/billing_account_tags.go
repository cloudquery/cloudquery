package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillingAccountTags() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_account_tags",
		Resolver:             fetchBillingAccountTags,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/tags/get?tabs=HTTP#tagsresult",
		Multiplex:            client.BillingAccountMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.TagsResult{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewTagsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	resp, err := svc.Get(ctx, *cl.BillingAccount.ID, nil)
	if err != nil {
		return err
	}
	res <- resp.TagsResult
	return nil
}
