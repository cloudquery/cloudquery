package external

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "github_external_groups",
		Resolver:  fetchGroups,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.ExternalGroup{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("GroupID"))...),
		Columns: []schema.Column{client.OrgColumn},
	}
}
