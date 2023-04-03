package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func orderableDbInstanceOptions() *schema.Table {
	tableName := "aws_docdb_orderable_db_instance_options"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_OrderableDBInstanceOption.html`,
		Resolver:    fetchDocdbOrderableDbInstanceOptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.OrderableDBInstanceOption{}),
		Columns:     []schema.Column{},
	}
}

func fetchDocdbOrderableDbInstanceOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	item := parent.Item.(types.DBEngineVersion)
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := &docdb.DescribeOrderableDBInstanceOptionsInput{Engine: item.Engine}

	p := docdb.NewDescribeOrderableDBInstanceOptionsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.OrderableDBInstanceOptions
	}
	return nil
}
