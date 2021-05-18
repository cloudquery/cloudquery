package resources

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamRoles() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_roles",
		Resolver:     fetchIamRoles,
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
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: resolveIamRolePolicies,
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
				Name: "path",
				Type: schema.TypeString,
			},
			{
				Name: "role_id",
				Type: schema.TypeString,
			},
			{
				Name: "role_name",
				Type: schema.TypeString,
			},
			{
				Name:     "assume_role_policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveIamRoleAssumeRolePolicyDocument,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "max_session_duration",
				Type: schema.TypeInt,
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
				Name:     "role_last_used_last_used_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RoleLastUsed.LastUsedDate"),
			},
			{
				Name:     "role_last_used_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleLastUsed.Region"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamRoleTags,
			},
		},
		Relations: []*schema.Table{
			iamRolePolicies(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
