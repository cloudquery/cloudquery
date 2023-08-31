package account

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchAccountUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		currentUser, err := c.Fastly.GetCurrentUser()
		if err != nil {
			return err
		}
		input := &fastly.ListCustomerUsersInput{
			CustomerID: currentUser.CustomerID,
		}
		r, err := c.Fastly.ListCustomerUsers(input)
		if err != nil {
			return err
		}
		res <- r
		return nil
	}
	return c.RetryOnError(ctx, "fastly_account_users", f)
}
