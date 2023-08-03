package fields

import (
	"context"

	"github.com/andygrunwald/go-jira"
	"github.com/cloudquery/cloudquery/plugins/source/jira/sync"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Fields() *schema.Table {
	return &schema.Table{
		Name:      "jira_fields",
		Transform: transformers.TransformWithStruct(&jira.Field{}, transformers.WithPrimaryKeys("ID")),
		Resolver:  fetchFields,
	}
}

func fetchFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*sync.Client)
	list, _, err := c.Jira.Field.GetListWithContext(ctx)
	if err != nil {
		return err
	}
	res <- list
	return nil
}
