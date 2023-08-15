package databases

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Databases() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_databases",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases",
		Resolver:    fetchDatabasesDatabases,
		Transform:   transformers.TransformWithStruct(&godo.Database{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			firewallRules(),
			replicas(),
			backups(),
		},
	}
}
