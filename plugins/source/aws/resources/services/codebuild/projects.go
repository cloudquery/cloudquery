// Code generated by codegen; DO NOT EDIT.

package codebuild

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "aws_codebuild_projects",
		Resolver:  fetchCodebuildProjects,
		Multiplex: client.ServiceAccountRegionMultiplexer("codebuild"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "artifacts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Artifacts"),
			},
			{
				Name:     "badge",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Badge"),
			},
			{
				Name:     "build_batch_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BuildBatchConfig"),
			},
			{
				Name:     "cache",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cache"),
			},
			{
				Name:     "concurrent_build_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ConcurrentBuildLimit"),
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "encryption_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionKey"),
			},
			{
				Name:     "environment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Environment"),
			},
			{
				Name:     "file_system_locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FileSystemLocations"),
			},
			{
				Name:     "last_modified",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModified"),
			},
			{
				Name:     "logs_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogsConfig"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "project_visibility",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProjectVisibility"),
			},
			{
				Name:     "public_project_alias",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicProjectAlias"),
			},
			{
				Name:     "queued_timeout_in_minutes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("QueuedTimeoutInMinutes"),
			},
			{
				Name:     "resource_access_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceAccessRole"),
			},
			{
				Name:     "secondary_artifacts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecondaryArtifacts"),
			},
			{
				Name:     "secondary_source_versions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecondarySourceVersions"),
			},
			{
				Name:     "secondary_sources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecondarySources"),
			},
			{
				Name:     "service_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceRole"),
			},
			{
				Name:     "source",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Source"),
			},
			{
				Name:     "source_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceVersion"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "timeout_in_minutes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TimeoutInMinutes"),
			},
			{
				Name:     "vpc_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcConfig"),
			},
			{
				Name:     "webhook",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Webhook"),
			},
		},
	}
}
