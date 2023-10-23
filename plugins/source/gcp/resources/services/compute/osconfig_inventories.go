package compute

import (
	"context"
	"errors"

	"cloud.google.com/go/compute/apiv1/computepb"
	osconfig "cloud.google.com/go/osconfig/apiv1"
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"google.golang.org/api/iterator"
)

func osConfigInventories() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_osconfig_inventories",
		Description: `https://cloud.google.com/compute/docs/osconfig/rest/v1/projects.locations.instances.inventories#Inventory`,
		Resolver:    fetchOSConfigInventories,
		Multiplex:   client.ProjectMultiplexEnabledServices("osconfig.googleapis.com"),
		Transform:   client.TransformWithStruct(new(pb.Inventory), transformers.WithPrimaryKeys("Name")),
		Columns:     schema.ColumnList{client.ProjectIDColumn(false)},
	}
}

func fetchOSConfigInventories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	zone := parent.Item.(*computepb.Zone)
	req := &pb.ListInventoriesRequest{
		Parent: "projects/" + c.ProjectId + "/locations/" + *zone.Name + "/instances/-",
		View:   pb.InventoryView_FULL,
	}

	gcpClient, err := osconfig.NewOsConfigZonalRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListInventories(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return err
		}
		res <- resp
	}
	return nil
}
