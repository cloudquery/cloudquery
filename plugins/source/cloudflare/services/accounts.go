package services

import (
	"context"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_accounts",
		Description: "Account represents the root object that owns resources.",
		Resolver:    fetchAccounts,
		Columns: []schema.Column{
			{
				Name:            "id",
				Description:     "The unique universal identifier for a Cloudflare account.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "Cloudflare account name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Cloudflare account type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_on",
				Description: "Creation timestamp of the account.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "settings",
				Description: "Outlines the available options for an account.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "cloudflare_account_members",
				Description: "AccountMember is the definition of a member of an account.",
				Resolver:    fetchAccountMembers,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAccountId,
					},
					{
						Name:        "id",
						Description: "The unique universal identifier for a Cloudflare account member.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name: "code",
						Type: schema.TypeString,
					},
					{
						Name:        "user",
						Description: "All the personal information about a member.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "status",
						Description: "Cloudflare account member status.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "cloudflare_account_member_roles",
						Description: "AccountRole defines the roles that a member can have attached.",
						Resolver:    fetchAccountMemberRoles,
						Columns: []schema.Column{
							{
								Name:        "account_member_cq_id",
								Description: "Unique CloudQuery ID of cloudflare_account_members table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "account_id",
								Description: "The Account ID of the resource.",
								Type:        schema.TypeString,
								Resolver:    client.ResolveAccountId,
							},
							{
								Name:        "id",
								Description: "The unique universal identifier for a Cloudflare account member role.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ID"),
							},
							{
								Name:        "name",
								Description: "Cloudflare account member role name.",
								Type:        schema.TypeString,
							},
							{
								Name:        "description",
								Description: "Cloudflare account member role description.",
								Type:        schema.TypeString,
							},
							{
								Name:        "permissions",
								Description: "Cloudflare account member role permissions.",
								Type:        schema.TypeJSON,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
			return errors.WithStack(err)
		}
		res <- accounts
		if !resp.HasMorePages() {
			break
		}
		opt.Page = resp.Page + 1
	}
	return nil
}
func fetchAccountMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	account := parent.Item.(cloudflare.Account)

	opt := cloudflare.PaginationOptions{
		Page:    1,
		PerPage: client.MaxItemsPerPage,
	}

	for {
		accountMembers, resp, err := svc.ClientApi.AccountMembers(ctx, account.ID, opt)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- accountMembers
		if !resp.HasMorePages() {
			break
		}
		opt.Page = resp.Page + 1
	}
	return nil
}
func fetchAccountMemberRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cloudflare.AccountMember)
	res <- r.Roles
	return nil
}
