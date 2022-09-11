package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type TaskDefinitionWrapper struct {
	*types.TaskDefinition
	Tags []types.Tag
}

func fetchEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ecs.ListTaskDefinitionsInput
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	for {
		listClustersOutput, err := svc.ListTaskDefinitions(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}

		for _, taskDefinitionArn := range listClustersOutput.TaskDefinitionArns {
			res <- taskDefinitionArn
		}
		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}

func getEcsTaskDefinition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error{
	c := meta.(*client.Client)
	svc := c.Services().ECS
	taskArn := resource.Item.(string)
	output, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskArn),
		Include:        []types.TaskDefinitionField{types.TaskDefinitionFieldTags},
	})
	if err != nil {
		return err
	}
	resource.Item = TaskDefinitionWrapper{
		output.TaskDefinition,
		output.Tags,
	} 
	return nil
}

