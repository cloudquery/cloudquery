package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/pipelines"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Pipelines() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_pipelines",
		Resolver:    fetchPipelines,
		Description: "https://developers.hubspot.com/docs/api/crm/pipelines",
		Transform:   transformers.TransformWithStruct(pipelines.Pipeline{}),
		// These are the object types that the pipelines endpoint supports
		// https://developers.hubspot.com/docs/api/crm/pipelines
		Multiplex: client.ObjectTypeMultiplex([]string{"deals", "tickets"}),
		Columns: []schema.Column{
			{
				Name:     "object_type",
				Type:     schema.TypeString,
				Resolver: client.ResolveObjectType,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
