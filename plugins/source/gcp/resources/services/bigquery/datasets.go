package bigquery

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/bigquery/v2"
)

func Datasets() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigquery_datasets",
		Description:         `https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets#Dataset`,
		PreResourceResolver: datasetGet,
		Resolver:            fetchDatasets,
		Multiplex:           client.ProjectMultiplexEnabledServices("bigquery.googleapis.com"),
		Transform:           client.TransformWithStruct(&pb.Dataset{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			Tables(),
		},
	}
}
