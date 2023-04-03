package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Regions() *schema.Table {
	tableName := "aws_ec2_regions"
	return &schema.Table{
		Name:        "aws_regions",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html`,
		Resolver:    fetchEc2Regions,
		Multiplex:   client.AccountMultiplex(tableName),
		Transform:   transformers.TransformWithStruct(&types.Region{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: resolveRegionEnabled,
			},
			{
				Name:     "partition",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSPartition,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionName"),
			},
		},
	}
}

func fetchEc2Regions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	output, err := c.Services().Ec2.DescribeRegions(ctx, &ec2.DescribeRegionsInput{AllRegions: aws.Bool(true)})
	if err != nil {
		return err
	}
	res <- output.Regions
	return nil
}

func resolveRegionEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := resource.Item.(types.Region)
	switch *region.OptInStatus {
	case "opt-in-not-required", "opted-in":
		return resource.Set(c.Name, true)
	case "not-opted-in":
		return resource.Set(c.Name, false)
	}
	return nil
}
