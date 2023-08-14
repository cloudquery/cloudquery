package accounts

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func AccountMembers() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_account_members",
		Resolver:  fetchAccountMembers,
		Transform: client.TransformWithStruct(&cloudflare.AccountMember{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
