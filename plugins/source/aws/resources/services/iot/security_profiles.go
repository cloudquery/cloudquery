package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SecurityProfiles() *schema.Table {
	tableName := "aws_iot_security_profiles"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeSecurityProfile.html`,
		Resolver:    fetchIotSecurityProfiles,
		Transform:   transformers.TransformWithStruct(&iot.DescribeSecurityProfileOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "targets",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotSecurityProfileTargets,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotSecurityProfileTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityProfileArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIotSecurityProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
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
	svc := cl.Services().Iot
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
	svc := cl.Services().Iot
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
