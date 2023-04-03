package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BillingGroups() *schema.Table {
	tableName := "aws_iot_billing_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeBillingGroup.html`,
		Resolver:    fetchIotBillingGroups,
		Transform:   transformers.TransformWithStruct(&iot.DescribeBillingGroupOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "things_in_group",
				Type:     schema.TypeStringArray,
				Resolver: resolveIotBillingGroupThingsInGroup,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIotBillingGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BillingGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIotBillingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListBillingGroupsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().Iot
	for {
		response, err := svc.ListBillingGroups(ctx, &input)
		if err != nil {
			return err
		}
		for _, g := range response.BillingGroups {
			group, err := svc.DescribeBillingGroup(ctx, &iot.DescribeBillingGroupInput{
				BillingGroupName: g.GroupName,
			}, func(options *iot.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
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
func resolveIotBillingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeBillingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListThingsInBillingGroupInput{
		BillingGroupName: i.BillingGroupName,
		MaxResults:       aws.Int32(250),
	}

	var things []string
	for {
		response, err := svc.ListThingsInBillingGroup(ctx, &input)
		if err != nil {
			return err
		}

		things = append(things, response.Things...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, things)
}
func resolveIotBillingGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeBillingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.BillingGroupArn,
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
