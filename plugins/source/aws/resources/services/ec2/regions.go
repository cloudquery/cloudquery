// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Regions() *schema.Table {
	return &schema.Table{
		Name:        "aws_regions",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html`,
		Resolver:    fetchEc2Regions,
		Multiplex:   client.AccountMultiplex,
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
			{
				Name:     "endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "opt_in_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OptInStatus"),
			},
		},
	}
}
