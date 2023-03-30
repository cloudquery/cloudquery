package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecurityProfiles() *schema.Table {
	tableName := "aws_iot_security_profiles"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeSecurityProfile.html`,
		Resolver:    fetchIotSecurityProfiles,
		Transform:   client.TransformWithStruct(&iot.DescribeSecurityProfileOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "targets",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotSecurityProfileTargets,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotSecurityProfileTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityProfileArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
