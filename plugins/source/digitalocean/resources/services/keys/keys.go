package keys

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_keys",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/SSH-Keys",
		Resolver:    fetchKeysKeys,
		Transform:   transformers.TransformWithStruct(&godo.Key{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
