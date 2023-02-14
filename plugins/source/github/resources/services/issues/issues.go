package issues

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Issues() *schema.Table {
	return &schema.Table{
		Name:      "github_issues",
		Resolver:  fetchIssues,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.Issue{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{client.OrgColumn},
	}
}
