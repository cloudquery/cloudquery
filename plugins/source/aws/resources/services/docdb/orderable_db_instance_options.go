package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func OrderableDbInstanceOptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_orderable_db_instance_options",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_OrderableDBInstanceOption.html`,
		Resolver:    fetchDocdbOrderableDbInstanceOptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform:  transformers.TransformWithStruct(&types.OrderableDBInstanceOption{}),
		Columns: []schema.Column{
		},
	}
}
