package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_ecs_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchEcsClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   transformers.TransformWithStruct(&types.Cluster{}),
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
			clusterTasks(),
			clusterServices(),
			clusterContainerInstances(),
		},
	}
}

func fetchEcsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ecs.ListClustersInput
	svc := meta.(*client.Client).Services().Ecs
	paginator := ecs.NewListClustersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		if len(page.ClusterArns) == 0 {
			return nil
		}
		describeClusterOutput, err := svc.DescribeClusters(ctx, &ecs.DescribeClustersInput{
			Clusters: page.ClusterArns,
			Include: []types.ClusterField{
				types.ClusterFieldAttachments,
				types.ClusterFieldTags,
				types.ClusterFieldSettings,
				types.ClusterFieldConfigurations,
				types.ClusterFieldStatistics,
			},
		})
		if err != nil {
			return err
		}
		res <- describeClusterOutput.Clusters
	}
	return nil
}
