// Code generated by codegen; DO NOT EDIT.

package workspaces

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Workspaces() *schema.Table {
	return &schema.Table{
		Name:        "aws_workspaces_workspaces",
		Description: "https://docs.aws.amazon.com/workspaces/latest/api/API_Workspace.html",
		Resolver:    fetchWorkspacesWorkspaces,
		Multiplex:   client.ServiceAccountRegionMultiplexer("workspaces"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWorkspaceArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "bundle_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BundleId"),
			},
			{
				Name:     "computer_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ComputerName"),
			},
			{
				Name:     "directory_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryId"),
			},
			{
				Name:     "error_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorCode"),
			},
			{
				Name:     "error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorMessage"),
			},
			{
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddress"),
			},
			{
				Name:     "modification_states",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ModificationStates"),
			},
			{
				Name:     "root_volume_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RootVolumeEncryptionEnabled"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetId"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserName"),
			},
			{
				Name:     "user_volume_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("UserVolumeEncryptionEnabled"),
			},
			{
				Name:     "volume_encryption_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VolumeEncryptionKey"),
			},
			{
				Name:     "workspace_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkspaceId"),
			},
			{
				Name:     "workspace_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceProperties"),
			},
		},
	}
}
