package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IamVirtualMfaDevices() *schema.Table {
	return &schema.Table{
		Name:          "aws_iam_virtual_mfa_devices",
		Description:   "Contains information about a virtual MFA device.",
		Resolver:      fetchIamVirtualMfaDevices,
		Multiplex:     client.AccountMultiplex,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:            "serial_number",
				Description:     "The serial number associated with VirtualMFADevice.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Resolver:    client.ResolveTags,
			},
			{
				Name:     "user",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("User"),
			},
			{
				Name:        "user_password_last_used",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user's password was last used to sign in to an AWS website. For a list of AWS websites that capture a user's last sign-in time, see the Credential reports (https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html) topic in the IAM User Guide. If a password is used more than once in a five-minute span, only the first use is returned in this field. If the field is null (no value), then it indicates that they never signed in with a password. This can be because:",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("User.PasswordLastUsed"),
			},
			{
				Name:        "user_tags",
				Description: "A list of tags that are associated with the user. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. ",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchIamVirtualMfaDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
