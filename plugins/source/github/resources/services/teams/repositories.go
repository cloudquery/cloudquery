package teams

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func repositories() *schema.Table {
	return &schema.Table{
		Name:      "github_team_repositories",
		Resolver:  fetchRepositories,
		Transform: client.TransformWithStruct(&github.Repository{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn, teamIDColumn},
	}
}

func fetchRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	t := parent.Item.(*github.Team)
	c := meta.(*client.Client)
	opts := &github.ListOptions{PerPage: 100}
	orgId, err := strconv.Atoi(strings.Split(*t.MembersURL, "/")[4])
	if err != nil {
		return err
	}
	for {
		repos, resp, err := c.Github.Teams.ListTeamReposByID(ctx, int64(orgId), *t.ID, opts)
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
