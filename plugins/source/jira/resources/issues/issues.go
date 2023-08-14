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
		Name:      "jira_issues",
		Transform: transformers.TransformWithStruct(&jira.Issue{}, transformers.WithPrimaryKeys("Self")),
		Resolver:  fetchResolver,
	}
}

func fetchResolver(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*sync.Client)
	issueList, _, err := c.Jira.Issue.SearchWithContext(ctx, "", nil)
	if err != nil {
		return err
	}
	res <- issueList
	return nil
}
