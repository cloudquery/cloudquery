package databases

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
