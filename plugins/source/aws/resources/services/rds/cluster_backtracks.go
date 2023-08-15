package rds

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterBacktracks() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_cluster_backtracks",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBClusterBacktracks.html`,
		Resolver:    fetchRdsClusterBacktracks,
		Transform: transformers.TransformWithStruct(
			&types.DBClusterBacktrack{},
			transformers.WithPrimaryKeys("BacktrackIdentifier"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "db_cluster_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchRdsClusterBacktracks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.DBCluster)

	if aws.ToInt64(cluster.BacktrackWindow) == 0 {
		// If this value is set to 0, backtracking is disabled for the DB cluster.
		return nil
	}

	config := rds.DescribeDBClusterBacktracksInput{
		DBClusterIdentifier: cluster.DBClusterIdentifier,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	p := rds.NewDescribeDBClusterBacktracksPaginator(svc, &config)
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- resp.DBClusterBacktracks
	}
	return nil
}
