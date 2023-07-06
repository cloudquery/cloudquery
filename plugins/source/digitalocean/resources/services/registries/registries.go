package registries

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_registries",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Container-Registry",
		Resolver:    fetchRegistriesRegistries,
		Transform:   transformers.TransformWithStruct(&godo.Registry{}),
		Columns: []schema.Column{
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			repositories(),
		},
	}
}
