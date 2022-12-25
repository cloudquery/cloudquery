package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func OrganizationsResources() []*Resource {
	resources := []*Resource{
		{
			TableDefinition: codegen.TableDefinition{
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
		},
		{
			TableDefinition: codegen.TableDefinition{
				SubService:   "organizations",
				Name:         "aws_organizations",
				Struct:       &types.Organization{},
				Description:  "https://docs.aws.amazon.com/organizations/latest/APIReference/API_Organization.html",
				SkipFields:   []string{"AvailablePolicyTypes"}, // deprecated and misleading field according to docs
				PKColumns:    []string{"account_id", "arn"},
				Multiplex:    `client.AccountMultiplex`,
				ExtraColumns: defaultAccountColumns,
			},
		},
	}

	for _, r := range resources {
		r.Service = "organizations"
	}
	return resources
}
