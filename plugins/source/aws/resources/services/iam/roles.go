package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource roles --config gen.hcl --output .
func Roles() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_roles",
		Description:  "An IAM role is an IAM identity that you can create in your account that has specific permissions.",
		Resolver:     fetchIamRoles,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreCommonErrors,
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
				Resolver:    resolveIamRolePolicies,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the role",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the role was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "path",
				Description: "The path to the role",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The stable and unique string identifying the role",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleId"),
			},
			{
				Name:        "role_name",
				Description: "The friendly name that identifies the role.",
				Type:        schema.TypeString,
			},
			{
				Name:        "assume_role_policy_document",
				Description: "The policy that grants an entity permission to assume the role.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRolesAssumeRolePolicyDocument,
			},
			{
				Name:        "description",
				Description: "A description of the role that you provide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_session_duration",
				Description: "The maximum session duration (in seconds) for the specified role",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "permissions_boundary_arn",
				Description: "The ARN of the policy used to set the permissions boundary for the user or role.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PermissionsBoundary.PermissionsBoundaryArn"),
			},
			{
				Name:        "permissions_boundary_type",
				Description: "The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:        "role_last_used_last_used_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601) that the role was last used",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("RoleLastUsed.LastUsedDate"),
			},
			{
				Name:        "role_last_used_region",
				Description: "The name of the Amazon Web Services Region in which the role was last used.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleLastUsed.Region"),
			},
			{
				Name:        "tags",
				Description: "A list of tags that are attached to the role",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iam_role_policies",
				Description: "Inline policies that are embedded in the specified IAM role",
				Resolver:    fetchIamRolePolicies,
				Columns: []schema.Column{
					{
						Name:        "role_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_roles table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "policy_document",
						Description: "The policy document",
						Type:        schema.TypeJSON,
						Resolver:    resolveRolePoliciesPolicyDocument,
					},
					{
						Name:        "policy_name",
						Description: "The name of the policy.",
						Type:        schema.TypeString,
					},
					{
						Name:        "role_name",
						Description: "The role the policy is associated with.",
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

func fetchIamRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listRoles, roleDetail))
}
func resolveIamRolePolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Role)
	cl := meta.(*client.Client)
	svc := cl.Services().IAM
	input := iam.ListAttachedRolePoliciesInput{
		RoleName: r.RoleName,
	}
	policies := map[string]*string{}
	for {
		response, err := svc.ListAttachedRolePolicies(ctx, &input)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		for _, p := range response.AttachedPolicies {
			policies[*p.PolicyArn] = p.PolicyName
		}
		if response.Marker == nil {
			break
		}
		input.Marker = response.Marker
	}
	return diag.WrapError(resource.Set("policies", policies))
}
func resolveRolesAssumeRolePolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Role)
	if r.AssumeRolePolicyDocument == nil {
		return nil
	}
	decodedDocument, err := url.QueryUnescape(*r.AssumeRolePolicyDocument)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set("assume_role_policy_document", decodedDocument))
}
func fetchIamRolePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().IAM
	role := parent.Item.(*types.Role)
	config := iam.ListRolePoliciesInput{
		RoleName: role.RoleName,
	}
	for {
		output, err := svc.ListRolePolicies(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		for _, p := range output.PolicyNames {
			policyResult, err := svc.GetRolePolicy(ctx, &iam.GetRolePolicyInput{PolicyName: &p, RoleName: role.RoleName})
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
func resolveRolePoliciesPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetRolePolicyOutput)

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

func listRoles(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	var config iam.ListRolesInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListRoles(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, role := range response.Roles {
			detailChan <- role
		}
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func roleDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	role := listInfo.(types.Role)
	svc := meta.(*client.Client).Services().IAM
	roleDetails, err := svc.GetRole(ctx, &iam.GetRoleInput{
		RoleName: role.RoleName,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- roleDetails.Role
}
