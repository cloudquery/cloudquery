package rds

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DbSnapshots() *schema.Table {
	tableName := "aws_rds_db_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSnapshot.html`,
		Resolver:    fetchRdsDbSnapshots,
		Transform:   transformers.TransformWithStruct(&types.DBSnapshot{}, transformers.WithSkipFields("TagList")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("DBSnapshotArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRDSDBSnapshotTags,
			},
			{
				Name:     "attributes",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRDSDBSnapshotAttributes,
			},
		},
	}
}

func fetchRdsDbSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	var input rds.DescribeDBSnapshotsInput
	paginator := rds.NewDescribeDBSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return nil
		}
		res <- page.DBSnapshots
	}
	return nil
}

func resolveRDSDBSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBSnapshot)
	tags := map[string]*string{}
	for _, t := range s.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}

func resolveRDSDBSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBSnapshot)
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	out, err := svc.DescribeDBSnapshotAttributes(
		ctx,
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: s.DBSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if out.DBSnapshotAttributesResult == nil {
		return nil
	}

	return resource.Set(column.Name, out.DBSnapshotAttributesResult.DBSnapshotAttributes)
}
