package ecs

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecs/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func TaskDefinitions() *schema.Table {
	tableName := "aws_ecs_task_definitions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html`,
		Resolver:            fetchEcsTaskDefinitions,
		PreResourceResolver: getTaskDefinition,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:           transformers.TransformWithStruct(&types.TaskDefinition{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskDefinitionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEcsTaskDefinitionTags,
			},
		},
	}
}

func fetchEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ecs.ListTaskDefinitionsInput
	svc := meta.(*client.Client).Services().Ecs
	paginator := ecs.NewListTaskDefinitionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.TaskDefinitionArns
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
