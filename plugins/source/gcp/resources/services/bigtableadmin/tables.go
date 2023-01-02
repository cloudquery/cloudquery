// Code generated by codegen; DO NOT EDIT.

package bigtableadmin

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigtableadmin_tables",
		Description:         `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.tables#Table`,
		Resolver:            fetchTables,
		PreResourceResolver: getTableInfo,
		Multiplex:           client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
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
				Name:     "families",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Families"),
			},
			{
				Name:     "family_infos",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FamilyInfos"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
		},
	}
}
