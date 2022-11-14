package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchQuicksightIngestions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	item := parent.Item.(types.DataSetSummary)
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	input := quicksight.ListIngestionsInput{
		AwsAccountId: aws.String(cl.AccountID),
		DataSetId:    item.DataSetId,
	}
	paginator := quicksight.NewListIngestionsPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.Ingestions
	}
	return nil
}
