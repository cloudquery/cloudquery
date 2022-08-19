package iam

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	smithy "github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/helpers"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cast"
)

//go:generate cq-gen --resource users --config gen.hcl --output .
func Users() *schema.Table {
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
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user was created.  This member is required.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("User.CreateDate"),
			},
			{
				Name:        "path",
				Description: "The path to the user",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.Path"),
			},
			{
				Name:        "id",
				Description: "The stable and unique string identifying the user",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.UserId"),
			},
			{
				Name:        "user_name",
				Description: "The friendly name identifying the user.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.UserName"),
			},
			{
				Name:        "password_last_used",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user's password was last used to sign in to an Amazon Web Services website",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("User.PasswordLastUsed"),
			},
			{
				Name:        "permissions_boundary_arn",
				Description: "The ARN of the policy used to set the permissions boundary for the user or role.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.PermissionsBoundary.PermissionsBoundaryArn"),
			},
			{
				Name:        "permissions_boundary_type",
				Description: "The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:        "tags",
				Description: "A list of tags that are associated with the user",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("User.Tags"),
			},
			{
				Name:     "user",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReportUser.User"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReportUser.ARN"),
			},
			{
				Name:     "password_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReportUser.PasswordStatus"),
			},
			{
				Name:     "password_last_changed",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReportUser.PasswordLastChanged"),
			},
			{
				Name:     "password_next_rotation",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReportUser.PasswordNextRotation"),
			},
			{
				Name:     "mfa_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReportUser.MfaActive"),
			},
			{
				Name:     "access_key1_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReportUser.AccessKey1Active"),
			},
			{
				Name:     "access_key2_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReportUser.AccessKey2Active"),
			},
			{
				Name:     "access_key1_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReportUser.AccessKey1LastRotated"),
			},
			{
				Name:     "access_key2_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReportUser.AccessKey2LastRotated"),
			},
			{
				Name:     "cert1_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReportUser.Cert1Active"),
			},
			{
				Name:     "cert2_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReportUser.Cert2Active"),
			},
			{
				Name:     "cert1_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReportUser.Cert1LastRotated"),
			},
			{
				Name:     "cert2_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReportUser.Cert2LastRotated"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_iam_user_access_keys",
				Resolver: fetchIamUserAccessKeys,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "access_key_id",
						Description: "The ID for this access key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessKeyMetadata.AccessKeyId"),
					},
					{
						Name:        "create_date",
						Description: "The date when the access key was created.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("AccessKeyMetadata.CreateDate"),
					},
					{
						Name:        "status",
						Description: "The status of the access key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessKeyMetadata.Status"),
					},
					{
						Name:        "user_name",
						Description: "The name of the IAM user that the key is associated with.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessKeyMetadata.UserName"),
					},
					{
						Name:        "last_used_date",
						Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the access key was most recently used. This field is null in the following situations:  * The user does not have an access key.  * An access key exists but has not been used since IAM began tracking this information.  * There is no sign-in data associated with the user.  This member is required.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("AccessKeyLastUsed.LastUsedDate"),
					},
					{
						Name:        "last_used_region",
						Description: "The Amazon Web Services Region where this access key was most recently used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessKeyLastUsed.Region"),
					},
					{
						Name:        "last_used_service_name",
						Description: "The name of the Amazon Web Services service with which this access key was most recently used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessKeyLastUsed.ServiceName"),
					},
				},
			},
			{
				Name:        "aws_iam_user_groups",
				Description: "Contains information about an IAM group entity",
				Resolver:    fetchIamUserGroups,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) specifying the group",
						Type:        schema.TypeString,
					},
					{
						Name:        "create_date",
						Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the group was created.  This member is required.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "group_id",
						Description: "The stable and unique string identifying the group",
						Type:        schema.TypeString,
					},
					{
						Name:        "group_name",
						Description: "The friendly name that identifies the group.  This member is required.",
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
				Name:        "aws_iam_user_attached_policies",
				Description: "Contains information about an attached policy",
				Resolver:    fetchIamUserAttachedPolicies,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "policy_arn",
						Description: "The Amazon Resource Name (ARN)",
						Type:        schema.TypeString,
					},
					{
						Name:        "policy_name",
						Description: "The friendly name of the attached policy.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_iam_user_policies",
				Description: "Contains the response to a successful GetUserPolicy request.",
				Resolver:    fetchIamUserPolicies,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "policy_document",
						Description: "The policy document",
						Type:        schema.TypeJSON,
						Resolver:    resolveUserPoliciesPolicyDocument,
					},
					{
						Name:        "policy_name",
						Description: "The name of the policy.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "user_name",
						Description: "The user the policy is associated with.  This member is required.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listUsers, userDetail))
}
func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(WrappedUser)
	cl := meta.(*client.Client)
	svc := cl.Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListAccessKeys(ctx, &config)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}

		for _, key := range output.AccessKeyMetadata {
			output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: key.AccessKeyId})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- wrappedKey{key, *output.AccessKeyLastUsed}
		}

		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsForUserInput
	p := parent.Item.(WrappedUser)
	if aws.ToString(p.UserName) == "<root_account>" {
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
func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAttachedUserPoliciesInput
	p := parent.Item.(WrappedUser)
	if aws.ToString(p.UserName) == "<root_account>" {
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
func fetchIamUserPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().IAM
	user := parent.Item.(WrappedUser)
	if aws.ToString(user.UserName) == "<root_account>" {
		return nil
	}
	config := iam.ListUserPoliciesInput{UserName: user.UserName}
	for {
		output, err := svc.ListUserPolicies(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		for _, p := range output.PolicyNames {
			policyCfg := &iam.GetUserPolicyInput{PolicyName: &p, UserName: user.UserName}
			policyResult, err := svc.GetUserPolicy(ctx, policyCfg)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- policyResult
		}
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
func resolveUserPoliciesPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetUserPolicyOutput)

	decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
	if err != nil {
		return diag.WrapError(err)
	}

	var document map[string]interface{}
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, document))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func listUsers(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	report, err := getCredentialReport(ctx, meta)
	if err != nil {
		return diag.WrapError(err)
	}

	for _, user := range report {
		detailChan <- user
	}
	return nil
}
func userDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	reportUser := listInfo.(*ReportUser)
	svc := meta.(*client.Client).Services().IAM
	userDetail, err := svc.GetUser(ctx, &iam.GetUserInput{
		UserName: aws.String(reportUser.User),
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- WrappedUser{*userDetail.User, reportUser}
}
func postIamUserResolver(_ context.Context, _ schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(WrappedUser)
	if r.ReportUser == nil {
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
		if err := resource.Set("cert1_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		cert1LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert1LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("cert1_last_rotated", cert1LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.Cert2LastRotated == "N/A" || r.Cert2LastRotated == "not_supported" {
		if err := resource.Set("cert2_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		cert2LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert2LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("cert2_last_rotated", cert2LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.AccessKey1LastRotated == "N/A" || r.AccessKey1LastRotated == "not_supported" {
		if err := resource.Set("access_key1_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		accessKey1LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey1LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("access_key1_last_rotated", accessKey1LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	if r.AccessKey2LastRotated == "N/A" || r.AccessKey2LastRotated == "not_supported" {
		if err := resource.Set("access_key2_last_rotated", nil); err != nil {
			return diag.WrapError(err)
		}
	} else {
		accessKey2LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey2LastRotated, location)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("access_key2_last_rotated", accessKey2LastRotated); err != nil {
			return diag.WrapError(err)
		}
	}

	return nil
}
func (r ReportUsers) GetUser(arn string) *ReportUser {
	for _, u := range r {
		if u.ARN == arn {
			return u
		}
	}
	return nil
}
func getCredentialReport(ctx context.Context, meta schema.ClientMeta) (ReportUsers, error) {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	svc := meta.(*client.Client).Services().IAM
	for {
		reportOutput, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err == nil && reportOutput != nil {
			var users ReportUsers
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
				var serviceError smithy.APIError
				if !errors.As(err, &serviceError) {
					return nil, diag.WrapError(err)
				}
				// LimitExceeded is the only specific error that should not stop processing
				// If Limit Exceeded is returned we should try and see if there is a credential report
				// already generated so we want to sleep for 5 seconds then continue
				if serviceError.ErrorCode() != "LimitExceeded" {
					return nil, diag.WrapError(err)
				}
				if err := helpers.Sleep(ctx, 5*time.Second); err != nil {
					return nil, diag.WrapError(err)
				}
			}
		case "ReportInProgress":
			meta.Logger().Debug("Waiting for credential report to be generated", "resource", "iam.users")
			if err := helpers.Sleep(ctx, 5*time.Second); err != nil {
				return nil, diag.WrapError(err)
			}
		default:
			return nil, diag.WrapError(err)
		}
	}
}
