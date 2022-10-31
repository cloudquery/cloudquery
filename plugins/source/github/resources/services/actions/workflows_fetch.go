package actions

import (
	"context"
	"net/url"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"sigs.k8s.io/yaml"
)

func fetchWorkflows(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
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
	client := meta.(*client.Client)
	workflow := resource.Item.(*github.Workflow)

	url, err := url.Parse(*workflow.URL)
	if err != nil {
		return err
	}

	pathParts := strings.Split(url.Path, "/")
	if len(pathParts) < 2 {
		return nil
	}
	opts := github.RepositoryContentGetOptions{}

	fileContent, _, _, err := client.Github.Repositories.GetContents(ctx, client.Org, pathParts[2], *workflow.Path, &opts)
	if err != nil {
		// This is not actually an error, it means that a workflow file has been deleted
		return err
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return err
	}
	jsonContent, err := yaml.YAMLToJSON([]byte(content))
	if err != nil {
		return err
	}
	return resource.Set(c.Name, jsonContent)
}
