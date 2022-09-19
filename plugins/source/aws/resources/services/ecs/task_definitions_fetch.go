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
	return client.ListAndDetailResolver(ctx, meta, res, listEcsTaskDefinitions, ecsTaskDefinitionDetail)
}

func ecsTaskDefinitionDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, detail interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().ECS
	taskArn := detail.(string)
	describeTaskDefinitionOutput, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskArn),
		Include:        []types.TaskDefinitionField{types.TaskDefinitionFieldTags},
	})
	if err != nil {
		errorChan <- err
		return
	}
	if describeTaskDefinitionOutput.TaskDefinition == nil {
		return
	}
	resultsChan <- TaskDefinitionWrapper{
		TaskDefinition: describeTaskDefinitionOutput.TaskDefinition,
		Tags:           describeTaskDefinitionOutput.Tags,
	}
}

func listEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, res chan<- interface{}) error {
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

func resolveEcsTaskDefinitionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]string{}
	for _, a := range r.Tags {
		if a.Key == nil {
			continue
		}
		j[*a.Key] = aws.ToString(a.Value)
	}
	return resource.Set(c.Name, j)
}
