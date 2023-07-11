package advisor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Suppressions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_advisor_suppressions",
		Description:          "https://learn.microsoft.com/en-us/rest/api/advisor/suppressions/list?tabs=HTTP#suppressioncontractlistresult",
		Resolver:             fetchSuppressions,
		PostResourceResolver: client.LowercaseIDResolver,
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_advisor_suppressions", client.Namespacemicrosoft_advisor),
		Transform:            transformers.TransformWithStruct(&armadvisor.SuppressionContract{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSuppressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armadvisor.NewSuppressionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
