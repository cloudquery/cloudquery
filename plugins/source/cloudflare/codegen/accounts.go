// Code generated by codegen using template resource_manual.go.tpl; DO NOT EDIT.

package codegen

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/services"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:     "cloudflare_accounts",
		Resolver: services.FetchAccounts,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "created_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOn"),
			},
			{
				Name:     "settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Settings"),
			},
		},

		Relations: []*schema.Table{
			accountMembers(),
		},
	}
}
