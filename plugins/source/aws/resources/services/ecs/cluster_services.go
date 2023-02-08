package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterServices() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_cluster_services",
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Service.html`,
		Resolver:    fetchEcsClusterServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ecs"),
		Transform:   transformers.TransformWithStruct(&types.Service{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				RetainOrder: true,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceArn"),
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
		Relations: []*schema.Table{
			ClusterTaskSets(),
		},
	}
}
