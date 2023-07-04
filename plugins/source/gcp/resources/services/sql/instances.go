package sql

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/sqladmin/v1beta4"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_sql_instances",
		Description: `https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/instances#DatabaseInstance`,
		Resolver:    fetchInstances,
		Multiplex:   client.ProjectMultiplexEnabledServices("sqladmin.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.DatabaseInstance{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
		Relations: []*schema.Table{
			Users(),
		},
	}
}
