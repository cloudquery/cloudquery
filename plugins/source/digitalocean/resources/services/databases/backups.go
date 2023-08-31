package databases

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func backups() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_database_backups",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_list_backups",
		Resolver:    fetchDatabasesBackups,
		Transform:   transformers.TransformWithStruct(&godo.DatabaseBackup{}),
		Columns:     []schema.Column{},
	}
}
