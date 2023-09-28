package cloudsupport

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	pb "google.golang.org/api/cloudsupport/v2beta"
)

func Cases() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudsupport_cases",
		Description: `https://cloud.google.com/support/docs/reference/rest/v2beta/cases#Case`,
		Resolver:    fetchCases,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudsupport.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Case{}, transformers.WithPrimaryKeys("Name")),
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
