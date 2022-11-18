package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "alternate_contacts",
			Struct:      &types.AlternateContact{},
			Description: "https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html",
			PKColumns:   []string{"alternate_contact_type"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:  "contacts",
			Struct:      &types.ContactInformation{},
			Description: "https://docs.aws.amazon.com/accounts/latest/reference/API_ContactInformation.html",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}

	for _, r := range resources {
		r.Service = "account"
		r.Multiplex = `client.AccountMultiplex`

	}
	return resources
}
