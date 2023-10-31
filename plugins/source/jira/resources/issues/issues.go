package issues

import (
	"context"

	"github.com/andygrunwald/go-jira"
	"github.com/cloudquery/cloudquery/plugins/source/jira/sync"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Issues() *schema.Table {
	return &schema.Table{
		Name:        "jira_issues",
		Description: "This table shows data for Jira Issues.",
		Transform:   transformers.TransformWithStruct(&jira.Issue{}, transformers.WithPrimaryKeys("Self")),
		Resolver:    fetchResolver,
	}
}

func fetchResolver(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*sync.Client)
	startAt := 0
	for {
		issueList, resp, err := c.Jira.Issue.SearchWithContext(ctx, "", &jira.SearchOptions{
			StartAt:    startAt,
			MaxResults: 1000,
		})
		if err != nil {
			return err
		}
		res <- issueList

		if resp.Total <= resp.StartAt+resp.MaxResults {
			break
		}
		startAt = resp.StartAt + resp.MaxResults
	}
	return nil
}
