package cdn

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func endpoints() *schema.Table {
	return &schema.Table{
		Name:     "azure_cdn_endpoints",
		Resolver: fetchEndpoints,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}