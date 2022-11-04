// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReservedInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_reserved_instances",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReservedInstances.html`,
		Resolver:    fetchEc2ReservedInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveReservedInstanceArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "currency_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrencyCode"),
			},
			{
				Name:     "duration",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Duration"),
			},
			{
				Name:     "end",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("End"),
			},
			{
				Name:     "fixed_price",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("FixedPrice"),
			},
			{
				Name:     "instance_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InstanceCount"),
			},
			{
				Name:     "instance_tenancy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceTenancy"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "offering_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OfferingClass"),
			},
			{
				Name:     "offering_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OfferingType"),
			},
			{
				Name:     "product_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductDescription"),
			},
			{
				Name:     "recurring_charges",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RecurringCharges"),
			},
			{
				Name:     "reserved_instances_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservedInstancesId"),
			},
			{
				Name:     "scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Scope"),
			},
			{
				Name:     "start",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Start"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "usage_price",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("UsagePrice"),
			},
		},
	}
}
