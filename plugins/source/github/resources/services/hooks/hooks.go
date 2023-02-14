package hooks

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Hooks() *schema.Table {
	return &schema.Table{
		Name:      "github_hooks",
		Resolver:  fetchHooks,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.Hook{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{deliveries()},
	}
}
