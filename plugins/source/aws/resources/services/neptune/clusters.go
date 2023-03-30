package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_neptune_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-clusters.html#DescribeDBClusters`,
		Resolver:    fetchNeptuneClusters,
		Transform:   transformers.TransformWithStruct(&types.DBCluster{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneClusterTags,
			},
		},
	}
}
