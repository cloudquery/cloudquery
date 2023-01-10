package bigtableadmin

import (
	pb "cloud.google.com/go/bigtable"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigtableadmin_tables",
		Description:         `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.tables#Table`,
		PreResourceResolver: getTableInfo,
		Resolver:            fetchTables,
		Multiplex:           client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:           transformers.TransformWithStruct(&pb.TableInfo{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "instance_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
