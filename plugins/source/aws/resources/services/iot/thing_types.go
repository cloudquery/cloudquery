package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ThingTypes() *schema.Table {
	tableName := "aws_iot_thing_types"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_ThingTypeDefinition.html`,
		Resolver:    fetchIotThingTypes,
		Transform:   client.TransformWithStruct(&types.ThingTypeDefinition{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotThingTypeTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingTypeArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
