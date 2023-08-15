package ecs

import (
	"context"
	"errors"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecs/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("TaskDefinitionArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveEcsTaskDefinitionTags,
			},
		},
	}
}

func fetchEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ecs.ListTaskDefinitionsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
	paginator := ecs.NewListTaskDefinitionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TaskDefinitionArns
	}
	return nil
}

func getTaskDefinition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
	taskArn := resource.Item.(string)

	describeTaskDefinitionOutput, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskArn),
		Include:        []types.TaskDefinitionField{types.TaskDefinitionFieldTags},
	}, func(options *ecs.Options) {
		options.Region = cl.Region
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
