package databases

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Databases() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_databases",
		Resolver:  fetchDatabasesDatabases,
		Transform: transformers.TransformWithStruct(&godo.Database{}),
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

		Relations: []*schema.Table{
			firewallRules(),
			replicas(),
			backups(),
		},
	}
}
