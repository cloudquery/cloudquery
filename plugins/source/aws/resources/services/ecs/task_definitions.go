package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TaskDefinitions() *schema.Table {
	return &schema.Table{
		Name:                "aws_ecs_task_definitions",
		Description:         `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html`,
		Resolver:            fetchEcsTaskDefinitions,
		PreResourceResolver: getTaskDefinition,
		Multiplex:           client.ServiceAccountRegionMultiplexer("ecs"),
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
