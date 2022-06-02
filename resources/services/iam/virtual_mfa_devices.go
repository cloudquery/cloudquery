package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamVirtualMfaDevices() *schema.Table {
	return &schema.Table{
		Name:          "aws_iam_virtual_mfa_devices",
		Description:   "Contains information about a virtual MFA device.",
		Resolver:      fetchIamVirtualMfaDevices,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"serial_number"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "serial_number",
				Description: "The serial number associated with VirtualMFADevice.",
				Type:        schema.TypeString,
			},
			{
				Name:        "base32_string_seed",
				Description: "The base32 seed defined as specified in RFC3548 (https://tools.ietf.org/html/rfc3548.txt). The Base32StringSeed is base64-encoded. ",
				Type:        schema.TypeByteArray,
			},
			{
				Name:        "enable_date",
				Description: "The date and time on which the virtual MFA device was enabled. ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "qr_code_png",
				Description: "A QR code PNG image that encodes otpauth://totp/$virtualMFADeviceName@$AccountName?secret=$Base32String where $virtualMFADeviceName is one of the create call arguments. AccountName is the user name if set (otherwise, the account ID otherwise), and Base32String is the seed in base32 format. The Base32String value is base64-encoded. ",
				Type:        schema.TypeByteArray,
				Resolver:    schema.PathResolver("QRCodePNG"),
			},
			{
				Name:        "tags",
				Description: "A list of tags that are attached to the virtual MFA device. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. ",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamVirtualMfaDeviceTags,
			},
			{
				Name:        "user_arn",
				Description: "The Amazon Resource Name (ARN) that identifies the user. For more information about ARNs and how to use ARNs in policies, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.Arn"),
			},
			{
				Name:        "user_create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("User.CreateDate"),
			},
			{
				Name:        "user_path",
				Description: "The path to the user. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. The ARN of the policy used to set the permissions boundary for the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.Path"),
			},
			{
				Name:        "user_id",
				Description: "The stable and unique string identifying the user. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.UserId"),
			},
			{
				Name:        "user_name",
				Description: "The friendly name identifying the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.UserName"),
			},
			{
				Name:        "user_password_last_used",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user's password was last used to sign in to an AWS website. For a list of AWS websites that capture a user's last sign-in time, see the Credential reports (https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html) topic in the IAM User Guide. If a password is used more than once in a five-minute span, only the first use is returned in this field. If the field is null (no value), then it indicates that they never signed in with a password. This can be because:",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("User.PasswordLastUsed"),
			},
			{
				Name:        "user_permissions_boundary_permissions_boundary_arn",
				Description: "The ARN of the policy used to set the permissions boundary for the user or role. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.PermissionsBoundary.PermissionsBoundaryArn"),
			},
			{
				Name:        "user_permissions_boundary_permissions_boundary_type",
				Description: "The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity. This data type can only have a value of Policy. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:        "user_tags",
				Description: "A list of tags that are associated with the user. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. ",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamVirtualMfaDeviceUserTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamVirtualMfaDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListVirtualMFADevicesInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListVirtualMFADevices(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, tags))
}
func resolveIamVirtualMfaDeviceUserTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VirtualMFADevice)
	if r.User == nil {
		return nil
	}
	tags := map[string]*string{}
	for _, t := range r.User.Tags {
		tags[*t.Key] = t.Value
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
