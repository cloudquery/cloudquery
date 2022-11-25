package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountResources() []*Resource {
	return []*Resource{
		{
			DataStruct:       &cloudflare.Account{},
			PKColumns:        []string{"id"},
			TableName:        "cloudflare_accounts",
			TableFuncName:    "Accounts",
			Filename:         "accounts.go",
			Service:          "accounts",
			Relations:        []string{"accountMembers()"},
			ResolverFuncName: "fetchAccounts",
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
			TableName:        "cloudflare_account_members",
			TableFuncName:    "accountMembers",
			Filename:         "account_members.go",
			Service:          "accounts",
			ResolverFuncName: "fetchAccountMembers",
		},
	}
}
