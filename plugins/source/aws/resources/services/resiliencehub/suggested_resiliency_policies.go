package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SuggestedResiliencyPolicies() *schema.Table {
	tableName := "aws_resiliencehub_suggested_resiliency_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResiliencyPolicy.html`,
		Resolver:    fetchSuggestedResiliencyPolicies,
		Transform:   transformers.TransformWithStruct(&types.ResiliencyPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), arnColumn("PolicyArn")},
	}
}

func fetchSuggestedResiliencyPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Resiliencehub
	p := resiliencehub.NewListSuggestedResiliencyPoliciesPaginator(svc, &resiliencehub.ListSuggestedResiliencyPoliciesInput{})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.ResiliencyPolicies
	}
	return nil
}
