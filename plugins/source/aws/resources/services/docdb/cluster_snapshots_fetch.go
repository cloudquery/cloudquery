package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	item := parent.Item.(types.DBCluster)
	c := meta.(*client.Client)
	svc := c.Services().DocDB

	input := &docdb.DescribeDBClusterSnapshotsInput{
		DBClusterIdentifier: item.DBClusterIdentifier,
	}
	p := docdb.NewDescribeDBClusterSnapshotsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.DBClusterSnapshots
	}
	return nil
}

func resolveDBClusterSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterSnapshot)
	return resolveDocDBTags(ctx, meta, resource, *item.DBClusterSnapshotArn, c.Name)
}
