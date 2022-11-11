package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ElasticbeanstalkResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "application_versions",
			Struct:      &types.ApplicationVersionDescription{},
			Description: "https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationVersionDescription.html",
			SkipFields:  []string{"ApplicationVersionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ApplicationVersionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "applications",
			Struct:      &types.ApplicationDescription{},
			Description: "https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationDescription.html",
			SkipFields:  []string{"ApplicationArn", "DateCreated"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ApplicationArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:    "date_created",
						Type:    schema.TypeTimestamp,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "environments",
			Struct:      &types.EnvironmentDescription{},
			Description: "https://docs.aws.amazon.com/elasticbeanstalk/latest/APIReference/API_EnvironmentDescription.html",
			SkipFields:  []string{"EnvironmentId", "EnvironmentArn"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("EnvironmentArn")`,
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `resolveElasticbeanstalkEnvironmentTags`,
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("EnvironmentId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "listeners",
					Type:     schema.TypeJSON,
					Resolver: `resolveElasticbeanstalkEnvironmentListeners`,
				},
			},
			Relations: []string{
				"ConfigurationSettings()",
				"ConfigurationOptions()",
			},
		},
		{
			SubService:  "configuration_settings",
			Struct:      &models.ConfigurationSettingsDescriptionWrapper{},
			Description: "https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationSettingsDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "environment_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
				}...),
		},
		{
			SubService:  "configuration_options",
			Struct:      &models.ConfigurationOptionDescriptionWrapper{},
			Description: "https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "environment_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elasticbeanstalk"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("elasticbeanstalk")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
