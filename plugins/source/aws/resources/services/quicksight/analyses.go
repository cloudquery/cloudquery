package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pkg/errors"
)

func Analyses() *schema.Table {
	tableName := "aws_quicksight_analyses"
	return &schema.Table{
		Name:                tableName,
		Description:         "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Analysis.html",
		Resolver:            fetchQuicksightAnalyses,
		PreResourceResolver: getAnalysis,
		Transform:           transformers.TransformWithStruct(&types.Analysis{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns:             []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
	}
}

func fetchQuicksightAnalyses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	input := quicksight.ListAnalysesInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	var ae smithy.APIError

	paginator := quicksight.NewListAnalysesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx, func(options *quicksight.Options) {
			options.Region = cl.Region
		})
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
	}, func(options *quicksight.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = out.Analysis
	return nil
}
