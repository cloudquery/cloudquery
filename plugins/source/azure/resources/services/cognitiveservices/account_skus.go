package cognitiveservices

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func accountSKUs() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cognitiveservices_account_skus",
		Resolver:             fetchAccountSKUs,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/accounts/list-skus?tabs=HTTP#accountsku",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cognitiveservices_account_skus", client.Namespacemicrosoft_cognitiveservices),
		Transform:            transformers.TransformWithStruct(&armcognitiveservices.AccountSKU{}),
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

func fetchAccountSKUs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
	resp, err := svc.ListSKUs(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.Value
	return nil
}
