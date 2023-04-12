package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GlobalClusters() *schema.Table {
	tableName := "aws_neptune_global_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-global-dbs.html#GlobalCluster`,
		Resolver:    fetchNeptuneGlobalClusters,
		Transform:   transformers.TransformWithStruct(&types.GlobalCluster{}),
		Multiplex:   client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchNeptuneGlobalClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config neptune.DescribeGlobalClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Neptune
	paginator := neptune.NewDescribeGlobalClustersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.GlobalClusters
	}
	return nil
}
