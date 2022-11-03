package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := docdb.DescribeDBClustersInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
	}

	p := docdb.NewDescribeDBClustersPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.DBClusters
	}
	return nil
}

func resolveDBClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBCluster)
	return resolveDocDBTags(ctx, meta, resource, *item.DBClusterArn, c.Name)
}
