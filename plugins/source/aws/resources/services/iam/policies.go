package iam

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamPolicyTags,
			},
			{
				Name:     "policy_version_list",
				Type:     schema.TypeJSON,
				Resolver: resolveIamPolicyVersionList,
			},
		},
		Relations: []*schema.Table{policyLastAccessedDetails()},
	}
}

func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := iam.GetAccountAuthorizationDetailsInput{
		Filter: []types.EntityType{
			types.EntityTypeAWSManagedPolicy, types.EntityTypeLocalManagedPolicy,
		},
	}
	svc := meta.(*client.Client).Services().Iam
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

func resolveIamPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ManagedPolicyDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	response, err := svc.ListPolicyTags(ctx, &iam.ListPolicyTagsInput{PolicyArn: r.Arn})
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
