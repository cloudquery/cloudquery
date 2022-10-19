package docdb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DocDB

	input := &docdb.DescribeDBClusterParameterGroupsInput{}

	for {
		output, err := svc.DescribeDBClusterParameterGroups(ctx, input)
		if err != nil {
			return err
		}
		if len(output.DBClusterParameterGroups) == 0 {
			return nil
		}
		res <- output.DBClusterParameterGroups
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveDBClusterParameterGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterParameterGroup)
	return resolveDocDBTags(ctx, meta, resource, *item.DBClusterParameterGroupArn, c.Name)
}
