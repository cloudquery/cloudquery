package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func alerts() *schema.Table {
	return &schema.Table{
		Name:      "github_organization_dependabot_alerts",
		Resolver:  fetchAlerts,
		Transform: client.TransformWithStruct(&github.DependabotAlert{}, transformers.WithPrimaryKeys("HTMLURL")),
		Columns:   []schema.Column{client.OrgColumn},
	}
}

func fetchAlerts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListAlertsOptions{ListCursorOptions: github.ListCursorOptions{PerPage: 100}}

	for {
		alerts, resp, err := c.Github.Dependabot.ListOrgAlerts(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- alerts
		opts.After = resp.After
		if len(resp.After) == 0 {
			break
		}
	}

	return nil
}
