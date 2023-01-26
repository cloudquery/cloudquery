package sql

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/sqladmin/v1beta4"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "gcp_sql_users",
		Description: `https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/users#User`,
		Resolver:    fetchUsers,
		Multiplex:   client.ProjectMultiplexEnabledServices("sqladmin.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.User{}, append(client.Options(), transformers.WithPrimaryKeys("Instance", "Name"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
