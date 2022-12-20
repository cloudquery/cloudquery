package quicksight

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchQuicksightAnalyses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	input := quicksight.ListAnalysesInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	var ae smithy.APIError

	paginator := quicksight.NewListAnalysesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}
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
