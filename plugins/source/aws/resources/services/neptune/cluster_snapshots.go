package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterSnapshots() *schema.Table {
	tableName := "aws_neptune_cluster_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-snapshots.html#DescribeDBClusterSnapshots`,
		Resolver:    fetchNeptuneClusterSnapshots,
		Transform:   transformers.TransformWithStruct(&types.DBClusterSnapshot{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterSnapshotArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneClusterSnapshotAttributes,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneClusterSnapshotTags,
			},
		},
	}
}
