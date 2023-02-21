package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Activities() *schema.Table {
	return &schema.Table{
		Name:        "aws_stepfunctions_activities",
		Description: `https://docs.aws.amazon.com/step-functions/latest/apireference/API_ListActivities.html`,
		Resolver:    fetchStepfunctionsActivities,
		Transform:   transformers.TransformWithStruct(&types.ActivityListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("states"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

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
