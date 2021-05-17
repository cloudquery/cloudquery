package resources

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
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamUsers() *schema.Table {
	return &schema.Table{
		Name:                 "aws_iam_users",
		Resolver:             fetchIamUsers,
		Multiplex:            client.AccountMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountFilter,
		PostResourceResolver: postIamUserResolver,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "password_last_used",
				Type: schema.TypeTimestamp,
			},
			// TODO: fix getcredentialreport
			//{
			//	Name:     "password_enabled",
			//	Type:     schema.TypeBool,
			//},
			//{
			//	Name:     "password_last_changed",
			//	Type:     schema.TypeTimestamp,
			//},
			//{
			//	Name:     "password_next_rotation",
			//	Type:     schema.TypeTimestamp,
			//},
			//{
			//	Name:     "mfa_active",
			//	Type:     schema.TypeBool,
			//},
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
				Resolver: resolveEc2UserTags,
			},
			{
				Name: "user_id",
				Type: schema.TypeString,
			},
			{
				Name: "user_name",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:                 "aws_iam_user_access_keys",
				Resolver:             fetchIamUserAccessKeys,
				PostResourceResolver: postIamUserAccessKeyResolver,
				Columns: []schema.Column{
					{
						Name:     "user_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
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
					// TODO: fix GetCredentialReport
					//{
					//	Name:     "last_rotated",
					//	Type:     schema.TypeTimestamp,
					//},
					{
						Name: "last_used_service_name",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_iam_user_groups",
				Resolver: fetchIamUserGroups,
				Columns: []schema.Column{
					{
						Name:     "user_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
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
				Columns: []schema.Column{
					{
						Name:     "user_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
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
		},
	}
}

type ReportUser struct {
	User                  string    `csv:"user"`
	ARN                   string    `csv:"arn"`
	UserCreationTime      time.Time `csv:"user_creation_time"`
	PasswordEnabled       string    `csv:"password_enabled"`
	PasswordLastUsed      string    `csv:"password_last_used"`
	PasswordLastChanged   string    `csv:"password_last_changed"`
	PasswordNextRotation  string    `csv:"password_next_rotation"`
	MFAActive             bool      `csv:"mfa_active"`
	AccessKey1Active      bool      `csv:"access_key_1_active"`
	AccessKey2Active      bool      `csv:"access_key_2_active"`
	AccessKey1LastRotated string    `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated string    `csv:"access_key_2_last_rotated"`
}

func GetCredentialReport(ctx context.Context, svc client.IamClient) ([]*ReportUser, error) {
	var err error
	var apiErr smithy.APIError
	//var reportOutput *iam.GetCredentialReportOutput
	for {
		_, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err != nil {
			if errors.As(err, &apiErr) {
				switch apiErr.ErrorCode() {
				case "ReportNotPresent", "ReportExpired":
					_, err := svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
					if err != nil {
						return nil, err
					}
				case "ReportInProgress":
					fmt.Println("Waiting for credential report to be generated", "resource", "iam.users")
					time.Sleep(2 * time.Second)
				default:
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			break
		}
	}
	var users []*ReportUser
	//err = gocsv.UnmarshalBytes(reportOutput.Content, &users)
	//if err != nil {
	//	return nil, err
	//}
	return users, nil
}

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config iam.ListUsersInput
	svc := meta.(*client.Client).Services().IAM
	_, err := GetCredentialReport(ctx, svc)
	if err != nil {
		return err
	}
	meta.(*client.Client).ReportUsers = nil

	for {
		output, err := svc.ListUsers(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Users
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

//func getMatchingReportUser(reportUsers []*ReportUser, user types.User) *ReportUser {
//	for _, reportUser := range reportUsers {
//		if *user.Arn == reportUser.ARN {
//			return reportUser
//		}
//	}
//	return nil
//}

func postIamUserResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	//r := resource.Item.(types.User)
	//reportUser := getMatchingReportUser(meta.(*provider.Client).ReportUsers.([]*ReportUser), r)
	//if reportUser != nil {
	//	location, err := time.LoadLocation("UTC")
	//	if err != nil {
	//		return err
	//	}
	//	passwordLastUsed, err := time.ParseInLocation(time.RFC3339, reportUser.PasswordLastUsed, location)
	//	if err == nil {
	//		resource.Set("password_last_used", passwordLastUsed)
	//	}
	//	passwordLastChanged, err := time.ParseInLocation(time.RFC3339, reportUser.PasswordLastChanged, location)
	//	if err == nil {
	//		resource.Set("password_last_changed", passwordLastChanged)
	//	}
	//}
	return nil
}

func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config iam.ListGroupsForUserInput
	p := parent.Item.(types.User)
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
	p := parent.Item.(types.User)
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListAccessKeys(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.AccessKeyMetadata
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func postIamUserAccessKeyResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(types.AccessKeyMetadata)
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
	p := parent.Item.(types.User)
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

func resolveEc2UserTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.User)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
