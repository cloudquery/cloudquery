package keys

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
