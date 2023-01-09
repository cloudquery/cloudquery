package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterTasks() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_cluster_tasks",
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Task.html`,
		Resolver:    fetchEcsClusterTasks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ecs"),
		Transform:   transformers.TransformWithStruct(&types.Task{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "task_protection",
				Type:     schema.TypeJSON,
				Resolver: getEcsTaskProtection,
			},
		},
	}
}
