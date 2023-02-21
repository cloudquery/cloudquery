package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_users",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html`,
		Resolver:            fetchIamUsers,
		PreResourceResolver: getUser,
		Transform:           transformers.TransformWithStruct(&types.User{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:          "password_last_used",
				Description:   "The date and time when the AWS account root user or IAM user's password was last used to sign in to an AWS website",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "password_enabled",
				Description: "When the user has a password, this value is TRUE. Otherwise it is FALSE",
				Type:        schema.TypeBool,
			},
			{
				Name:        "password_status",
				Description: "When the user has a password, this value is TRUE. Otherwise it is FALSE.The value for the AWS account root user is always not_supported",
				Type:        schema.TypeString,
			},
			{
				Name:          "password_last_changed",
				Description:   "The date and time when the user's password was last set, in ISO 8601 date-time format. If the user does not have a password, the value in this field is N/A (not applicable). The value for the AWS account (root) is always NULL",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "password_next_rotation",
				Description:   "When the account has a password policy that requires password rotation, this field contains the date and time, in ISO 8601 date-time format, when the user is required to set a new password. The value for the AWS account (root) is always NULL",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "mfa_active",
				Description: "When a multi-factor authentication (MFA) device has been enabled for the user, this value is TRUE. Otherwise it is FALSE",
				Type:        schema.TypeBool,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (https://www.iso.org/iso/iso8601), when the user was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "path",
				Description: "The path to the user",
				Type:        schema.TypeString,
			},
			{
				Name:          "permissions_boundary_arn",
				Description:   "The ARN of the policy used to set the permissions boundary for the user or role",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PermissionsBoundary.PermissionsBoundaryArn"),
				IgnoreInTests: true,
			},
			{
				Name:        "permissions_boundary_type",
				Description: "The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:        "tags",
				Description: "A list of tags that are associated with the user",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "user_id",
				Description: "The stable and unique string identifying the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_name",
				Description: "The friendly name of the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "access_key_1_active",
				Description: "When the user has an access key and the access key's status is Active, this value is TRUE. Otherwise it is FALSE",
				Type:        schema.TypeBool,
			},
			client.DefaultAccountIDColumn(true),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			UserAccessKeys(),
			UserGroups(),
			UserAttachedPolicies(),
			UserPolicies(),
			SshPublicKeys(),
			SigningCertificates(),
		},
	}
}
