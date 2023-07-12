package teams

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:      "github_teams",
		Resolver:  fetchTeams,
		Multiplex: client.OrgMultiplex,
		Transform: client.TransformWithStruct(&github.Team{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{members(), repositories()},
	}
}

func fetchTeams(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{PerPage: 100}
	for {
		repos, resp, err := c.Github.Teams.ListTeams(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- repos

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}
	return nil
}

var teamIDColumn = schema.Column{
	Name:       "team_id",
	Type:       arrow.PrimitiveTypes.Int64,
	Resolver:   schema.ParentColumnResolver("id"),
	PrimaryKey: true,
}
