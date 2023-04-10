package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ThingGroups() *schema.Table {
	tableName := "aws_iot_thing_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeThingGroup.html`,
		Resolver:    fetchIotThingGroups,
		Transform:   transformers.TransformWithStruct(&iot.DescribeThingGroupOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "things_in_group",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotThingGroupThingsInGroup,
			},
			{
				Name:     "policies",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotThingGroupPolicies,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotThingGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIotThingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListThingGroupsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().Iot
	paginator := iot.NewListThingGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, g := range page.ThingGroups {
			// TODO: Handle resolution in parallel with PreResourceResolver
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
	}
	return nil
}
func ResolveIotThingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeThingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListThingsInThingGroupInput{
		ThingGroupName: i.ThingGroupName,
		MaxResults:     aws.Int32(250),
	}

	var things []string
	paginator := iot.NewListThingsInThingGroupPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		things = append(things, page.Things...)
	}
	return resource.Set(c.Name, things)
}
func ResolveIotThingGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeThingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListAttachedPoliciesInput{
		Target:   i.ThingGroupArn,
		PageSize: aws.Int32(250),
	}

	var policies []string
	paginator := iot.NewListAttachedPoliciesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, p := range page.Policies {
			policies = append(policies, *p.PolicyArn)
		}
	}
	return resource.Set(c.Name, policies)
}
func ResolveIotThingGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeThingGroupOutput)
	svc := meta.(*client.Client).Services().Iot
	return resolveIotTags(ctx, svc, resource, c, i.ThingGroupArn)

}
