package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DevEndpoints() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_dev_endpoints",
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_DevEndpoint.html`,
		Resolver:    fetchGlueDevEndpoints,
		Transform:   transformers.TransformWithStruct(&types.DevEndpoint{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueDevEndpointArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueDevEndpointTags,
			},
		},
	}
}
