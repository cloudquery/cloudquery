package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	item := parent.Item.(types.DBCluster)
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := &docdb.DescribeDBInstancesInput{Filters: []types.Filter{{Name: aws.String("db-cluster-id"), Values: []string{*item.DBClusterIdentifier}}}}

	p := docdb.NewDescribeDBInstancesPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.DBInstances
	}
	return nil
}

func resolveDBInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBInstance)
	return resolveDocDBTags(ctx, meta, resource, *item.DBInstanceArn, c.Name)
}
