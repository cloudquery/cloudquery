package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "db_cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchRdsClusterBacktracks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := rds.DescribeDBClusterBacktracksInput{
		DBClusterIdentifier: parent.Item.(types.DBCluster).DBClusterIdentifier,
	}
	c := meta.(*client.Client)
	svc := c.Services().Rds
	p := rds.NewDescribeDBClusterBacktracksPaginator(svc, &config)
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- resp.DBClusterBacktracks
	}
	return nil
}
