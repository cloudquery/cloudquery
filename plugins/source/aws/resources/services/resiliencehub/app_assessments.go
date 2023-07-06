package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func appAssesments() *schema.Table {
	tableName := "aws_resiliencehub_app_assessments"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppAssessment.html`,
		Resolver:            fetchAppAssessments,
		PreResourceResolver: describeAppAssessments,
		Transform:           transformers.TransformWithStruct(&types.AppAssessment{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:             []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARNTop, arnColumn("AssessmentArn")},
		Relations: []*schema.Table{
			appComponentCompliances(),
			appComponentRecommendations(),
			testRecommendations(),
			alarmRecommendations(),
			recommendationTemplates(),
			sopRecommendations(),
		},
	}
}

func describeAppAssessments(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Resiliencehub
	out, err := svc.DescribeAppAssessment(ctx,
		&resiliencehub.DescribeAppAssessmentInput{AssessmentArn: resource.Item.(types.AppAssessmentSummary).AssessmentArn},
		func(options *resiliencehub.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	resource.SetItem(out.Assessment)
	return nil
}

func fetchAppAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Resiliencehub
	p := resiliencehub.NewListAppAssessmentsPaginator(svc, &resiliencehub.ListAppAssessmentsInput{AppArn: parent.Item.(*types.App).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.AssessmentSummaries
	}
	return nil
}
