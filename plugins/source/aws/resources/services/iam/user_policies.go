package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamUserPolicies() *schema.Table {
	return &schema.Table{
		Name:          "aws_iam_user_policies",
		Description:   "Inline policies that are embedded in the specified IAM user",
		Resolver:      fetchIamUserPolicies,
		IgnoreError:   client.IgnoreCommonErrors,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "user_cq_id",
				Description: "Unique CloudQuery ID of aws_iam_users table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "user_id",
				Description: "user ID the policy belongs too.",
				Type:        schema.TypeString,
				Resolver:    schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:        "policy_document",
				Description: "The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamUserPolicyPolicyDocument,
			},
			{
				Name:        "policy_name",
				Description: "The name of the policy.",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_name",
				Description: "The user the policy is associated with.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamUserPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().IAM
	user := parent.Item.(wrappedUser)
	if aws.ToString(user.UserName) == rootName {
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
func resolveIamUserPolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
