package spaces

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"

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
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("ACLs"),
			},
			{
				Name:     "bucket",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Bucket"),
			},
			{
				Name:     "location",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "public",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("Public"),
			},
		},

		Relations: []*schema.Table{
			cors(),
		},
	}
}
