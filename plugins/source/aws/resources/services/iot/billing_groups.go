// Code generated by codegen; DO NOT EDIT.

package iot

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BillingGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_billing_groups",
		Resolver:  fetchIotBillingGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("iot"),
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
			{
				Name:     "billing_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BillingGroupId"),
			},
			{
				Name:     "billing_group_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BillingGroupMetadata"),
			},
			{
				Name:     "billing_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BillingGroupName"),
			},
			{
				Name:     "billing_group_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BillingGroupProperties"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
