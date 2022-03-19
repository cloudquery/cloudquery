package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_policies",
		Description:  "The output from the GetPolicy operation.",
		Resolver:     fetchIotPolicies,
		Multiplex:    client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotPolicyTags,
			},
			{
				Name:        "creation_date",
				Description: "The date the policy was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default_version_id",
				Description: "The default policy version ID.",
				Type:        schema.TypeString,
			},
			{
				Name:        "generation_id",
				Description: "The generation ID of the policy.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_date",
				Description: "The date the policy was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "The policy ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyArn"),
			},
			{
				Name:        "document",
				Description: "The JSON document that describes the policy.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyDocument"),
			},
			{
				Name:        "name",
				Description: "The policy name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyName"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListPoliciesInput{
		PageSize: aws.Int32(250),
	}

	for {
		response, err := svc.ListPolicies(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		for _, s := range response.Policies {
			profile, err := svc.GetPolicy(ctx, &iot.GetPolicyInput{
				PolicyName: s.PolicyName,
			}, func(options *iot.Options) {
				options.Region = client.Region
			})
			if err != nil {
				return diag.WrapError(err)
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
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.PolicyArn,
	}
	tags := make(map[string]interface{})

	for {
		response, err := svc.ListTagsForResource(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})

		if err != nil {
			return diag.WrapError(err)
		}
		for _, t := range response.Tags {
			tags[*t.Key] = t.Value
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
