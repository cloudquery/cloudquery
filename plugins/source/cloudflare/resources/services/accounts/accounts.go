package accounts

import (
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_accounts",
		Resolver:  fetchAccounts,
		Transform: transformers.TransformWithStruct(&cloudflare.Account{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			AccountMembers(),
		},
	}
}
