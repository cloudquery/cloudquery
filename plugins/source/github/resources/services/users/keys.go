package users

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v59/github"
)

func keys() *schema.Table {
	return &schema.Table{
		Name:      "github_user_keys",
		Resolver:  fetchMembers,
		Transform: client.TransformWithStruct(&github.Key{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			{
				Name:       "user_id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{PerPage: 100}
	for {
		keys, resp, err := c.Github.Users.ListKeys(ctx, "", opts)
		if err != nil {
			return err
		}
		res <- keys

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return nil
}
