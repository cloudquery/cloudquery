package cognitiveservices

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func accountUsages() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cognitiveservices_account_usages",
		Resolver:             fetchAccountUsages,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/accounts/list-usages?tabs=HTTP#usage",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cognitiveservices_account_usages", client.Namespacemicrosoft_cognitiveservices),
		Transform:            transformers.TransformWithStruct(&armcognitiveservices.Usage{}),
		Columns: schema.ColumnList{
			client.SubscriptionID,
			schema.Column{
				Name:     "account_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchAccountUsages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armcognitiveservices.Account)
	cl := meta.(*client.Client)
	svc, err := armcognitiveservices.NewAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	resp, err := svc.ListUsages(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.Value
	return nil
}
