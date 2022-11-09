package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchQuicksightAnalyses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	input := quicksight.ListAnalysesInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	paginator := quicksight.NewListAnalysesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.AnalysisSummaryList
	}
	return nil
}

func getAnalysis(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	item := resource.Item.(types.AnalysisSummary)

	out, err := svc.DescribeAnalysis(ctx, &quicksight.DescribeAnalysisInput{
		AwsAccountId: aws.String(cl.AccountID),
		AnalysisId:   item.AnalysisId,
	})
	if err != nil {
		return err
	}

	resource.Item = out.Analysis
	return nil
}
