// Code generated by codegen; DO NOT EDIT.

package iot

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Things() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_things",
		Resolver:  fetchIotThings,
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
				Name:     "principals",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotThingPrincipals,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
			{
				Name:     "thing_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingName"),
			},
			{
				Name:     "thing_type_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingTypeName"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
		},
	}
}
