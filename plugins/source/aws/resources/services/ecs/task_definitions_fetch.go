package ecs

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecs/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ecs.ListTaskDefinitionsInput
	svc := meta.(*client.Client).Services().Ecs
	for {
		listClustersOutput, err := svc.ListTaskDefinitions(ctx, &config)
		if err != nil {
			return err
		}
		res <- listClustersOutput.TaskDefinitionArns

		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}

func getTaskDefinition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Ecs
	taskArn := resource.Item.(string)

	describeTaskDefinitionOutput, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskArn),
		Include:        []types.TaskDefinitionField{types.TaskDefinitionFieldTags},
	})
	if err != nil {
		return err
	}
	if describeTaskDefinitionOutput.TaskDefinition == nil {
		return errors.New("nil TaskDefinition encountered")
	}
	resource.Item = models.TaskDefinitionWrapper{
		TaskDefinition: describeTaskDefinitionOutput.TaskDefinition,
		Tags:           describeTaskDefinitionOutput.Tags,
	}
	return nil
}

func resolveEcsTaskDefinitionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(models.TaskDefinitionWrapper)
	return resource.Set(c.Name, client.TagsToMap(r.Tags))
}
