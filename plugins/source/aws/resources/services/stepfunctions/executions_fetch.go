package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchStepfunctionsExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.ListExecutionsInput{
		MaxResults: 1000,
	}
	paginator := sfn.NewListExecutionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Executions
	}
	return nil
}
