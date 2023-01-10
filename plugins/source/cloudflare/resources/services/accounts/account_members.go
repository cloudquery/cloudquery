package accounts

import (
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AccountMembers() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_account_members",
		Resolver:  fetchAccountMembers,
		Transform: transformers.TransformWithStruct(&cloudflare.AccountMember{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
