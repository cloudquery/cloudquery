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
)

func policyVersions() *schema.Table {
	table_name := "aws_iam_policy_versions"
	return &schema.Table{
		Name:        table_name,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_PolicyVersion.html`,
		Resolver:    fetchPolicyVersion,
		Transform:   transformers.TransformWithStruct(&iam.GetPolicyVersionOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "policy_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
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
	policyVersion, err := svc.GetPolicyVersion(ctx, &config)

	if err != nil {
		return err
	}

	doc, err := url.QueryUnescape(aws.ToString(policyVersion.PolicyVersion.Document))

	if err != nil {
		return err
	}

	policyVersion.PolicyVersion.Document = &doc

	res <- policyVersion

	return nil
}
