package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotBillingGroups() *schema.Table {
	return &schema.Table{
		Name:          "aws_iot_billing_groups",
		Description:   "Billing groups are groups of things created for billing purposes that collect billable information for the things",
		Resolver:      fetchIotBillingGroups,
		Multiplex:     client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
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
				Name:        "things_in_group",
				Description: "Lists the things in the specified group",
				Type:        schema.TypeStringArray,
				Resolver:    ResolveIotBillingGroupThingsInGroup,
			},
			{
				Name:        "tags",
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotBillingGroupTags,
			},
			{
				Name:        "arn",
				Description: "The ARN of the billing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingGroupArn"),
			},
			{
				Name:        "id",
				Description: "The ID of the billing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingGroupId"),
			},
			{
				Name:        "creation_date",
				Description: "The date the billing group was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("BillingGroupMetadata.CreationDate"),
			},
			{
				Name:        "name",
				Description: "The name of the billing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingGroupName"),
			},
			{
				Name:        "description",
				Description: "The description of the billing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingGroupProperties.BillingGroupDescription"),
			},
			{
				Name:        "version",
				Description: "The version of the billing group.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotBillingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	input := iot.ListBillingGroupsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListBillingGroups(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, g := range response.BillingGroups {
			group, err := svc.DescribeBillingGroup(ctx, &iot.DescribeBillingGroupInput{
				BillingGroupName: g.GroupName,
			}, func(options *iot.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- group
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotBillingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeBillingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListThingsInBillingGroupInput{
		BillingGroupName: i.BillingGroupName,
		MaxResults:       aws.Int32(250),
	}

	var things []string
	for {
		response, err := svc.ListThingsInBillingGroup(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		things = append(things, response.Things...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, things)
}
func ResolveIotBillingGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeBillingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.BillingGroupArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})

		if err != nil {
			return diag.WrapError(err)
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
