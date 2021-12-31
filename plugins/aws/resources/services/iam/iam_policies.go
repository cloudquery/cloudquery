package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_policies",
		Description:  "Contains information about a managed policy, including the policy's ARN, versions, and the number of principal entities (users, groups, and roles) that the policy is attached to.",
		Resolver:     fetchIamPolicies,
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN). ARNs are unique identifiers for AWS resources. For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the AWS General Reference. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "attachment_count",
				Description: "The number of principal entities (users, groups, and roles) that the policy is attached to. ",
				Type:        schema.TypeInt,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy was created. ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default_version_id",
				Description: "The identifier for the version of the policy that is set as the default (operative) version. For more information about policy versions, see Versioning for managed policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html) in the IAM User Guide. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A friendly description of the policy. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_attachable",
				Description: "Specifies whether the policy can be attached to an IAM user, group, or role. ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "path",
				Description: "The path to the policy. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "permissions_boundary_usage_count",
				Description: "The number of entities (users and roles) for which the policy is used as the permissions boundary. For more information about permissions boundaries, see Permissions boundaries for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the IAM User Guide. ",
				Type:        schema.TypeInt,
			},
			{
				Name:        "id",
				Description: "The stable and unique string identifying the policy. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyId"),
			},
			{
				Name:        "name",
				Description: "The friendly name (not ARN) identifying the policy. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyName"),
			},
			{
				Name:        "update_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy was last updated. When a policy has only one version, this field contains the date and time when the policy was created. When a policy has more than one version, this field contains the date and time when the most recent policy version was created. ",
				Type:        schema.TypeTimestamp,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iam_policy_versions",
				Description: "Contains information about a version of a managed policy.",
				Resolver:    fetchIamPolicyVersions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"policy_cq_id", "version_id"}},
				Columns: []schema.Column{
					{
						Name:        "policy_cq_id",
						Description: "Policy CloudQuery ID the policy versions belongs too.",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "policy_id",
						Description: "Policy ID the policy versions belongs too.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "create_date",
						Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy version was created. ",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "document",
						Description: "The policy document. The policy document is returned in the response to the GetPolicyVersion and GetAccountAuthorizationDetails operations. It is not returned in the response to the CreatePolicyVersion or ListPolicyVersions operations. The policy document returned in this structure is URL-encoded compliant with RFC 3986 (https://tools.ietf.org/html/rfc3986). You can use a URL decoding method to convert the policy back to plain JSON text. For example, if you use Java, you can use the decode method of the java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs provide similar functionality. ",
						Type:        schema.TypeJSON,
						Resolver:    resolveIamPolicyVersionDocument,
					},
					{
						Name:        "is_default_version",
						Description: "Specifies whether the policy version is set as the policy's default version. ",
						Type:        schema.TypeBool,
					},
					{
						Name:        "version_id",
						Description: "The identifier for the policy version. Policy version identifiers always begin with v (always lowercase). When a policy is created, the first policy version is v1. ",
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
func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.GetAccountAuthorizationDetailsInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.GetAccountAuthorizationDetails(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Policies
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func fetchIamPolicyVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.ManagedPolicyDetail)
	res <- r.PolicyVersionList
	return nil
}
func resolveIamPolicyVersionDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicyVersion)
	if r.Document != nil {
		decodedDocument, err := url.QueryUnescape(*r.Document)
		if err != nil {
			return err
		}
		data := make(map[string]interface{})
		if err := json.Unmarshal([]byte(decodedDocument), &data); err != nil {
			return err
		}
		return resource.Set("document", data)
	}
	return nil
}
