package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamRolePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_role_policies",
		Description: "Inline policies that are embedded in the specified IAM role",
		Resolver:    fetchIamRolePolicies,
		Multiplex:   client.AccountMultiplex,
		IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"role_cq_id", "policy_name"}},
		Columns: []schema.Column{
			{
				Name:        "role_cq_id",
				Description: "Role CloudQuery ID the policy belongs too.",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "role_id",
				Description: "Role ID the policy belongs too.",
				Type:        schema.TypeString,
				Resolver:    schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "policy_document",
				Description: "The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamRolePolicyPolicyDocument,
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
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamRolePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var ae smithy.APIError
	svc := meta.(*client.Client).Services().IAM
	role := parent.Item.(types.Role)
	config := iam.ListRolePoliciesInput{
		RoleName: role.RoleName,
	}
	for {
		output, err := svc.ListRolePolicies(ctx, &config)
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchEntity" {
				return nil
			}
			return err
		}
		for _, p := range output.PolicyNames {
			policyResult, err := svc.GetRolePolicy(ctx, &iam.GetRolePolicyInput{PolicyName: &p, RoleName: role.RoleName})
			if err != nil {
				return err
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
func resolveIamRolePolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*iam.GetRolePolicyOutput)
	if !ok {
		return fmt.Errorf("not role policy")
	}

	decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
	if err != nil {
		return err
	}

	var document map[string]interface{}
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, document)
}
