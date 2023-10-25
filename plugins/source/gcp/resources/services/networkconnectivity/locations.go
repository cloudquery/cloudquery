package networkconnectivity

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"google.golang.org/api/networkconnectivity/v1"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_networkconnectivity_locations",
		Description: `https://cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest/v1/projects.locations/list`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("networkconnectivity.googleapis.com"),
		Transform:   client.TransformWithStruct(&networkconnectivity.Location{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			InternalRanges(),
		},
	}
}

func fetchLocations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	ncClient, err := networkconnectivity.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := ncClient.Projects.Locations.List("projects/" + c.ProjectId)
	err = it.Pages(ctx, func(resp *networkconnectivity.ListLocationsResponse) error {
		if resp != nil {
			res <- resp.Locations
		}
		return nil
	})

	return err
}
