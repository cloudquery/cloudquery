package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func OrderableDbInstanceOptions() *schema.Table {
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
