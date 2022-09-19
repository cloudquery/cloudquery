// Code generated by codegen; DO NOT EDIT.

package fsx

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StorageVirtualMachines() *schema.Table {
	return &schema.Table{
		Name:      "aws_fsx_storage_virtual_machines",
		Resolver:  fetchFsxStorageVirtualMachines,
		Multiplex: client.ServiceAccountRegionMultiplexer("fsx"),
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
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "active_directory_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveDirectoryConfiguration"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Endpoints"),
			},
			{
				Name:     "file_system_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemId"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "lifecycle_transition_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LifecycleTransitionReason"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "root_volume_security_style",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RootVolumeSecurityStyle"),
			},
			{
				Name:     "storage_virtual_machine_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageVirtualMachineId"),
			},
			{
				Name:     "subtype",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subtype"),
			},
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UUID"),
			},
		},
	}
}
