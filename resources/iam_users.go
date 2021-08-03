package resources

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cast"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const rootName = "<root_account>"

func IamUsers() *schema.Table {
	return &schema.Table{
		Name:                 "aws_iam_users",
		Resolver:             fetchIamUsers,
		Multiplex:            client.AccountMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountFilter,
		PostResourceResolver: postIamUserResolver,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserId"),
			},
			{
				Name: "password_last_used",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name: "password_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "password_status",
				Type: schema.TypeString,
			},
			{
				Name: "password_last_changed",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "password_next_rotation",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "mfa_active",
				Type: schema.TypeBool,
			},
			{
				Name: "create_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "path",
				Type: schema.TypeString,
			},
			{
				Name:     "permissions_boundary_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PermissionsBoundary.PermissionsBoundaryArn"),
			},
			{
				Name:     "permissions_boundary_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveUserTags,
			},
			{
				Name: "user_id",
				Type: schema.TypeString,
			},
			{
				Name: "user_name",
				Type: schema.TypeString,
			},
			{
				Name: "access_key_1_active",
				Type: schema.TypeBool,
			},
			{
				Name: "access_key_1_last_rotated",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "access_key_2_active",
				Type: schema.TypeBool,
			},
			{
				Name: "access_key_2_last_rotated",
				Type: schema.TypeTimestamp,
			},

			{
				Name: "cert_1_active",
				Type: schema.TypeBool,
			},
			{
				Name: "cert_1_last_rotated",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "cert_2_active",
				Type: schema.TypeBool,
			},
			{
				Name: "cert_2_last_rotated",
				Type: schema.TypeTimestamp,
			},
		},
		Relations: []*schema.Table{
			{
				Name:                 "aws_iam_user_access_keys",
				Resolver:             fetchIamUserAccessKeys,
				PostResourceResolver: postIamUserAccessKeyResolver,
				Options:              schema.TableCreationOptions{PrimaryKeys: []string{"user_cq_id", "access_key_id"}},
				Columns: []schema.Column{
					{
						Name:     "user_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("user_id"),
					},
					{
						Name: "access_key_id",
						Type: schema.TypeString,
					},
					{
						Name: "create_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "last_used",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_rotated",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_used_service_name",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_iam_user_groups",
				Resolver: fetchIamUserGroups,
				Options:  schema.TableCreationOptions{PrimaryKeys: []string{"user_cq_id", "group_id"}},
				Columns: []schema.Column{
					{
						Name:     "user_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("user_id"),
					},
					{
						Name: "arn",
						Type: schema.TypeString,
					},
					{
						Name: "create_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "group_id",
						Type: schema.TypeString,
					},
					{
						Name: "group_name",
						Type: schema.TypeString,
					},
					{
						Name: "path",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_iam_user_attached_policies",
				Resolver: fetchIamUserAttachedPolicies,
				Options:  schema.TableCreationOptions{PrimaryKeys: []string{"user_cq_id", "policy_name"}},
				Columns: []schema.Column{
					{
						Name:     "user_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("user_id"),
					},
					{
						Name: "policy_arn",
						Type: schema.TypeString,
					},
					{
						Name: "policy_name",
						Type: schema.TypeString,
					},
				},
			},
			IamUserPolicies(),
		},
	}
}

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config iam.ListUsersInput
	svc := meta.(*client.Client).Services().IAM
	report, err := getCredentialReport(ctx, meta)
	if err != nil {
		return err
	}
	meta.(*client.Client).ReportUsers = nil

	root := report.GetUser(fmt.Sprintf("arn:aws:iam::%s:root", meta.(*client.Client).AccountID))
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
			return err
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
		return err
	}

	// Only set if cast is successful
	if enabled, err := cast.ToBoolE(r.PasswordStatus); err == nil {
		if err := resource.Set("password_enabled", enabled); err != nil {
			return err
		}
	}

	if r.PasswordNextRotation == "N/A" || r.PasswordNextRotation == "not_supported" {
		if err := resource.Set("password_next_rotation", nil); err != nil {
			return err
		}
	} else {
		passwordNextRotation, err := time.ParseInLocation(time.RFC3339, r.PasswordNextRotation, location)
		if err != nil {
			return err
		}
		if err := resource.Set("password_next_rotation", passwordNextRotation); err != nil {
			return err
		}
	}

	if r.PasswordLastChanged == "N/A" || r.PasswordLastChanged == "not_supported" {
		if err := resource.Set("password_last_changed", nil); err != nil {
			return err
		}
	} else {
		passwordLastChanged, err := time.ParseInLocation(time.RFC3339, r.PasswordLastChanged, location)
		if err != nil {
			return err
		}
		if err := resource.Set("password_last_changed", passwordLastChanged); err != nil {
			return err
		}
	}

	if r.Cert1LastRotated == "N/A" || r.Cert1LastRotated == "not_supported" {
		if err := resource.Set("cert_1_last_rotated", nil); err != nil {
			return err
		}
	} else {
		cert1LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert1LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("cert_1_last_rotated", cert1LastRotated); err != nil {
			return err
		}
	}

	if r.Cert2LastRotated == "N/A" || r.Cert2LastRotated == "not_supported" {
		if err := resource.Set("cert_2_last_rotated", nil); err != nil {
			return err
		}
	} else {
		cert2LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert2LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("cert_2_last_rotated", cert2LastRotated); err != nil {
			return err
		}
	}

	if r.AccessKey1LastRotated == "N/A" || r.AccessKey1LastRotated == "not_supported" {
		if err := resource.Set("access_key_1_last_rotated", nil); err != nil {
			return err
		}
	} else {
		accessKey1LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey1LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("access_key_1_last_rotated", accessKey1LastRotated); err != nil {
			return err
		}
	}

	if r.AccessKey2LastRotated == "N/A" || r.AccessKey2LastRotated == "not_supported" {
		if err := resource.Set("access_key_2_last_rotated", nil); err != nil {
			return err
		}
	} else {
		accessKey2LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey2LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("access_key_2_last_rotated", accessKey2LastRotated); err != nil {
			return err
		}
	}

	return nil
}

