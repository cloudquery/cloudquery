package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchStepfunctionsMapRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.ListMapRunsInput{
		MaxResults:   1000,
		ExecutionArn: parent.Item.(*sfn.DescribeExecutionOutput).ExecutionArn,
	}
	paginator := sfn.NewListMapRunsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.MapRuns
	}
	return nil
}

func getMapRun(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.DescribeMapRunInput{
		MapRunArn: resource.Item.(types.MapRunListItem).MapRunArn,
	}
	output, err := svc.DescribeMapRun(ctx, &config)
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}
