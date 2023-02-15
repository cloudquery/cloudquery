package organizations

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func members() *schema.Table {
	return &schema.Table{
		Name:     "github_organization_members",
		Resolver: fetchMembers,
		Transform: transformers.TransformWithStruct(&github.User{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{
			client.OrgColumn,
			{
				Name:     "membership",
				Type:     schema.TypeJSON,
				Resolver: resolveMembership,
			},
		},
	}
}
