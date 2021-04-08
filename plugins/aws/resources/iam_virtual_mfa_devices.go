package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamVirtualMfaDevices() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_virtual_mfa_devices",
		Resolver:     fetchIamVirtualMfaDevices,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
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
				Name: "serial_number",
				Type: schema.TypeString,
			},
			{
				Name: "base32_string_seed",
				Type: schema.TypeByteArray,
			},
			{
				Name: "enable_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "qr_code_png",
				Type:     schema.TypeByteArray,
				Resolver: schema.PathResolver("QRCodePNG"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamVirtualMfaDeviceTags,
			},
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Arn"),
			},
			{
				Name:     "user_create_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("User.CreateDate"),
			},
			{
				Name:     "user_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Path"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.UserId"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.UserName"),
			},
			{
				Name:     "user_password_last_used",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("User.PasswordLastUsed"),
			},
			{
				Name:     "user_permissions_boundary_permissions_boundary_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.PermissionsBoundary.PermissionsBoundaryArn"),
			},
			{
				Name:     "user_permissions_boundary_permissions_boundary_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:     "user_tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamVirtualMfaDeviceUserTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamVirtualMfaDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config iam.ListVirtualMFADevicesInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListVirtualMFADevices(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.VirtualMFADevices
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}

	return nil
}
func resolveIamVirtualMfaDeviceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VirtualMFADevice)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
func resolveIamVirtualMfaDeviceUserTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VirtualMFADevice)
	tags := map[string]*string{}
	for _, t := range r.User.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("user_tags", tags)
	return nil
}
