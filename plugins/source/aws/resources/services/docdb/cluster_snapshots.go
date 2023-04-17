package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func clusterSnapshots() *schema.Table {
	tableName := "aws_docdb_cluster_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterSnapshot.html`,
		Resolver:    fetchDocdbClusterSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.DBClusterSnapshot{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDBClusterSnapshotTags,
			},
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
				Resolver: resolveDocdbClusterSnapshotAttributes,
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_cluster_snapshot_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterSnapshotIdentifier"),
			},
		},
	}
}

func fetchDocdbClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	item := parent.Item.(types.DBCluster)
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := &docdb.DescribeDBClusterSnapshotsInput{
		DBClusterIdentifier: item.DBClusterIdentifier,
	}
	p := docdb.NewDescribeDBClusterSnapshotsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.DBClusterSnapshots
	}
	return nil
}

func resolveDocdbClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterSnapshot)
	cli := meta.(*client.Client)
	svc := cli.Services().Docdb

	input := &docdb.DescribeDBClusterSnapshotAttributesInput{
		DBClusterSnapshotIdentifier: item.DBClusterSnapshotIdentifier,
	}

	output, err := svc.DescribeDBClusterSnapshotAttributes(ctx, input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.DBClusterSnapshotAttributesResult.DBClusterSnapshotAttributes)
}

func resolveDBClusterSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterSnapshot)
	return resolveDocDBTags(ctx, meta, resource, *item.DBClusterSnapshotArn, c.Name)
}
