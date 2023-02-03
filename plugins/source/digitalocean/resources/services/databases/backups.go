package databases

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func backups() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_database_backups",
		Resolver:  fetchDatabasesBackups,
		Transform: transformers.TransformWithStruct(&godo.DatabaseBackup{}),
		Columns:   []schema.Column{},
	}
}
