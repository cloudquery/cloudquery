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

func IamGroupPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_group_policies",
		Description:  "Inline policies that are embedded in the specified IAM group",
		Resolver:     fetchIamGroupPolicies,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"group_cq_id", "policy_name"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "group_cq_id",
				Description: "Unique CloudQuery ID of aws_iam_groups table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "group_id",
				Description: "Group ID the policy belongs too.",
				Type:        schema.TypeString,
				Resolver:    schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:        "group_name",
				Description: "The group the policy is associated with.",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy_document",
				Description: "The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamGroupPolicyPolicyDocument,
			},
			{
				Name:        "policy_name",
				Description: "The name of the policy.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var ae smithy.APIError
	svc := meta.(*client.Client).Services().IAM
	group := parent.Item.(types.Group)
	config := iam.ListGroupPoliciesInput{
		GroupName: group.GroupName,
	}
	for {
		output, err := svc.ListGroupPolicies(ctx, &config)
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchEntity" {
				return nil
			}
			return err
		}
		for _, p := range output.PolicyNames {
			policyResult, err := svc.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{PolicyName: &p, GroupName: group.GroupName})
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
func resolveIamGroupPolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*iam.GetGroupPolicyOutput)
	if !ok {
		return fmt.Errorf("not group policy")
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
