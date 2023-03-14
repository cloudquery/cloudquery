package installations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
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
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
