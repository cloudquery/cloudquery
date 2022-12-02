package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func fetchProjectsReleases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	project := parent.Item.(*gitlab.Project)
	opt := &gitlab.ListReleasesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}

	for {
		// Get the first page with projects.
		members, resp, err := c.Gitlab.Releases.ListReleases(project.ID, opt)
		if err != nil {
			return err
		}
		res <- members
		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}

	return nil
}
