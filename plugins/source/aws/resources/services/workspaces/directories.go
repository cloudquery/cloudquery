// Code generated by codegen; DO NOT EDIT.

package workspaces

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Directories() *schema.Table {
	return &schema.Table{
		Name:        "aws_workspaces_directories",
		Description: "https://docs.aws.amazon.com/workspaces/latest/api/API_WorkspaceDirectory.html",
		Resolver:    fetchWorkspacesDirectories,
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
				Resolver: resolveDirectoryArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "alias",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Alias"),
			},
			{
				Name:     "customer_user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerUserName"),
			},
			{
				Name:     "directory_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryId"),
			},
			{
				Name:     "directory_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryName"),
			},
			{
				Name:     "directory_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryType"),
			},
			{
				Name:     "dns_ip_addresses",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DnsIpAddresses"),
			},
			{
				Name:     "iam_role_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamRoleId"),
			},
			{
				Name:     "ip_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("IpGroupIds"),
			},
			{
				Name:     "registration_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistrationCode"),
			},
			{
				Name:     "saml_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamlProperties"),
			},
			{
				Name:     "selfservice_permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfservicePermissions"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubnetIds"),
			},
			{
				Name:     "tenancy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tenancy"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:     "workspace_creation_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceCreationProperties"),
			},
			{
				Name:     "workspace_security_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkspaceSecurityGroupId"),
			},
		},
	}
}
