package accounts

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_accounts",
		Resolver:  fetchAccountsAccounts,
		Transform: transformers.TransformWithStruct(&godo.Account{}),
		Columns: []schema.Column{
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UUID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
