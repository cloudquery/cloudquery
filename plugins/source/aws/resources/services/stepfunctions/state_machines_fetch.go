package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchStepfunctionsStateMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
