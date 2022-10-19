package docdb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	item := parent.Item.(types.DBClusterParameterGroup)
	c := meta.(*client.Client)
	svc := c.Services().DocDB

	input := &docdb.DescribeDBClusterParametersInput{
		DBClusterParameterGroupName: item.DBClusterParameterGroupName,
	}

	for {
		output, err := svc.DescribeDBClusterParameters(ctx, input)
		if err != nil {
			return err
		}

		if len(output.Parameters) == 0 {
			return nil
		}
		res <- output.Parameters
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}
