package bigquery

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/bigquery/v2"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigquery_tables",
		Description:         `https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#Table`,
		PreResourceResolver: tableGet,
		Resolver:            fetchTables,
		Multiplex:           client.ProjectMultiplexEnabledServices("bigquery.googleapis.com"),
		Transform:           client.TransformWithStruct(&pb.Table{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
