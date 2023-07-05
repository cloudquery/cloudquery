package accounts

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_accounts",
		Resolver:  fetchAccounts,
		Transform: client.TransformWithStruct(&cloudflare.Account{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			AccountMembers(),
		},
	}
}
