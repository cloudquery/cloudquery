// Code generated by codegen; DO NOT EDIT.

package spaces

import (
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
)

func Spaces() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_spaces",
		Resolver:  fetchSpacesSpaces,
		Multiplex: client.SpacesRegionMultiplex,
		Columns: []schema.Column{
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
			{
				Name:     "ac_ls",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ACLs"),
			},
		},

		Relations: []*schema.Table{
			Cors(),
		},
	}
}
