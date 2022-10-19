package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
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

	for {
		output, err := svc.DescribeDBClusterSnapshots(ctx, input)
		if err != nil {
			return err
		}

		if len(output.DBClusterSnapshots) == 0 {
			return nil
		}
		res <- output.DBClusterSnapshots
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveDBClusterSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterSnapshot)
	return resolveDocDBTags(ctx, meta, resource, *item.DBClusterSnapshotArn, c.Name)
}
