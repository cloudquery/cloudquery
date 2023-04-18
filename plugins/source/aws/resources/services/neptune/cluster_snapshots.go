package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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

func fetchNeptuneClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Neptune
	input := neptune.DescribeDBClusterSnapshotsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}
	paginator := neptune.NewDescribeDBClusterSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil
		}
		res <- page.DBClusterSnapshots
	}
	return nil
}

func resolveNeptuneClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	c := meta.(*client.Client)
	svc := c.Services().Neptune
	out, err := svc.DescribeDBClusterSnapshotAttributes(
		ctx,
		&neptune.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		func(o *neptune.Options) {
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

func resolveNeptuneClusterSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	out, err := svc.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{ResourceName: s.DBClusterSnapshotArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
