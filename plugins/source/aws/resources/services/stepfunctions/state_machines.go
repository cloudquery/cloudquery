package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func StateMachines() *schema.Table {
	tableName := "aws_stepfunctions_state_machines"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html`,
		Resolver:            fetchStepfunctionsStateMachines,
		PreResourceResolver: getStepFunction,
		Transform:           transformers.TransformWithStruct(&sfn.DescribeStateMachineOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "states"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateMachineArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveStepFunctionTags,
			},
		},
		Relations: []*schema.Table{
			executions(),
		},
	}
}

func fetchStepfunctionsStateMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.ListStateMachinesInput{}
	paginator := sfn.NewListStateMachinesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.StateMachines
	}
	return nil
}

func getStepFunction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sfn
	sm := resource.Item.(types.StateMachineListItem)

	stateMachineDetails, err := svc.DescribeStateMachine(ctx, &sfn.DescribeStateMachineInput{StateMachineArn: sm.StateMachineArn})
	if err != nil {
		return err
	}

	resource.Item = stateMachineDetails
	return nil
}

func resolveStepFunctionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sm := resource.Item.(*sfn.DescribeStateMachineOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Sfn
	tagParams := sfn.ListTagsForResourceInput{
		ResourceArn: sm.StateMachineArn,
	}
	tags, err := svc.ListTagsForResource(ctx, &tagParams)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tags.Tags))
}
