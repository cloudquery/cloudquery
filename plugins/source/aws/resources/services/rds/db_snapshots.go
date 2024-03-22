package rds

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/rds/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func DbSnapshots() *schema.Table {
	tableName := "aws_rds_db_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSnapshot.html`,
		Resolver:    fetchRdsDbSnapshots,
		Transform:   transformers.TransformWithStruct(&types.DBSnapshot{}, transformers.WithSkipFields("TagList", "Attributes"), transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DBSnapshotArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTagPath("TagList"),
			},
			{
				Name:     "attributes",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Attributes"),
			},
		},
	}
}

func fetchRdsDbSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	var input rds.DescribeDBSnapshotsInput
	paginator := rds.NewDescribeDBSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return nil
		}
		for _, DBSnapshot := range page.DBSnapshots {
			attributes, err := fetchSnapshotAttributes(ctx, meta, DBSnapshot.DBSnapshotIdentifier)
			if err != nil {
				// Log error and continue resolving
				cl.Logger().Warn().Err(err).Msg("failed to fetch snapshot attributes")
				res <- models.ExtendedSnapshots{
					DBSnapshot: DBSnapshot,
				}
				continue
			}
			newAttributes := make([]models.ExtendedAttributes, len(attributes))
			for _, attribute := range attributes {
				newAttributes = append(newAttributes, models.ExtendedAttributes{DBSnapshotAttribute: attribute})
			}
			res <- models.ExtendedSnapshots{
				DBSnapshot: DBSnapshot,
				Attributes: newAttributes,
			}
		}
	}
	return nil
}

func fetchSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, dbSnapshotIdentifier *string) ([]types.DBSnapshotAttribute, error) {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	out, err := svc.DescribeDBSnapshotAttributes(
		ctx,
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: dbSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	if out.DBSnapshotAttributesResult == nil {
		return nil, nil
	}
	return out.DBSnapshotAttributesResult.DBSnapshotAttributes, nil
}