func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
			return err
		}
		res <- output.Groups
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
			return err
		}

		keys := make([]wrappedKey, len(output.AccessKeyMetadata))
		for i, key := range output.AccessKeyMetadata {
			switch i {
			case 0:
				rotated := parent.Get("access_key_1_last_rotated")
				if rotated != nil {
					keys[i] = wrappedKey{key, rotated.(time.Time)}
				}
			case 1:
				rotated := parent.Get("access_key_2_last_rotated")
				if rotated != nil {
					keys[i] = wrappedKey{key, rotated.(time.Time)}
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
	svc := meta.(*client.Client).Services().IAM
	output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: r.AccessKeyId})
	if err != nil {
		return err
	}
	if output.AccessKeyLastUsed != nil {
		if err := resource.Set("last_used", output.AccessKeyLastUsed.LastUsedDate); err != nil {
			return err
		}
		if err := resource.Set("last_used_service_name", output.AccessKeyLastUsed.ServiceName); err != nil {
			return err
		}
	}
	return nil
}

func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
			return err
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
			return err
		}
		for _, t := range tagsOutput.Tags {
			tags[*t.Key] = t.Value
		}
	}
	return resource.Set("tags", tags)
}

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
				return nil, err
			}
			return users, nil
		}
		if !errors.As(err, &apiErr) {
			return nil, err
		}
		switch apiErr.ErrorCode() {
		case "ReportNotPresent", "ReportExpired":
			_, err := svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
			if err != nil {
				return nil, err
			}
		case "ReportInProgress":
			meta.Logger().Debug("Waiting for credential report to be generated", "resource", "iam.users")
			time.Sleep(5 * time.Second)
		default:
			return nil, err
		}
	}
}
