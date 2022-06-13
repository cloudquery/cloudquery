package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource directories --config gen.hcl --output .
func Directories() *schema.Table {
	return &schema.Table{
		Name:         "aws_workspaces_directories",
		Description:  "Describes a directory that is used with Amazon WorkSpaces.",
		Resolver:     fetchWorkspacesDirectories,
		Multiplex:    client.ServiceAccountRegionMultiplexer("workspaces"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the workspaces directory",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.WorkspacesService, func(resource *schema.Resource) ([]string, error) {
					return []string{"directory", *resource.Item.(types.WorkspaceDirectory).DirectoryId}, nil
				}),
			},
			{
				Name:        "alias",
				Description: "The directory alias.",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_user_name",
				Description: "The user name for the service account.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The directory identifier.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryId"),
			},
			{
				Name:        "name",
				Description: "The name of the directory.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryName"),
			},
			{
				Name:        "type",
				Description: "The directory type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryType"),
			},
			{
				Name:        "dns_ip_addresses",
				Description: "The IP addresses of the DNS servers for the directory.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "iam_role_id",
				Description: "The identifier of the IAM role",
				Type:        schema.TypeString,
			},
			{
				Name:          "ip_group_ids",
				Description:   "The identifiers of the IP access control groups associated with the directory.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "registration_code",
				Description: "The registration code for the directory",
				Type:        schema.TypeString,
			},
			{
				Name:        "change_compute_type",
				Description: "Specifies whether users can change the compute type (bundle) for their WorkSpace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfservicePermissions.ChangeComputeType"),
			},
			{
				Name:        "increase_volume_size",
				Description: "Specifies whether users can increase the volume size of the drives on their WorkSpace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfservicePermissions.IncreaseVolumeSize"),
			},
			{
				Name:        "rebuild_workspace",
				Description: "Specifies whether users can rebuild the operating system of a WorkSpace to its original state.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfservicePermissions.RebuildWorkspace"),
			},
			{
				Name:        "restart_workspace",
				Description: "Specifies whether users can restart their WorkSpace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfservicePermissions.RestartWorkspace"),
			},
			{
				Name:        "switch_running_mode",
				Description: "Specifies whether users can switch the running mode of their WorkSpace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfservicePermissions.SwitchRunningMode"),
			},
			{
				Name:        "state",
				Description: "The state of the directory's registration with Amazon WorkSpaces",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "The identifiers of the subnets used with the directory.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tenancy",
				Description: "Specifies whether the directory is dedicated or shared",
				Type:        schema.TypeString,
			},
			{
				Name:        "device_type_android",
				Description: "Indicates whether users can use Android and Android-compatible Chrome OS devices to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeAndroid"),
			},
			{
				Name:        "device_type_chrome_os",
				Description: "Indicates whether users can use Chromebooks to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeChromeOs"),
			},
			{
				Name:        "device_type_ios",
				Description: "Indicates whether users can use iOS devices to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeIos"),
			},
			{
				Name:        "device_type_linux",
				Description: "Indicates whether users can use Linux clients to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeLinux"),
			},
			{
				Name:        "device_type_osx",
				Description: "Indicates whether users can use macOS clients to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeOsx"),
			},
			{
				Name:        "device_type_web",
				Description: "Indicates whether users can access their WorkSpaces through a web browser.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeWeb"),
			},
			{
				Name:        "device_type_windows",
				Description: "Indicates whether users can use Windows clients to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeWindows"),
			},
			{
				Name:        "device_type_zero_client",
				Description: "Indicates whether users can use zero client devices to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeZeroClient"),
			},
			{
				Name:          "custom_security_group_id",
				Description:   "The identifier of the default security group to apply to WorkSpaces when they are created",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("WorkspaceCreationProperties.CustomSecurityGroupId"),
				IgnoreInTests: true,
			},
			{
				Name:          "default_ou",
				Description:   "The organizational unit (OU) in the directory for the WorkSpace machine accounts.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("WorkspaceCreationProperties.DefaultOu"),
				IgnoreInTests: true,
			},
			{
				Name:        "enable_internet_access",
				Description: "Specifies whether to automatically assign an Elastic public IP address to WorkSpaces in this directory by default",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("WorkspaceCreationProperties.EnableInternetAccess"),
			},
			{
				Name:        "enable_maintenance_mode",
				Description: "Specifies whether maintenance mode is enabled for WorkSpaces",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("WorkspaceCreationProperties.EnableMaintenanceMode"),
			},
			{
				Name:        "enable_work_docs",
				Description: "Specifies whether the directory is enabled for Amazon WorkDocs.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("WorkspaceCreationProperties.EnableWorkDocs"),
			},
			{
				Name:        "user_enabled_as_local_administrator",
				Description: "Specifies whether WorkSpace users are local administrators on their WorkSpaces.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("WorkspaceCreationProperties.UserEnabledAsLocalAdministrator"),
			},
			{
				Name:        "workspace_security_group_id",
				Description: "The identifier of the security group that is assigned to new WorkSpaces.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWorkspacesDirectories(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Workspaces
	input := workspaces.DescribeWorkspaceDirectoriesInput{}
	for {
		output, err := svc.DescribeWorkspaceDirectories(ctx, &input, func(o *workspaces.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Directories
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
