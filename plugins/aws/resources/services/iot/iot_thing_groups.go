package iot

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotThingGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_thing_groups",
		Description:  "Groups allow you to manage several things at once by categorizing them into groups",
		Resolver:     fetchIotThingGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreError:  client.IgnoreCommonErrors,
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
				Name:        "arn",
				Description: "The thing group ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupArn"),
			},
			{
				Name:        "id",
				Description: "The thing group ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupId"),
			},
			{
				Name:        "creation_date",
				Description: "The UNIX timestamp of when the thing group was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ThingGroupMetadata.CreationDate"),
			},
			{
				Name:        "parent_group_name",
				Description: "The parent thing group name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupMetadata.ParentGroupName"),
			},
			{
				Name:        "root_to_parent_thing_groups",
				Description: "The root parent thing group.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIotThingGroupsRootToParentThingGroups,
			},
			{
				Name:        "name",
				Description: "The name of the thing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupName"),
			},
			{
				Name:        "attribute_payload_attributes",
				Description: "A JSON string containing up to three key-value pair in JSON format",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ThingGroupProperties.AttributePayload.Attributes"),
			},
			{
				Name:        "attribute_payload_merge",
				Description: "Specifies whether the list of attributes provided in the AttributePayload is merged with the attributes stored in the registry, instead of overwriting them. To remove an attribute, call UpdateThing with an empty attribute value",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ThingGroupProperties.AttributePayload.Merge"),
			},
			{
				Name:        "thing_group_description",
				Description: "The thing group description.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupProperties.ThingGroupDescription"),
			},
			{
				Name:        "version",
				Description: "The version of the thing group.",
				Type:        schema.TypeBigInt,
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
		response, err := svc.ListThingGroups(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, g := range response.ThingGroups {
			group, err := svc.DescribeThingGroup(ctx, &iot.DescribeThingGroupInput{
				ThingGroupName: g.GroupName,
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
		response, err := svc.ListThingsInThingGroup(ctx, &input, func(options *iot.Options) {
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
		response, err := svc.ListAttachedPolicies(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
func resolveIotThingGroupsRootToParentThingGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeThingGroupOutput)
	if i.ThingGroupMetadata == nil {
		return nil
	}

	data, err := json.Marshal(i.ThingGroupMetadata.RootToParentThingGroups)
	if err != nil {
		return diag.WrapError(err)
	}

	return resource.Set(c.Name, data)
}
