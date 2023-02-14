package organizations

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func alerts() *schema.Table {
	return &schema.Table{
		Name:     "github_organization_dependabot_alerts",
		Resolver: fetchAlerts,
		Transform: transformers.TransformWithStruct(&github.DependabotAlert{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("HTMLURL"))...),
		Columns: []schema.Column{client.OrgColumn},
	}
}
