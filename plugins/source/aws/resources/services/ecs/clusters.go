package ecs

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	tableName := "aws_ecs_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchEcsClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   client.TransformWithStruct(&types.Cluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArn"),
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
			ClusterTasks(),
			ClusterServices(),
			ClusterContainerInstances(),
		},
	}
}
