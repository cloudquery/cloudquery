// Code generated by codegen; DO NOT EDIT.

package artifactregistry

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_locations",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/Shared.Types/ListLocationsResponse#Location`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocationId"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Repositories(),
		},
	}
}
