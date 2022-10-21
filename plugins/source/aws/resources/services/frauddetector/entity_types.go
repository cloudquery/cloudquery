// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EntityTypes() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_entity_types",
		Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_EntityType.html",
		Resolver:    fetchFrauddetectorEntityTypes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("frauddetector"),
		Columns: []schema.Column{
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
	}
}
