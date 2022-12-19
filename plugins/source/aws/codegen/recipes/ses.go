package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SESResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "configuration_sets",
			Struct:              &sesv2.GetConfigurationSetOutput{},
			Description:         "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetConfigurationSet.html",
			SkipFields:          []string{"ResultMetadata"},
			PreResourceResolver: "getConfigurationSet",
			ExtraColumns: append(
				defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `resolveConfigurationSetArn`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			),
			NameTransformer: CreateReplaceTransformer(map[string]string{
				"configuration_set_name": "name",
			}),
			Relations: []string{"ConfigurationSetEventDestinations()"},
		},
		{
			SubService:  "configuration_set_event_destinations",
			Struct:      types.EventDestination{},
			Description: "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html",
			PKColumns:   []string{"name"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				codegen.ColumnDefinition{
					Name:     "configuration_set_name",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("name")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			),
		},
		{
			SubService:          "contact_lists",
			Struct:              &sesv2.GetContactListOutput{},
			Description:         "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html",
			PreResourceResolver: "getContactList",
			PKColumns:           []string{"name"},
			SkipFields:          []string{"ResultMetadata"},
			ExtraColumns:        defaultRegionalColumnsPK,
			NameTransformer: CreateReplaceTransformer(map[string]string{
				"contact_list_name": "name",
			}),
		},
		{
			SubService:            "identities",
			Struct:                &models.EmailIdentity{},
			UnwrapEmbeddedStructs: true,
			Description:           "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailIdentity.html",
			PreResourceResolver:   "getIdentity",
			SkipFields:            []string{"ResultMetadata"},
			ExtraColumns: append(
				defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: "resolveIdentityArn",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			),
		},
		{
			SubService:          "templates",
			Struct:              &models.Template{},
			PreResourceResolver: "getTemplate",
			ExtraColumns: append(
				defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `resolveTemplateArn`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ses"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("email")`
	}
	return resources
}
