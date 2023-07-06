package iot

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SecurityProfiles() *schema.Table {
	tableName := "aws_iot_security_profiles"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeSecurityProfile.html`,
		Resolver:            fetchIotSecurityProfiles,
		PreResourceResolver: getSecurityProfile,
		Transform:           transformers.TransformWithStruct(&iot.DescribeSecurityProfileOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "targets",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: ResolveIotSecurityProfileTargets,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: ResolveIotSecurityProfileTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("SecurityProfileArn"),
				PrimaryKey: true,
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
	paginator := iot.NewListSecurityProfilesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SecurityProfileIdentifiers
	}
	return nil
}

func getSecurityProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot

	output, err := svc.DescribeSecurityProfile(ctx, &iot.DescribeSecurityProfileInput{
		SecurityProfileName: resource.Item.(types.SecurityProfileIdentifier).Name,
	}, func(options *iot.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output
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
	paginator := iot.NewListTargetsForSecurityProfilePaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		for _, t := range page.SecurityProfileTargets {
			targets = append(targets, *t.Arn)
		}
	}
	return resource.Set(c.Name, targets)
}
func ResolveIotSecurityProfileTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeSecurityProfileOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	return resolveIotTags(ctx, meta, svc, resource, c, i.SecurityProfileArn)
}
