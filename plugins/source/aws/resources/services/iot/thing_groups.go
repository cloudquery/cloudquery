package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ThingGroups() *schema.Table {
	tableName := "aws_iot_thing_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeThingGroup.html`,
		Resolver:    fetchIotThingGroups,
		Transform:   transformers.TransformWithStruct(&iot.DescribeThingGroupOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "things_in_group",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotThingGroupThingsInGroup,
			},
			{
				Name:     "policies",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotThingGroupPolicies,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotThingGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
