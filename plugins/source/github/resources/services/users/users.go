package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v59/github"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:     "github_users",
		Resolver: fetchUser,
		// No multiplexer is needed here because we are fetching the current user
		Transform: client.TransformWithStruct(&github.User{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{keys()},
	}
}

func fetchUser(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	if len(c.Spec.AppAuth) > 0 {
		c.Logger().Debug().Msg("getting current user is not supported with app auth")
		return nil
	}
	user, _, err := c.Github.Users.Get(ctx, "")
	if err != nil {
		return err
	}
	res <- user
	return nil
}
