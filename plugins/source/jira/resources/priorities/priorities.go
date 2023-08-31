package priorities

import (
	"context"

	"github.com/andygrunwald/go-jira"
	"github.com/cloudquery/cloudquery/plugins/source/jira/sync"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Priorities() *schema.Table {
	return &schema.Table{
		Name:      "jira_priorities",
		Transform: transformers.TransformWithStruct(&jira.Priority{}, transformers.WithPrimaryKeys("Self")),
		Resolver:  fetchPriorities,
	}
}

func fetchPriorities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*sync.Client)
	list, _, err := c.Jira.Priority.GetListWithContext(ctx)
	if err != nil {
		return err
	}
	res <- list
	return nil
}
