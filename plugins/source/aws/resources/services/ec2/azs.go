package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AvailabilityZones() *schema.Table {
	tableName := "aws_availability_zones"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AvailabilityZone.html`,
		Resolver:    fetchAvailabilityZones,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.AvailabilityZone{}, transformers.WithPrimaryKeys("RegionName", "ZoneId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "enabled",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: resolveAZEnabled,
			},
			{
				Name:     "partition",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAWSPartition,
			},
			{
				Name:     "region",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("RegionName"),
			},
		},
	}
}

func fetchAvailabilityZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	output, err := svc.DescribeAvailabilityZones(ctx, &ec2.DescribeAvailabilityZonesInput{AllAvailabilityZones: aws.Bool(true)}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.AvailabilityZones
	return nil
}

func resolveAZEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := resource.Item.(types.AvailabilityZone)
	switch region.OptInStatus {
	case "opt-in-not-required", "opted-in":
		return resource.Set(c.Name, true)
	case "not-opted-in":
		return resource.Set(c.Name, false)
	}
	return nil
}
