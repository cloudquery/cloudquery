package resiliencehub

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func describeAppAssessments(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Resiliencehub
	out, err := svc.DescribeAppAssessment(ctx,
		&resiliencehub.DescribeAppAssessmentInput{AssessmentArn: resource.Item.(types.AppAssessmentSummary).AssessmentArn},
	)
	if err != nil {
		return err
	}
	resource.SetItem(out.Assessment)
	return nil
}

func fetchAppAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Resiliencehub
	p := resiliencehub.NewListAppAssessmentsPaginator(svc, &resiliencehub.ListAppAssessmentsInput{AppArn: parent.Item.(*types.App).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- out.AssessmentSummaries
	}
	return nil
}
