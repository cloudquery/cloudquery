package spaces

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
)

func Spaces() *schema.Table {
	return &schema.Table{
		Name:                 "digitalocean_spaces",
		Resolver:             fetchSpacesSpaces,
		PostResourceResolver: resolveSpaceAttributes,
		Multiplex:            client.SpacesRegionMultiplex,
		Transform:            transformers.TransformWithStruct(&WrappedBucket{}),
		Columns: []schema.Column{
			{
				Name:     "acls",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ACLs"),
			},
			{
				Name:     "bucket",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Bucket"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "public",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Public"),
			},
		},

		Relations: []*schema.Table{
			Cors(),
		},
	}
}
