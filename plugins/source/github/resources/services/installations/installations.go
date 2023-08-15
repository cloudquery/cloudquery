package installations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func Installations() *schema.Table {
	return &schema.Table{
		Name:      "github_installations",
		Resolver:  fetchInstallations,
		Multiplex: client.OrgMultiplex,
		Transform: client.TransformWithStruct(&github.Installation{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn},
	}
}

func fetchInstallations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{PerPage: 100}
	for {
		installations, resp, err := c.Github.Organizations.ListInstallations(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- installations.Installations

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return nil
}
