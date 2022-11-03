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
			SubService:          "templates",
			Struct:              &models.Template{},
			SkipFields:          []string{},
			PreResourceResolver: "getTemplate",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveSesTemplateArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:          "configuration_sets",
			Struct:              &sesv2.GetConfigurationSetOutput{},
			Description:         "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetConfigurationSet.html",
			SkipFields:          []string{"ConfigurationSetName", "ResultMetadata"},
			PreResourceResolver: "getConfigurationSet",
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ConfigurationSetName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{"ConfigurationSetEventDestinations()"},
		},
		{
			SubService:  "configuration_set_event_destinations",
			Struct:      types.EventDestination{},
			Description: "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html",
			SkipFields:  []string{"Name"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "configuration_set_name",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("name")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Name")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:          "contact_lists",
			Struct:              &sesv2.GetContactListOutput{},
			Description:         "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html",
			PreResourceResolver: "getContactList",
			SkipFields:          []string{"ContactListName", "ResultMetadata"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ContactListName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:            "identities",
			Struct:                &models.EmailIdentityWrapper{},
			UnwrapEmbeddedStructs: true,
			Description:           "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailIdentity.html",
			PreResourceResolver:   "getEmailIdentity",
			SkipFields:            []string{"ARN", "ResultMetadata"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveEmailIdentityArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ses"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("email")`
	}
	return resources
}
