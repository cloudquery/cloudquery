package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BillingGroups() *schema.Table {
	tableName := "aws_iot_billing_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeBillingGroup.html`,
		Resolver:            fetchIotBillingGroups,
		PreResourceResolver: getBillingGroup,
		Transform:           transformers.TransformWithStruct(&iot.DescribeBillingGroupOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
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
	paginator := iot.NewListBillingGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.BillingGroups
	}
	return nil
}

func getBillingGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot

	output, err := svc.DescribeBillingGroup(ctx, &iot.DescribeBillingGroupInput{
		BillingGroupName: resource.Item.(types.GroupNameAndArn).GroupName,
	})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

// TODO: Move this to a new table
func resolveIotBillingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeBillingGroupOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListThingsInBillingGroupInput{
		BillingGroupName: i.BillingGroupName,
		MaxResults:       aws.Int32(250),
	}
	var things []string
	paginator := iot.NewListThingsInBillingGroupPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		things = append(things, page.Things...)
	}
	return resource.Set(c.Name, things)
}
func resolveIotBillingGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeBillingGroupOutput)
	svc := meta.(*client.Client).Services().Iot
	return resolveIotTags(ctx, svc, resource, c, i.BillingGroupArn)
}
