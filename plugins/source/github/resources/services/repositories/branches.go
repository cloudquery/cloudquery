package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/go-github/v49/github"
)

func branches() *schema.Table {
	return &schema.Table{
		Name:      "github_repository_branches",
		Resolver:  fetchBranches,
		Transform: client.TransformWithStruct(&github.Branch{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.OrgColumn,
			client.RepositoryIDColumn,
			{
				Name:     "protection",
				Type:     types.ExtensionTypes.JSON,
				Resolver: resolveBranchProtection,
			},
		},
	}
}

func fetchBranches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)
	opts := &github.BranchListOptions{ListOptions: github.ListOptions{PerPage: 100}}

	for {
		branches, resp, err := c.Github.Repositories.ListBranches(ctx, c.Org, *repo.Name, opts)
		if err != nil {
			return err
		}
		res <- branches
		opts.ListOptions.Page = resp.NextPage
		if opts.ListOptions.Page == resp.LastPage {
			break
		}
	}

	return nil
}

func resolveBranchProtection(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	branch := resource.Item.(*github.Branch)
	repo := resource.Parent.Item.(*github.Repository)

	if !*branch.Protected {
		return nil
	}

	protection, _, err := cl.Github.Repositories.GetBranchProtection(ctx, cl.Org, *repo.Name, *branch.Name)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, protection)
}
