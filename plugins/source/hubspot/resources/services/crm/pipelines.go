package crm

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/clarkmcc/go-hubspot/generated/v3/pipelines"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "object_type",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveObjectType,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}
