package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterTaskSets() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_cluster_task_sets",
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskSet.html`,
		Resolver:    fetchEcsClusterTaskSets,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ecs"),
		Transform:   transformers.TransformWithStruct(&types.TaskSet{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskSetArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
