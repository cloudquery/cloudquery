package databases

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func replicas() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_database_replicas",
		Resolver:  fetchDatabasesReplicas,
		Transform: transformers.TransformWithStruct(&godo.DatabaseReplica{}),
		Columns:   []schema.Column{},
	}
}
