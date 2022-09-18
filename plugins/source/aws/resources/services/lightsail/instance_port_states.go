// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InstancePortStates() *schema.Table {
	return &schema.Table{
		Name:      "aws_lightsail_instance_port_states",
		Resolver:  fetchLightsailInstancePortStates,
		Multiplex: client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "cidr_list_aliases",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CidrListAliases"),
			},
			{
				Name:     "cidrs",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Cidrs"),
			},
			{
				Name:     "from_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("FromPort"),
			},
			{
				Name:     "ipv_6_cidrs",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Ipv6Cidrs"),
			},
			{
				Name:     "protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Protocol"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "to_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ToPort"),
			},
		},
	}
}
