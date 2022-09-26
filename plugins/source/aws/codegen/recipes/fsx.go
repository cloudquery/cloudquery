package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FSXResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "backups",
			Struct:     &types.Backup{},
			SkipFields: []string{"BackupId", "Tags"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("BackupId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `client.ResolveTags`,
				},
			},
		},
		{
			SubService: "data_repository_associations",
			Struct:     &types.DataRepositoryAssociation{},
			SkipFields: []string{"ResourceARN", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService: "data_repository_tasks",
			Struct:     &types.DataRepositoryTask{},
			SkipFields: []string{"ResourceARN", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService: "file_systems",
			Struct:     &types.FileSystem{},
			SkipFields: []string{"AdministrativeActions", "ResourceARN", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
					{
						Name:          "administrative_actions",
						Type:          schema.TypeJSON,
						Resolver:      `schema.PathResolver("AdministrativeActions")`,
						IgnoreInTests: true,
					},
				}...),
		},
		{
			SubService: "snapshots",
			Struct:     &types.Snapshot{},
			SkipFields: []string{"AdministrativeActions", "ResourceARN", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
					{
						Name:          "administrative_actions",
						Type:          schema.TypeJSON,
						Resolver:      `schema.PathResolver("AdministrativeActions")`,
						IgnoreInTests: true,
					},
				}...),
		},
		{
			SubService: "storage_virtual_machines",
			Struct:     &types.StorageVirtualMachine{},
			SkipFields: []string{"ResourceARN", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService: "volumes",
			Struct:     &types.Volume{},
			SkipFields: []string{"AdministrativeActions", "ResourceARN", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
					{
						Name:          "administrative_actions",
						Type:          schema.TypeJSON,
						Resolver:      `schema.PathResolver("AdministrativeActions")`,
						IgnoreInTests: true,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "fsx"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("fsx")`
	}
	return resources
}
