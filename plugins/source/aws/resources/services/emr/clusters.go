package emr

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ClusterArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
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
	cl := meta.(*client.Client)
	svc := cl.Services().Emr
	paginator := emr.NewListClustersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Clusters
	}
	return nil
}

func getCluster(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Emr
	response, err := svc.DescribeCluster(ctx, &emr.DescribeClusterInput{ClusterId: resource.Item.(types.ClusterSummary).Id}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.Cluster
	return nil
}
