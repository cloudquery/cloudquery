package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Things() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_things",
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_ThingAttribute.html`,
		Resolver:    fetchIotThings,
		Transform:   transformers.TransformWithStruct(&types.ThingAttribute{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
		},
	}
}
