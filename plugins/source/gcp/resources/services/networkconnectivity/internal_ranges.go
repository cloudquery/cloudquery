package networkconnectivity

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"google.golang.org/api/networkconnectivity/v1"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func InternalRanges() *schema.Table {
	return &schema.Table{
		Name:        "gcp_networkconnectivity_internal_ranges",
		Description: `https://cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest/v1/projects.locations.internalRanges/list`,
		Resolver:    fetchInternalRanges,
		Multiplex:   client.ProjectMultiplexEnabledServices("networkconnectivity.googleapis.com"),
		Transform:   client.TransformWithStruct(&networkconnectivity.InternalRange{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchInternalRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(*networkconnectivity.Location)

	ns, err := networkconnectivity.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	listInternalRanges := ns.Projects.Locations.InternalRanges.List(p.Name)
	err = listInternalRanges.Pages(ctx, func(resp *networkconnectivity.ListInternalRangesResponse) error {
		if resp != nil {
			res <- resp.InternalRanges
		}
		return nil
	})

	return err
}
