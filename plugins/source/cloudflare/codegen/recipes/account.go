package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &cloudflare.Account{},
			PKColumns:  []string{"id"},
			Service:    "accounts",
			Relations:  []string{"AccountMembers()"},
		},
		{
			DataStruct: &cloudflare.AccountMember{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Service: "accounts",
		},
	}
}
