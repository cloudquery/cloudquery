package bigtableadmin

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
			client.ProjectIDColumn(true),
			{
				Name:       "instance_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("name"),
				PrimaryKey: true,
			},
		},
	}
}
