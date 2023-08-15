package apps

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/digitalocean/godo"
)

func Apps() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_apps",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/apps_get",
		Resolver:    fetchApps,
		Transform:   transformers.TransformWithStruct(&godo.App{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       types.ExtensionTypes.UUID,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: schema.Tables{
			alerts(),
		},
	}
}
