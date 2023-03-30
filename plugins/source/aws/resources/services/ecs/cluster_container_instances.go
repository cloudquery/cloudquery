package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterContainerInstances() *schema.Table {
	tableName := "aws_ecs_cluster_container_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerInstance.html`,
		Resolver:    fetchEcsClusterContainerInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   client.TransformWithStruct(&types.ContainerInstance{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
