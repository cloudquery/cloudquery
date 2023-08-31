package accounts

import (
	"context"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	opt := cloudflare.AccountsListParams{
		PaginationOptions: cloudflare.PaginationOptions{
			Page:    1,
			PerPage: client.MaxItemsPerPage,
		},
	}

	for {
		accounts, resp, err := svc.ClientApi.Accounts(ctx, opt)
		if err != nil {
			return err
		}
		res <- accounts
		if !resp.HasMorePages() {
			break
		}
		opt.Page = resp.Page + 1
	}
	return nil
}
func fetchAccountMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	account := parent.Item.(cloudflare.Account)

	opt := cloudflare.PaginationOptions{
		Page:    1,
		PerPage: client.MaxItemsPerPage,
	}

	for {
		accountMembers, resp, err := svc.ClientApi.AccountMembers(ctx, account.ID, opt)
		if err != nil {
			return err
		}
		res <- accountMembers
		if !resp.HasMorePages() {
			break
		}
		opt.Page = resp.Page + 1
	}
	return nil
}
