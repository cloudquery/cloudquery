package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource workspaces --config gen.hcl --output .
func Workspaces() *schema.Table {
	return &schema.Table{
		Name:         "aws_workspaces_workspaces",
		Description:  "Describes a WorkSpace.",
		Resolver:     fetchWorkspacesWorkspaces,
		Multiplex:    client.ServiceAccountRegionMultiplexer("workspaces"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Description: "The Amazon Resource Name (ARN) for the workspaces workspace",
				Type:        schema.TypeString,
			},
			{
				Name:        "bundle_id",
				Description: "The identifier of the bundle used to create the WorkSpace.",
				Type:        schema.TypeString,
			},
			{
				Name:        "computer_name",
				Description: "The name of the WorkSpace, as seen by the operating system",
				Type:        schema.TypeString,
			},
			{
				Name:        "directory_id",
				Description: "The identifier of the Directory Service directory for the WorkSpace.",
				Type:        schema.TypeString,
			},
			{
				Name:        "error_code",
				Description: "The error code that is returned if the WorkSpace cannot be created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "error_message",
				Description: "The text of the error message that is returned if the WorkSpace cannot be created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ip_address",
				Description: "The IP address of the WorkSpace.",
				Type:        schema.TypeString,
			},
			{
				Name:        "modification_states",
				Description: "The modification states of the WorkSpace.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "root_volume_encryption_enabled",
				Description: "Indicates whether the data stored on the root volume is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "state",
				Description: "The operational state of the WorkSpace",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "The identifier of the subnet for the WorkSpace.",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_name",
				Description: "The user for the WorkSpace.",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_volume_encryption_enabled",
				Description: "Indicates whether the data stored on the user volume is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "volume_encryption_key",
				Description: "The symmetric KMS key used to encrypt data stored on your WorkSpace",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier of the WorkSpace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceId"),
			},
			{
				Name:        "compute_type_name",
				Description: "The compute type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceProperties.ComputeTypeName"),
			},
			{
				Name:        "root_volume_size_gib",
				Description: "The size of the root volume",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("WorkspaceProperties.RootVolumeSizeGib"),
			},
			{
				Name:        "running_mode",
				Description: "The running mode",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceProperties.RunningMode"),
			},
			{
				Name:        "running_mode_auto_stop_timeout_in_minutes",
				Description: "The time after a user logs off when WorkSpaces are automatically stopped. Configured in 60-minute intervals.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("WorkspaceProperties.RunningModeAutoStopTimeoutInMinutes"),
			},
			{
				Name:        "user_volume_size_gib",
				Description: "The size of the user storage",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("WorkspaceProperties.UserVolumeSizeGib"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWorkspacesWorkspaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	// GENERATED from github.com/cloudquery/cq-gen/providers/aws.PaginatorTemplate. Do not edit.
	paginator := workspaces.NewDescribeWorkspacesPaginator(meta.(*client.Client).Services().Workspaces, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.Workspaces
	}
	return nil
}
