package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_emr_clusters"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_Cluster.html`,
		Resolver:            fetchEmrClusters,
		PreResourceResolver: getCluster,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:           transformers.TransformWithStruct(&types.Cluster{}),
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
		Relations: []*schema.Table{clusterInstanceFleets(), clusterInstanceGroups(), clusterInstances()},
	}
}

func fetchEmrClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := emr.ListClustersInput{
		ClusterStates: []types.ClusterState{
			types.ClusterStateRunning,
			types.ClusterStateStarting,
			types.ClusterStateBootstrapping,
			types.ClusterStateWaiting,
		},
	}
	c := meta.(*client.Client)
	svc := c.Services().Emr
	paginator := emr.NewListClustersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Clusters
	}
	return nil
}

func getCluster(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Emr
	response, err := svc.DescribeCluster(ctx, &emr.DescribeClusterInput{ClusterId: resource.Item.(types.ClusterSummary).Id})
	if err != nil {
		return err
	}
	resource.Item = response.Cluster
	return nil
}
