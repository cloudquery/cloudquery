package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	optGroups := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
			Page:    0,
		},
	}
	opt := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
			Page:    0,
		},
	}

	for {
		// Get the first page with projects.
		groups, respList, err := c.Gitlab.Groups.ListGroups(optGroups)
		if err != nil {
			return err
		}
		if len(groups) == 0 {
			return nil
		}

		for _, group := range groups {

			for {
				// Get the first page with projects.
				projects, resp, err := c.Gitlab.Groups.ListGroupProjects(group.ID, opt, gitlab.WithContext(ctx))
				if err != nil {
					return err
				}
				res <- projects

				// Exit the loop when we've seen all pages.
				if resp.NextPage == 0 {
					break
				}

				// Update the page number to get the next page.
				opt.Page = resp.NextPage
			}
		}

		// Exit the loop when we've seen all pages.
		if respList.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = respList.NextPage
	}
	return nil
}
