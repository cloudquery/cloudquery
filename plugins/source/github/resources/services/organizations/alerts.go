package organizations

import (
	"context"

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

func fetchAlerts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	alerts, _, err := c.Github.Dependabot.ListOrgAlerts(ctx, c.Org, nil)
	if err != nil {
		return err
	}

	res <- alerts

	return nil
}
