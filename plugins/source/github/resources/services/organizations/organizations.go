package organizations

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:      "github_organizations",
		Resolver:  fetchOrganizations,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.Organization{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{alerts(), members(), secrets()},
	}
}
