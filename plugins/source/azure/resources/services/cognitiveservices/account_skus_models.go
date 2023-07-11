package cognitiveservices

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func accountModels() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cognitiveservices_account_models",
		Resolver:             fetchAccountModels,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices@v1.3.0#AccountModel",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cognitiveservices_account_models", client.Namespacemicrosoft_cognitiveservices),
		Transform:            transformers.TransformWithStruct(&armcognitiveservices.AccountModel{}),
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

func fetchAccountModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
	pager := svc.NewListModelsPager(group, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
