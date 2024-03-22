package rds

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Clusters() *schema.Table {
	tableName := "aws_rds_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBCluster.html`,
		Resolver:    fetchRdsClusters,
		Transform:   transformers.TransformWithStruct(&types.DBCluster{}, transformers.WithSkipFields("TagList")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DBClusterArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTagPath("TagList"),
			},
		},
		Relations: []*schema.Table{
			clusterBacktracks(),
		},
	}
}

func fetchRdsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config rds.DescribeDBClustersInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds

	paginator := rds.NewDescribeDBClustersPaginator(svc, &config)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DBClusters
	}
	return nil
}
