package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Regions() *schema.Table {
	return &schema.Table{
		Name:        "aws_regions",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html`,
		Resolver:    fetchEc2Regions,
		Multiplex:   client.AccountMultiplex,
		Transform:   transformers.TransformWithStruct(&types.Region{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: resolveRegionEnabled,
			},
			{
				Name:     "partition",
				Type:     schema.TypeString,
				Resolver: resolveRegionPartition,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionName"),
			},
		},
	}
}
