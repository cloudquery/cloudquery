package iam

import (
	"context"
	"net/url"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Policies() *schema.Table {
	tableName := "aws_iam_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ManagedPolicyDetail.html`,
		Resolver:    fetchIamPolicies,
		Transform:   transformers.TransformWithStruct(&types.Policy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PolicyId"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIamPolicyTags,
			},
			{
				Name:     "policy_version",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: fetchPolicyVersion,
			},
		},
		Relations: []*schema.Table{
			policyLastAccessedDetails(),
		},
	}
}

func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListPoliciesPaginator(svc, nil)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Policies
	}
	return nil
}

func resolveIamPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Policy)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	response, err := svc.ListPolicyTags(ctx, &iam.ListPolicyTagsInput{PolicyArn: r.Arn}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set("tags", client.TagsToMap(response.Tags))
}

func fetchPolicyVersion(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	policy := resource.Item.(types.Policy)
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

	return resource.Set("policy_version", policyVersion)
}
