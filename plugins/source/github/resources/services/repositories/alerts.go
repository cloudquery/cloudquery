package repositories

import (
	"context"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func alerts() *schema.Table {
	return &schema.Table{
		Name:      "github_repository_dependabot_alerts",
		Resolver:  fetchAlerts,
		Transform: client.TransformWithStruct(&github.DependabotAlert{}, transformers.WithPrimaryKeys("Number")),
		Columns:   []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)
	opts := &github.ListAlertsOptions{ListCursorOptions: github.ListCursorOptions{PerPage: 100}}

	for {
		alerts, resp, err := c.Github.Dependabot.ListRepoAlerts(ctx, c.Org, *repo.Name, opts)
		if err != nil {
			return err
		}
		res <- alerts
		opts.Page = strconv.FormatInt(int64(resp.NextPage), 10)
		if resp.NextPage == 0 {
			break
		}
	}

	return nil
}
