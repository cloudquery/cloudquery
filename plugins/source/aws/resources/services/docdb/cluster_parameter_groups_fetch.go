package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := &docdb.DescribeDBClusterParameterGroupsInput{}

	p := docdb.NewDescribeDBClusterParameterGroupsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.DBClusterParameterGroups
	}
	return nil
}

func resolveDocdbClusterParameterGroupParameters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterParameterGroup)
	svc := meta.(*client.Client).Services().Docdb

	input := &docdb.DescribeDBClusterParametersInput{
		DBClusterParameterGroupName: item.DBClusterParameterGroupName,
	}

	var params []types.Parameter
	p := docdb.NewDescribeDBClusterParametersPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		params = append(params, response.Parameters...)
	}
	return resource.Set(c.Name, params)
}

func resolveDBClusterParameterGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterParameterGroup)
	return resolveDocDBTags(ctx, meta, resource, *item.DBClusterParameterGroupArn, c.Name)
}
