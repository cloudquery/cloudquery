package iam

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cast"
)

type wrappedKey struct {
	types.AccessKeyMetadata
	LastRotated time.Time
}

type wrappedUser struct {
	types.User
	*reportUser
	isRoot bool
}

type reportUser struct {
	User                  string    `csv:"user"`
	ARN                   string    `csv:"arn"`
	UserCreationTime      time.Time `csv:"user_creation_time"`
	PasswordStatus        string    `csv:"password_enabled"`
	PasswordLastChanged   string    `csv:"password_last_changed"`
	PasswordNextRotation  string    `csv:"password_next_rotation"`
	MfaActive             bool      `csv:"mfa_active"`
	AccessKey1Active      bool      `csv:"access_key_1_active"`
	AccessKey2Active      bool      `csv:"access_key_2_active"`
	AccessKey1LastRotated string    `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated string    `csv:"access_key_2_last_rotated"`
	Cert1Active           bool      `csv:"cert_1_active"`
	Cert2Active           bool      `csv:"cert_2_active"`
	Cert1LastRotated      string    `csv:"cert_1_last_rotated"`
	Cert2LastRotated      string    `csv:"cert_2_last_rotated"`
}

type reportUsers []*reportUser

const rootName = "<root_account>"

func IamUsers() *schema.Table {
	return &schema.Table{
		Name:                 "aws_iam_users",
		Resolver:             fetchIamUsers,
		Multiplex:            client.AccountMultiplex,
		IgnoreError:          client.IgnoreCommonErrors,
		DeleteFilter:         client.DeleteAccountFilter,
		PostResourceResolver: postIamUserResolver,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "id",
				Description: "The stable and unique string identifying the user",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserId"),
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
				Resolver:    resolveUserTags,
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
			{
				Name:        "access_key_1_last_rotated",
				Description: "The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "access_key_2_active",
				Description: "When the user has an access key and the access key's status is Active, this value is TRUE. Otherwise it is FALSE",
				Type:        schema.TypeBool,
			},
			{
				Name:          "access_key_2_last_rotated",
				Description:   "The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},

			{
				Name:        "cert_1_active",
				Description: "When the user has an X.509 signing certificate and that certificate's status is Active, this value is TRUE. Otherwise it is FALSE",
				Type:        schema.TypeBool,
			},
			{
				Name:          "cert_1_last_rotated",
				Description:   "The date and time, in ISO 8601 date-time format, when the user's signing certificate was created or last changed",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "cert_2_active",
				Description: "When the user has an X.509 signing certificate and that certificate's status is Active, this value is TRUE. Otherwise it is FALSE",
				Type:        schema.TypeBool,
			},
			{
				Name:          "cert_2_last_rotated",
				Description:   "The date and time, in ISO 8601 date-time format, when the user's signing certificate was created or last changed",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:                 "aws_iam_user_access_keys",
				Resolver:             fetchIamUserAccessKeys,
				PostResourceResolver: postIamUserAccessKeyResolver,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "user_id",
						Description: "The stable and unique string identifying the user",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("user_id"),
					},
					{
						Name:        "access_key_id",
						Description: "The ID for this access key",
						Type:        schema.TypeString,
					},
					{
						Name:        "create_date",
						Description: "The date when the access key was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "status",
						Description: "The status of the access key. Active means that the key is valid for API calls; Inactive means it is not",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_used",
						Description: "The date and time, in ISO 8601 date-time format, when the user's second access key was most recently used to sign an AWS API request",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "last_rotated",
						Description: "The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "last_used_service_name",
						Description: "The AWS service that was most recently accessed with the user's second access key",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_iam_user_groups",
				Resolver:      fetchIamUserGroups,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "user_id",
						Description: "The stable and unique string identifying the user",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("user_id"),
					},
					{
						Name:        "group_arn",
						Description: "The Amazon Resource Name (ARN) specifying the group",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Arn"),
					},
					{
						Name:        "create_date",
						Description: "The date and time, in ISO 8601 date-time format, when the group was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "group_id",
						Description: "The stable and unique string identifying the group",
						Type:        schema.TypeString,
					},
					{
						Name:        "group_name",
						Description: "The friendly name that identifies the group",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "The path to the group",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_iam_user_attached_policies",
				Resolver: fetchIamUserAttachedPolicies,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "user_id",
						Description: "The stable and unique string identifying the user",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("user_id"),
					},
					{
						Name:        "policy_arn",
						Description: "The Amazon Resource Name (ARN) of the policy",
						Type:        schema.TypeString,
					},
					{
						Name:        "policy_name",
						Description: "The friendly name of the attached policy",
						Type:        schema.TypeString,
					},
				},
			},
			IamUserPolicies(),
		},
	}
}

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	var config iam.ListUsersInput

	cl := meta.(*client.Client)
	svc := cl.Services().IAM
	report, err := getCredentialReport(ctx, meta)
	if err != nil {
		return diag.WrapError(err)
	}

	root := report.GetUser(fmt.Sprintf("arn:%s:iam::%s:root", cl.Partition, cl.AccountID))
	if root != nil {
		res <- wrappedUser{
			User: types.User{
				Arn:        aws.String(root.ARN),
				CreateDate: aws.Time(root.UserCreationTime),
				UserId:     aws.String("root"),
				UserName:   aws.String(root.User),
			},
			reportUser: root,
			isRoot:     true,
		}
	}

	for {
		output, err := svc.ListUsers(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}

		wUsers := make([]wrappedUser, len(output.Users))
		for i, u := range output.Users {
			ru := report.GetUser(aws.ToString(u.Arn))
			if ru == nil {
				meta.Logger().Warn("failed to find user in credential report", "arn", u.Arn)
				ru = &reportUser{}
			}
			wUsers[i] = wrappedUser{
				User:       u,
				reportUser: ru,
				isRoot:     false,
			}
		}

		res <- wUsers
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func postIamUserResolver(_ context.Context, _ schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(wrappedUser)
	if r.reportUser == nil {
		return nil
	}

	location, err := time.LoadLocation("UTC")
	if err != nil {
		return diag.WrapError(err)
	}

	// Only set if cast is successful
	if enabled, err := cast.ToBoolE(r.PasswordStatus); err == nil {
		if err := resource.Set("password_enabled", enabled); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.reportUser.ARN == "" {
		if err := resource.Set("password_next_rotation", nil); err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("password_last_changed", nil); err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("cert_1_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("cert_2_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("access_key_1_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}

		return diag.WrapError(resource.Set("access_key_2_last_rotated", nil))
	}

	if r.PasswordNextRotation == "N/A" || r.PasswordNextRotation == "not_supported" {
		if err := resource.Set("password_next_rotation", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		passwordNextRotation, err := time.ParseInLocation(time.RFC3339, r.PasswordNextRotation, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("password_next_rotation", passwordNextRotation); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.PasswordLastChanged == "N/A" || r.PasswordLastChanged == "not_supported" {
		if err := resource.Set("password_last_changed", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		passwordLastChanged, err := time.ParseInLocation(time.RFC3339, r.PasswordLastChanged, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("password_last_changed", passwordLastChanged); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.Cert1LastRotated == "N/A" || r.Cert1LastRotated == "not_supported" {
		if err := resource.Set("cert_1_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		cert1LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert1LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("cert_1_last_rotated", cert1LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.Cert2LastRotated == "N/A" || r.Cert2LastRotated == "not_supported" {
		if err := resource.Set("cert_2_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		cert2LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert2LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("cert_2_last_rotated", cert2LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.AccessKey1LastRotated == "N/A" || r.AccessKey1LastRotated == "not_supported" {
		if err := resource.Set("access_key_1_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		accessKey1LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey1LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("access_key_1_last_rotated", accessKey1LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.AccessKey2LastRotated == "N/A" || r.AccessKey2LastRotated == "not_supported" {
		if err := resource.Set("access_key_2_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		accessKey2LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey2LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("access_key_2_last_rotated", accessKey2LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	return nil
}

func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsForUserInput
	p := parent.Item.(wrappedUser)
	if aws.ToString(p.UserName) == rootName {
		return nil
	}
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListGroupsForUser(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Groups
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(wrappedUser)
	svc := meta.(*client.Client).Services().IAM
	if aws.ToString(p.UserName) == rootName {
		return nil
	}
	config.UserName = p.UserName
	for {
		output, err := svc.ListAccessKeys(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}

		keys := make([]wrappedKey, len(output.AccessKeyMetadata))
		for i, key := range output.AccessKeyMetadata {
			switch i {
			case 0:
				rotated := parent.Get("access_key_1_last_rotated")
				if rotated != nil {
					keys[i] = wrappedKey{key, rotated.(time.Time)}
				} else {
					keys[i] = wrappedKey{key, *key.CreateDate}
				}
			case 1:
				rotated := parent.Get("access_key_2_last_rotated")
				if rotated != nil {
					keys[i] = wrappedKey{key, rotated.(time.Time)}
				} else {
					keys[i] = wrappedKey{key, *key.CreateDate}
				}
			default:
				keys[i] = wrappedKey{key, time.Time{}}
			}
		}
		res <- keys
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func postIamUserAccessKeyResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(wrappedKey)
	if r.AccessKeyId == nil {
		return nil
	}
	svc := meta.(*client.Client).Services().IAM
	output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: r.AccessKeyId})
	if err != nil {
		return diag.WrapError(err)
	}
	if output.AccessKeyLastUsed != nil {
		if err := resource.Set("last_used", output.AccessKeyLastUsed.LastUsedDate); err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("last_used_service_name", output.AccessKeyLastUsed.ServiceName); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAttachedUserPoliciesInput
	p := parent.Item.(wrappedUser)
	if aws.ToString(p.UserName) == rootName {
		return nil
	}
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListAttachedUserPolicies(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.AttachedPolicies
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func resolveUserTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	svc := meta.(*client.Client).Services().IAM
	r := resource.Item.(wrappedUser)
	tags := map[string]*string{}

	if !r.isRoot {
		tagsOutput, err := svc.ListUserTags(ctx, &iam.ListUserTagsInput{UserName: r.UserName})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, t := range tagsOutput.Tags {
			tags[*t.Key] = t.Value
		}
	}
	return diag.WrapError(resource.Set("tags", tags))
}

func (r reportUsers) GetUser(arn string) *reportUser {
	for _, u := range r {
		if u.ARN == arn {
			return u
		}
	}
	return nil
}

func getCredentialReport(ctx context.Context, meta schema.ClientMeta) (reportUsers, error) {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	svc := meta.(*client.Client).Services().IAM
	for {
		reportOutput, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err == nil && reportOutput != nil {
			var users reportUsers
			err = gocsv.UnmarshalBytes(reportOutput.Content, &users)
			if err != nil {
				return nil, diag.WrapError(err)
			}
			return users, nil
		}
		if !errors.As(err, &apiErr) {
			return nil, diag.WrapError(err)
		}
		switch apiErr.ErrorCode() {
		case "ReportNotPresent", "ReportExpired":
			_, err := svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
			if err != nil {
				return nil, diag.WrapError(err)
			}
		case "ReportInProgress":
			meta.Logger().Debug("Waiting for credential report to be generated", "resource", "iam.users")
			time.Sleep(5 * time.Second)
		default:
			return nil, diag.WrapError(err)
		}
	}
}
