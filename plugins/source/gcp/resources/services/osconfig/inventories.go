package osconfig

import (
	"context"
	"errors"

	osconfig "cloud.google.com/go/osconfig/apiv1"
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func Inventories() *schema.Table {
	return &schema.Table{
		Name:        "gcp_osconfig_inventories",
		Description: `https://cloud.google.com/compute/docs/osconfig/rest/v1/projects.locations.instances.inventories#Inventory`,
		Resolver:    fetchInventories,
		Multiplex:   client.ProjectLocationMultiplexEnabledServices("osconfig.googleapis.com", nil),
		Transform:   client.TransformWithStruct(new(pb.Inventory), transformers.WithPrimaryKeys("Name")),
		Columns:     schema.ColumnList{client.ProjectIDColumn(false)},
	}
}

func fetchInventories(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListInventoriesRequest{
		Parent: "projects/" + c.ProjectId + "/locations/" + c.Location + "/instances/-",
		View:   pb.InventoryView_FULL,
	}

	gcpClient, err := osconfig.NewOsConfigZonalRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	defer gcpClient.Close()
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
