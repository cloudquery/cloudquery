package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Events() *schema.Table {
	tableName := "aws_docdb_events"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html`,
		Resolver:    fetchDocdbEvents,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.Event{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
