package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountResources() []Resource {
	return []Resource{
		{
			CFStruct:         &cloudflare.Account{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_accounts",
			TableFuncName:    "Accounts",
			Filename:         "accounts.go",
			Package:          "accounts",
			Relations:        []string{"accountMembers()"},
			ResolverFuncName: "fetchAccounts",
		},
		{
			CFStruct: &cloudflare.AccountMember{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_account_members",
			TableFuncName:    "accountMembers",
			Filename:         "account_members.go",
			Package:          "accounts",
			ResolverFuncName: "fetchAccountMembers",
		},
	}
}
