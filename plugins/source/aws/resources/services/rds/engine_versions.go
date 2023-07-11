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

func EngineVersions() *schema.Table {
	tableName := "aws_rds_engine_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html`,
		Resolver:    fetchRdsEngineVersions,
		Transform:   transformers.TransformWithStruct(&types.DBEngineVersion{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "_engine_version_hash",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveObjectHash,
				PrimaryKey: true,
			},
			{
				Name:     "tag_list",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTagField("TagList"),
			},
		},

		Relations: []*schema.Table{
			clusterParameters(),
		},
	}
}

func fetchRdsEngineVersions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	input := &rds.DescribeDBEngineVersionsInput{}
	p := rds.NewDescribeDBEngineVersionsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.DBEngineVersions
	}
	return nil
}
