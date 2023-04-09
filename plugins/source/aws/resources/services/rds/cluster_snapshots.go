package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
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

func fetchRdsClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Rds
	var input rds.DescribeDBClusterSnapshotsInput
	paginator := rds.NewDescribeDBClusterSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil
		}
		res <- page.DBClusterSnapshots
	}
	return nil
}

func resolveRDSClusterSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	return resource.Set(c.Name, client.TagsToMap(s.TagList))
}

func resolveRDSClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	c := meta.(*client.Client)
	svc := c.Services().Rds
	out, err := svc.DescribeDBClusterSnapshotAttributes(
		ctx,
		&rds.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = c.Region
		},
	)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if out.DBClusterSnapshotAttributesResult == nil {
		return nil
	}

	return resource.Set(column.Name, out.DBClusterSnapshotAttributesResult.DBClusterSnapshotAttributes)
}
