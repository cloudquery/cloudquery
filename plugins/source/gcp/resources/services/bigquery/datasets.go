package bigquery

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	pb "google.golang.org/api/bigquery/v2"
)

type datasetWrapper struct {
	*pb.Dataset
	svc *pb.Service
}

func Datasets() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigquery_datasets",
		Description:         `https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets#Dataset`,
		PreResourceResolver: datasetGet,
		Resolver:            fetchDatasets,
		Multiplex:           client.ProjectMultiplexEnabledServices("bigquery.googleapis.com"),
		Transform:           client.TransformWithStruct(&datasetWrapper{}, transformers.WithPrimaryKeys("Id"), transformers.WithUnwrapStructFields("Dataset")),
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
