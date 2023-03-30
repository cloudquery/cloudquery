package emr

import (
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	tableName := "aws_emr_clusters"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_Cluster.html`,
		Resolver:            fetchEmrClusters,
		PreResourceResolver: getCluster,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:           client.TransformWithStruct(&types.Cluster{}),
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
