package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchStepfunctionsActivities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.ListActivitiesInput{
		MaxResults: 1000,
	}
	paginator := sfn.NewListActivitiesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Activities
	}
	return nil
}
