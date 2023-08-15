package stepfunctions

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("StateMachineArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveStepFunctionTags,
			},
		},
		Relations: []*schema.Table{
			executions(),
		},
	}
}

func fetchStepfunctionsStateMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sfn
	config := sfn.ListStateMachinesInput{}
	paginator := sfn.NewListStateMachinesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *sfn.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.StateMachines
	}
	return nil
}

func getStepFunction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sfn
	sm := resource.Item.(types.StateMachineListItem)

	stateMachineDetails, err := svc.DescribeStateMachine(ctx,
		&sfn.DescribeStateMachineInput{StateMachineArn: sm.StateMachineArn},
		func(o *sfn.Options) {
			o.Region = cl.Region
		},
	)
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
	tags, err := svc.ListTagsForResource(ctx, &tagParams, func(o *sfn.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tags.Tags))
}
