package bigtableadmin

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigtableadmin_tables",
		Description:         `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.tables#Table`,
		PreResourceResolver: getTableInfo,
		Resolver:            fetchTables,
		Multiplex:           client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:           client.TransformWithStruct(&tableInfoWithName{}, transformers.WithUnwrapStructFields("TableInfo"), transformers.WithPrimaryKeys("Name")),
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
		},
	}
}
