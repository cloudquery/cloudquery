package databases

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
