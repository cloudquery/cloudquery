package apps

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/digitalocean/godo"
)

func alerts() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_apps_alerts",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/apps_list_alerts",
		Resolver:    fetchAppsAlerts,
		Transform:   transformers.TransformWithStruct(&godo.AppAlert{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       types.ExtensionTypes.UUID,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
