package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BillingGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_billing_groups",
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeBillingGroup.html`,
		Resolver:    fetchIotBillingGroups,
		Transform:   transformers.TransformWithStruct(&iot.DescribeBillingGroupOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
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
		},
	}
}
