package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AvailabilityZones() *schema.Table {
	return &schema.Table{
		Name:        "aws_availability_zones",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html`,
		Resolver:    fetchAvailabilityZones,
		Multiplex:   client.AccountMultiplex,
		Transform:   transformers.TransformWithStruct(&types.AvailabilityZone{}, transformers.WithPrimaryKeys("RegionName", "ZoneId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: resolveAZEnabled,
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
