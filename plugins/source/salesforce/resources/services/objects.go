package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/salesforce/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func rawResolver(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, r.Item)
}

func objectResolver(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return r.Set(c.Name, cl.Object)
}

func Objects() *schema.Table {
	return &schema.Table{
		Name:        "salesforce_objects",
		Description: "https://developer.salesforce.com/docs/atlas.en-us.object_reference.meta/object_reference/sforce_api_objects_list.htm",
		Resolver:    fetchObjects,
		Multiplex:   client.MultiplexStandardObjects,
		Columns: []schema.Column{
			{
				Name:        "_cq_raw",
				Description: "Raw JSON response",
				Type:        schema.TypeJSON,
				Resolver:    rawResolver,
			},
			{
				Name:        "object_type",
				Description: "Name of the object.",
				Type:        schema.TypeString,
				Resolver:    objectResolver,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "id",
				Description: "Unique identifier for the object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
