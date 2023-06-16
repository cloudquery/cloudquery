package accounts

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_accounts",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Account",
		Resolver:    fetchAccountsAccounts,
		Transform:   transformers.TransformWithStruct(&godo.Account{}),
		Columns: []schema.Column{
			{
				Name:       "uuid",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("UUID"),
				PrimaryKey: true,
			},
		},
	}
}
