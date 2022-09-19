package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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

func resolveIamPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ManagedPolicyDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().IAM
	response, err := svc.ListPolicyTags(ctx, &iam.ListPolicyTagsInput{PolicyArn: r.Arn})
	if err != nil {
		if cl.IsNotFoundError(err) {
			meta.Logger().Debug().Err(err).Msg("ListPolicyTags: Policy does not exist")
			return nil
		}
		return err
	}
	return resource.Set("tags", client.TagsToMap(response.Tags))
}
