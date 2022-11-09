package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FSXResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "backups",
			Struct:      &types.Backup{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_Backup.html",
			SkipFields:  []string{"BackupId"},
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
			},
		},
		{
			SubService:  "data_repository_associations",
			Struct:      &types.DataRepositoryAssociation{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_DataRepositoryAssociation.html",
			SkipFields:  []string{"ResourceARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "data_repository_tasks",
			Struct:      &types.DataRepositoryTask{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_DataRepositoryTask.html",
			SkipFields:  []string{"ResourceARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "file_caches",
			Struct:      &types.FileCache{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileCache.html",
			SkipFields:  []string{"ResourceARN"},
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
						Resolver: `resolveFileCacheTags`,
					},
				}...),
		},
		{
			SubService:  "file_systems",
			Struct:      &types.FileSystem{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileSystem.html",
			SkipFields:  []string{"AdministrativeActions", "ResourceARN", "Tags"},
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
			SubService:  "snapshots",
			Struct:      &types.Snapshot{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_Snapshot.html",
			SkipFields:  []string{"AdministrativeActions", "ResourceARN"},
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
						Name:          "administrative_actions",
						Type:          schema.TypeJSON,
						Resolver:      `schema.PathResolver("AdministrativeActions")`,
						IgnoreInTests: true,
					},
				}...),
		},
		{
			SubService:  "storage_virtual_machines",
			Struct:      &types.StorageVirtualMachine{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachine.html",
			SkipFields:  []string{"ResourceARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ResourceARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "volumes",
			Struct:      &types.Volume{},
			Description: "https://docs.aws.amazon.com/fsx/latest/APIReference/API_Volume.html",
			SkipFields:  []string{"AdministrativeActions", "ResourceARN"},
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
