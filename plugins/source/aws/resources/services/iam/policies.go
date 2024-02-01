package iam

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
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
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("PolicyId"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIamPolicyTags,
			},
		},
		Relations: []*schema.Table{
			policyLastAccessedDetails(),
			policyVersions(),
		},
	}
}

func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListPoliciesPaginator(svc, &iam.ListPoliciesInput{
		MaxItems: aws.Int32(1000),
	})
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

	parsedArn, err := arn.Parse(aws.ToString(r.Arn))
	if err != nil {
		return err
	}
	// AWS Managed Policies do not include an accountID in the ARN and cannot have any tags so no need to call API
	if parsedArn.AccountID != cl.AccountID {
		return nil
	}
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
