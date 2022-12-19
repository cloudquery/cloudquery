package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SESResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "active_receipt_rule_sets",
			Struct:      &ses.DescribeActiveReceiptRuleSetOutput{},
			Description: `https://docs.aws.amazon.com/ses/latest/APIReference/API_DescribeActiveReceiptRuleSet.html`,
			SkipFields:  []string{"Metadata", "ResultMetadata"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				codegen.ColumnDefinition{
					Name:     "name",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Metadata.Name")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				codegen.ColumnDefinition{
					Name:     "created_timestamp",
					Type:     schema.TypeTimestamp,
					Resolver: `schema.PathResolver("Metadata.CreatedTimestamp")`,
				},
			),
		},
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
			NameTransformer:     CreateReplaceTransformer(map[string]string{"contact_list_name": "name"}),
		},
		{
			SubService:          "custom_verification_email_templates",
			Struct:              &sesv2.GetCustomVerificationEmailTemplateOutput{},
			Description:         "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetCustomVerificationEmailTemplate.html",
			PreResourceResolver: "getCustomVerificationEmailTemplate",
			SkipFields:          []string{"ResultMetadata"},
			ExtraColumns: append(
				defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: "resolveCustomVerificationEmailTemplateArn",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			),
			NameTransformer: CreateReplaceTransformer(map[string]string{"template_": ""}),
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
