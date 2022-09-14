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
			Relations:        []string{"accountMembers()"},
			ResolverFuncName: "services.FetchAccounts",
		},
		{
			CFStruct: &cloudflare.AccountMember{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_cq_id",
					Type:     schema.TypeUUID,
					Resolver: "schema.ParentIDResolver",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_account_members",
			TableFuncName:    "accountMembers",
			Filename:         "accounts_account_members.go",
			ResolverFuncName: "services.FetchAccountMembers",
		},
	}
}
