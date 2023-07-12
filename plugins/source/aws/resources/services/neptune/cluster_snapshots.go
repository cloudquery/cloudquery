package neptune

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("DBClusterSnapshotArn"),
				PrimaryKey: true,
			},
			{
				Name:     "attributes",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveNeptuneClusterSnapshotAttributes,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveNeptuneClusterSnapshotTags,
			},
		},
	}
}

func fetchNeptuneClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	input := neptune.DescribeDBClusterSnapshotsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}
	paginator := neptune.NewDescribeDBClusterSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *neptune.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return nil
		}
		res <- page.DBClusterSnapshots
	}
	return nil
}

func resolveNeptuneClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	out, err := svc.DescribeDBClusterSnapshotAttributes(
		ctx,
		&neptune.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		func(o *neptune.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		if cl.IsNotFoundError(err) {
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
	out, err := svc.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{ResourceName: s.DBClusterSnapshotArn}, func(options *neptune.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
