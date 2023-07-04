package databases

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func replicas() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_database_replicas",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_list_replicas",
		Resolver:    fetchDatabasesReplicas,
		Transform:   transformers.TransformWithStruct(&godo.DatabaseReplica{}),
		Columns:     []schema.Column{},
	}
}
