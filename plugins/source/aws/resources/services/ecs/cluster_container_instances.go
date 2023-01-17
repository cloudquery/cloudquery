package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterContainerInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_cluster_container_instances",
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerInstance.html`,
		Resolver:    fetchEcsClusterContainerInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ecs"),
		Transform:   transformers.TransformWithStruct(&types.ContainerInstance{}),
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
