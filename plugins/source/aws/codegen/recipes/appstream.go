package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AppstreamResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "app_blocks",
			Struct:      &types.AppBlock{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AppBlock.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "applications",
			Struct:      &types.Application{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Application.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				`ApplicationFleetAssociations()`,
			},
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:  "application_fleet_associations",
			Struct:      &types.ApplicationFleetAssociation{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ApplicationFleetAssociation.html",
			SkipFields:  []string{"ApplicationArn", "FleetName"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "application_arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ApplicationArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "fleet_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("FleetName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:  "directory_configs",
			Struct:      &types.DirectoryConfig{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_DirectoryConfig.html",
			SkipFields:  []string{"DirectoryName"},
			ExtraColumns: append(defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "directory_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DirectoryName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "fleets",
			Struct:      &types.Fleet{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Fleet.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "image_builders",
			Struct:      &types.ImageBuilder{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ImageBuilder.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "images",
			Struct:      &types.Image{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Image.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: true,
			MaxResults:                        25,
		},
		{
			SubService:  "stacks",
			Struct:      &types.Stack{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Stack.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				`StackEntitlements()`,
				`StackUserAssociations()`,
			},
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:  "stack_entitlements",
			Struct:      &types.Entitlement{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Entitlement.html",
			SkipFields:  []string{"StackName", "Name"},
			ExtraColumns: append(defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "stack_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("StackName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Name")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:  "stack_user_associations",
			Struct:      &types.UserStackAssociation{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UserStackAssociation.html",
			SkipFields:  []string{"StackName", "UserName", "AuthenticationType"},
			ExtraColumns: append(defaultRegionalColumnsPK, []codegen.ColumnDefinition{
				{
					Name:     "stack_name",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("StackName")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "user_name",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("UserName")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "authentication_type",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("AuthenticationType")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:  "usage_report_subscriptions",
			Struct:      &types.UsageReportSubscription{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UsageReportSubscription.html",
			SkipFields:  []string{"S3BucketName"},
			ExtraColumns: append(defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "s3_bucket_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("S3BucketName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "users",
			Struct:      &types.User{},
			Description: "https://docs.aws.amazon.com/appstream2/latest/APIReference/API_User.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			ShouldGenerateResolverAndMockTest: false,
		},
	}

	for _, resource := range resources {
		resource.Service = "appstream"
		resource.Multiplex = `client.ServiceAccountRegionMultiplexer("appstream2")`

		resource.ResolverAndMockTestTemplate = "describe_resources_1"
	}

	return resources
}
