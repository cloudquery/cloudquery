package actions

import (
	"context"
	"net/url"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func fetchWorkflows(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		repos, resp, err := c.Github.Repositories.ListByOrg(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		for _, repo := range repos {
			actionOpts := &github.ListOptions{PerPage: 100}
			for {
				workflows, resp, err := c.Github.Actions.ListWorkflows(ctx, *repo.Owner.Login, *repo.Name, actionOpts)
				if err != nil {
					return err
				}
				res <- workflows.Workflows
				opts.Page = resp.NextPage
				if opts.Page == resp.LastPage {
					break
				}
			}
		}
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}

func resolveContents(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	workflow := resource.Item.(*github.Workflow)

	parsedUrl, err := url.Parse(*workflow.HTMLURL)
	if err != nil {
		return err
	}

	pathParts := strings.Split(parsedUrl.Path, "/")
	if len(pathParts) < 2 {
		return nil
	}
	owner := pathParts[1]
	repo := pathParts[2]
	ref := pathParts[4]
	path := *workflow.Path
	opts := github.RepositoryContentGetOptions{Ref: ref}

	fileContent, _, _, err := cl.Github.Repositories.GetContents(ctx, owner, repo, path, &opts)
	if err != nil {
		// This is not actually an error, it means that a workflow file has been deleted
		return resource.Set(c.Name, nil)
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return err
	}
	return resource.Set(c.Name, content)
}
