package iam

import (
	"context"
	"net/url"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func policyVersions() *schema.Table {
	table_name := "aws_iam_policy_versions"
	return &schema.Table{
		Name:        table_name,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_PolicyVersion.html`,
		Resolver:    fetchPolicyVersion,
		Transform:   transformers.TransformWithStruct(&types.PolicyVersion{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "policy_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:     "document",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolvePolicyDocument,
			},
		},
	}
}

func fetchPolicyVersion(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	policy := parent.Item.(types.Policy)
	config := iam.GetPolicyVersionInput{
		PolicyArn: policy.Arn,
		VersionId: policy.DefaultVersionId,
	}
	policyVersionOutput, err := svc.GetPolicyVersion(ctx, &config)

	if err != nil {
		return err
	}

	res <- policyVersionOutput.PolicyVersion

	return nil
}

func resolvePolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.PolicyVersion)
	doc, err := url.QueryUnescape(aws.ToString(r.Document))
	if err != nil {
		return err
	}
	return resource.Set("document", doc)
}
