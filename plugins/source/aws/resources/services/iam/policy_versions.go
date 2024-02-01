package iam

import (
	"context"
	"net/url"

	"github.com/apache/arrow/go/v15/arrow"
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
		Name:                table_name,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_PolicyVersion.html`,
		Resolver:            fetchPolicyVersion,
		PreResourceResolver: getPolicy,
		Transform:           transformers.TransformWithStruct(&types.PolicyVersion{}, transformers.WithPrimaryKeyComponents("VersionId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "policy_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "document_json",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolvePolicyDocument,
			},
		},
	}
}
func fetchPolicyVersion(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	policy := parent.Item.(types.Policy)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListPolicyVersionsPaginator(svc, &iam.ListPolicyVersionsInput{
		PolicyArn: policy.Arn,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Versions
	}

	return nil
}

func getPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	policy := resource.Parent.Item.(types.Policy)
	pv := resource.Item.(types.PolicyVersion)
	out, err := svc.GetPolicyVersion(ctx,
		&iam.GetPolicyVersionInput{
			PolicyArn: policy.Arn,
			VersionId: pv.VersionId},
		func(options *iam.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		cl.Logger().Warn().Err(err).Msg("Failed to get policy version")
		return err
	}

	resource.SetItem(*out.PolicyVersion)

	return nil
}

func resolvePolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicyVersion)
	doc, err := url.QueryUnescape(aws.ToString(r.Document))
	if err != nil {
		return err
	}
	return resource.Set(c.Name, doc)
}
