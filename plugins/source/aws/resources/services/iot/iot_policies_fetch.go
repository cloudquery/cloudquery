package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIotPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListPoliciesInput{
		PageSize: aws.Int32(250),
	}

	for {
		response, err := svc.ListPolicies(ctx, &input)
		if err != nil {
			return err
		}

		for _, s := range response.Policies {
			profile, err := svc.GetPolicy(ctx, &iot.GetPolicyInput{
				PolicyName: s.PolicyName,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- profile
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func ResolveIotPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.GetPolicyOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.PolicyArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input)

		if err != nil {
			return err
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
