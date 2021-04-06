package resources

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func IamPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_policies",
		Resolver:     fetchIamPolicies,
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
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name: "attachment_count",
				Type: schema.TypeInt,
			},
			{
				Name: "create_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "default_version_id",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "is_attachable",
				Type: schema.TypeBool,
			},
			{
				Name: "path",
				Type: schema.TypeString,
			},
			{
				Name: "permissions_boundary_usage_count",
				Type: schema.TypeInt,
			},
			{
				Name: "policy_id",
				Type: schema.TypeString,
			},
			{
				Name: "policy_name",
				Type: schema.TypeString,
			},
			{
				Name: "update_date",
				Type: schema.TypeTimestamp,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_iam_policy_versions",
				Resolver: fetchIamPolicyVersions,
				Columns: []schema.Column{
					{
						Name:     "policy_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "create_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "document",
						Type:     schema.TypeJSON,
						Resolver: resolveIamPolicyVersionDocument,
					},
					{
						Name: "is_default_version",
						Type: schema.TypeBool,
					},
					{
						Name: "version_id",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func fetchIamPolicyVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
		resource.Set("document", decodedDocument)
	}
	return nil
}
