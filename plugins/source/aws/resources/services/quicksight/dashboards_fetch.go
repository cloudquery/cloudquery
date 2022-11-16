package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func fetchQuicksightDashboards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	input := quicksight.ListDashboardsInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	var ae smithy.APIError

	paginator := quicksight.NewListDashboardsPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}
			return err
		}
		res <- result.DashboardSummaryList
	}
	return nil
}
