package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterSnapshots() *schema.Table {
	tableName := "aws_rds_cluster_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterSnapshot.html`,
		Resolver:    fetchRdsClusterSnapshots,
		Transform:   transformers.TransformWithStruct(&types.DBClusterSnapshot{}, transformers.WithSkipFields("TagList")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRDSClusterSnapshotTags,
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: resolveRDSClusterSnapshotAttributes,
			},
		},
	}
}
