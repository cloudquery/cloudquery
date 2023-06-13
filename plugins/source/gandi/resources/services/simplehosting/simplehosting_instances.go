package simplehosting

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func SimplehostingInstances() *schema.Table {
	return &schema.Table{
		Name:     "gandi_simplehosting_instances",
		Resolver: fetchSimplehostingInstances,
		Columns: []schema.Column{
			{
				Name:        "sharing_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "size",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "status",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "database",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Database"),
			},
			{
				Name:     "language",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Language"),
			},
			{
				Name:     "datacenter",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Datacenter"),
			},
		},

		Relations: []*schema.Table{
			SimplehostingInstanceVhosts(),
		},
	}
}
