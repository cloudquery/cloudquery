package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IotSecurityProfiles() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_security_profiles",
		Resolver:  fetchIotSecurityProfiles,
		Multiplex: client.ServiceAccountRegionMultiplexer("iot"),
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
				Name:        "targets",
				Description: "Targets associated with the security profile",
				Type:        schema.TypeStringArray,
				Resolver:    ResolveIotSecurityProfileTargets,
			},
			{
				Name:        "tags",
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotSecurityProfileTags,
			},
			{
				Name:        "additional_metrics_to_retain",
				Description: "Please use DescribeSecurityProfileResponse$additionalMetricsToRetainV2 instead. A list of metrics whose data is retained (stored)",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "additional_metrics_to_retain_v2",
				Description: "A list of metrics whose data is retained (stored)",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("AdditionalMetricsToRetainV2"),
			},
			{
				Name:        "alert_targets",
				Description: "Where the alerts are sent",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "creation_date",
				Description: "The time the security profile was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_modified_date",
				Description: "The time the security profile was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:            "arn",
				Description:     "The ARN of the security profile.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("SecurityProfileArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "description",
				Description: "A description of the security profile (associated with the security profile when it was created or updated).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecurityProfileDescription"),
			},
			{
				Name:        "name",
				Description: "The name of the security profile.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecurityProfileName"),
			},
			{
				Name:        "version",
				Description: "The version of the security profile",
				Type:        schema.TypeInt,
			},
			{
				Name:        "behaviors",
				Type: 			schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotSecurityProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListSecurityProfilesInput{
		MaxResults: aws.Int32(250),
	}

	for {
		response, err := svc.ListSecurityProfiles(ctx, &input)
		if err != nil {
			return err
		}

		for _, s := range response.SecurityProfileIdentifiers {
			profile, err := svc.DescribeSecurityProfile(ctx, &iot.DescribeSecurityProfileInput{
				SecurityProfileName: s.Name,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- profile
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotSecurityProfileTargets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeSecurityProfileOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTargetsForSecurityProfileInput{
		SecurityProfileName: i.SecurityProfileName,
		MaxResults:          aws.Int32(250),
	}

	var targets []string
	for {
		response, err := svc.ListTargetsForSecurityProfile(ctx, &input)
		if err != nil {
			return err
		}

		for _, t := range response.SecurityProfileTargets {
			targets = append(targets, *t.Arn)
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, targets)
}
func ResolveIotSecurityProfileTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeSecurityProfileOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.SecurityProfileArn,
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
