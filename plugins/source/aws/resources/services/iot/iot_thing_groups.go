package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IotThingGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_thing_groups",
		Description: "Groups allow you to manage several things at once by categorizing them into groups",
		Resolver:    fetchIotThingGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
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
				Resolver:    ResolveIotThingGroupThingsInGroup,
			},
			{
				Name:          "policies",
				Description:   "Policies associated with the thing group",
				Type:          schema.TypeStringArray,
				Resolver:      ResolveIotThingGroupPolicies,
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotThingGroupTags,
			},
			{
				Name:          "index_name",
				Description:   "The dynamic thing group index name.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "query_string",
				Description:   "The dynamic thing group search query string.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "query_version",
				Description:   "The dynamic thing group query version.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "status",
				Description: "The dynamic thing group status.",
				Type:        schema.TypeString,
			},
			{
				Name:            "arn",
				Description:     "The thing group ARN.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ThingGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "id",
				Description: "The thing group ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupId"),
			},
			{
				Name:     "thing_group_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ThingGroupMetadata"),
			},
			{
				Name:        "name",
				Description: "The name of the thing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupName"),
			},
			{
				Name:     "thing_group_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ThingGroupProperties"),
			},
			{
				Name:        "version",
				Description: "The version of the thing group.",
				Type:        schema.TypeInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotThingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	input := iot.ListThingGroupsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListThingGroups(ctx, &input)
		if err != nil {
			return err
		}
		for _, g := range response.ThingGroups {
			group, err := svc.DescribeThingGroup(ctx, &iot.DescribeThingGroupInput{
				ThingGroupName: g.GroupName,
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
func ResolveIotThingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeThingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListThingsInThingGroupInput{
		ThingGroupName: i.ThingGroupName,
		MaxResults:     aws.Int32(250),
	}

	var things []string
	for {
		response, err := svc.ListThingsInThingGroup(ctx, &input)
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
func ResolveIotThingGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeThingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListAttachedPoliciesInput{
		Target:   i.ThingGroupArn,
		PageSize: aws.Int32(250),
	}

	var policies []string
	for {
		response, err := svc.ListAttachedPolicies(ctx, &input)
		if err != nil {
			return err
		}

		for _, p := range response.Policies {
			policies = append(policies, *p.PolicyArn)
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return resource.Set(c.Name, policies)
}
func ResolveIotThingGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeThingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.ThingGroupArn,
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
