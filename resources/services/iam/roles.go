package iam

import (
	"context"
	"errors"
	"net/url"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamRoles() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_roles",
		Description:  "An IAM role is an IAM identity that you can create in your account that has specific permissions.",
		Resolver:     fetchIamRoles,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "policies",
				Description: "List of policies attached to group.",
				Type:        schema.TypeJSON,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Resolver:    resolveIamRolePolicies,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the role. For more information about ARNs and how to use them in policies, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide guide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the role was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "path",
				Description: "The path to the role. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The stable and unique string identifying the role. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleId"),
			},
			{
				Name:        "name",
				Description: "The friendly name that identifies the role.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleName"),
			},
			{
				Name:        "assume_role_policy_document",
				Description: "The policy that grants an entity permission to assume the role. ",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamRoleAssumeRolePolicyDocument,
			},
			{
				Name:        "description",
				Description: "A description of the role that you provide. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_session_duration",
				Description: "The maximum session duration (in seconds) for the specified role. Anyone who uses the AWS CLI, or API to assume the role can specify the duration using the optional DurationSeconds API parameter or duration-seconds CLI parameter. ",
				Type:        schema.TypeInt,
			},
			{
				Name:          "permissions_boundary_arn",
				Description:   "The ARN of the policy used to set the permissions boundary for the user or role. ",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PermissionsBoundary.PermissionsBoundaryArn"),
				IgnoreInTests: true,
			},
			{
				Name:        "permissions_boundary_type",
				Description: "The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity. This data type can only have a value of Policy. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:          "role_last_used_last_used_date",
				Description:   "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601) that the role was last used. This field is null if the role has not been used within the IAM tracking period. For more information about the tracking period, see Regions where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period) in the IAM User Guide. ",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("RoleLastUsed.LastUsedDate"),
				IgnoreInTests: true,
			},
			{
				Name:          "role_last_used_region",
				Description:   "The name of the AWS Region in which the role was last used. ",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("RoleLastUsed.Region"),
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "A list of tags that are attached to the role. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. ",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamRoleTags,
			},
		},
		Relations: []*schema.Table{
			IamRolePolicies(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListRolesInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListRoles(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Roles
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveIamRolePolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Role)
	svc := meta.(*client.Client).Services().IAM
	input := iam.ListAttachedRolePoliciesInput{
		RoleName: r.RoleName,
	}
	policies := map[string]*string{}
	for {
		response, err := svc.ListAttachedRolePolicies(ctx, &input)
		if err != nil {
			var ae smithy.APIError
			if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchEntity" {
				return nil
			}
			return err
		}
		for _, p := range response.AttachedPolicies {
			policies[*p.PolicyArn] = p.PolicyName
		}
		if response.Marker == nil {
			break
		}
		input.Marker = response.Marker
	}
	return resource.Set("policies", policies)
}

func resolveIamRoleAssumeRolePolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Role)
	if r.AssumeRolePolicyDocument != nil {
		decodedDocument, err := url.QueryUnescape(*r.AssumeRolePolicyDocument)
		if err != nil {
			return err
		}
		return resource.Set("assume_role_policy_document", decodedDocument)
	}
	return nil
}
func resolveIamRoleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Role)
	svc := meta.(*client.Client).Services().IAM
	response, err := svc.ListRoleTags(ctx, &iam.ListRoleTagsInput{RoleName: r.RoleName})
	if err != nil {
		return err
	}
	tags := map[string]*string{}
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
