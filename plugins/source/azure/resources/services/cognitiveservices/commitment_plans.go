package cognitiveservices

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func CommitmentPlans() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cognitiveservices_commitment_plans",
		Resolver:             fetchCommitmentPlans,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/commitment-plans/list?tabs=HTTP#commitmentplan",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cognitiveservices_commitment_plans", client.Namespacemicrosoft_cognitiveservices),
		Transform:            transformers.TransformWithStruct(&armcognitiveservices.CommitmentPlan{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchCommitmentPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcognitiveservices.NewCommitmentPlansClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPlansBySubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
