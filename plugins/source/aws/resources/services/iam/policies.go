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
		Transform:   transformers.TransformWithStruct(&types.ManagedPolicyDetail{}),
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
				Name:     "policy_version_list",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIamPolicyVersionList,
			},
		},
		Relations: []*schema.Table{
			policyLastAccessedDetails(),
		},
	}
}

func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := iam.GetAccountAuthorizationDetailsInput{
		Filter: []types.EntityType{
			types.EntityTypeAWSManagedPolicy, types.EntityTypeLocalManagedPolicy,
		},
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	paginator := iam.NewGetAccountAuthorizationDetailsPaginator(svc, &config)

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
	r := resource.Item.(types.ManagedPolicyDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
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

func resolveIamPolicyVersionList(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ManagedPolicyDetail)
	for i := range r.PolicyVersionList {
		if v, err := url.QueryUnescape(aws.ToString(r.PolicyVersionList[i].Document)); err == nil {
			r.PolicyVersionList[i].Document = &v
		}
	}
	return resource.Set(c.Name, r.PolicyVersionList)
}
