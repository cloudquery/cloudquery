package registries

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func repositories() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_registry_repositories",
		Description: "Deprecated. https://docs.digitalocean.com/reference/api/api-reference/#operation/registry_list_repositories",
		Resolver:    fetchRegistriesRepositories,
		Transform:   transformers.TransformWithStruct(&godo.Repository{}),
		Columns: []schema.Column{
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
		},
	}
}
