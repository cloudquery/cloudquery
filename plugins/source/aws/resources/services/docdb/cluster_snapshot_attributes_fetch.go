package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	item := parent.Item.(types.DBClusterSnapshot)
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := &docdb.DescribeDBClusterSnapshotAttributesInput{
		DBClusterSnapshotIdentifier: item.DBClusterSnapshotIdentifier,
	}

	output, err := svc.DescribeDBClusterSnapshotAttributes(ctx, input)
	if err != nil {
		return err
	}

	res <- output.DBClusterSnapshotAttributesResult

	return nil
}
