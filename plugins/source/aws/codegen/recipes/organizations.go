package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func OrganizationsResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "accounts",
			Struct:      &types.Account{},
			Description: "https://docs.aws.amazon.com/organizations/latest/APIReference/API_Account.html",
			SkipFields:  []string{"Arn"},
			Multiplex:   `client.AccountMultiplex`,
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveAccountTags`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "organizations"
	}
	return resources
}
