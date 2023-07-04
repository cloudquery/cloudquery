package cloudiot

import (
	"context"

	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"

	iot "cloud.google.com/go/iot/apiv1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

// Doesn't seem like there's a way to automatically get this list.
// Also this service will deprecated by the end of 2023
var locations = []string{"us-central1", "europe-west1", "asia-east1"}

func fetchDeviceRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	eg, gctx := errgroup.WithContext(ctx)
	for i := range locations {
		location := locations[i]
		eg.Go(func() error {
			req := &pb.ListDeviceRegistriesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/" + location,
			}
			it := gcpClient.ListDeviceRegistries(gctx, req, c.CallOptions...)
			for {
				resp, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return err
				}

				res <- resp
			}
			return nil
		})
	}
	return eg.Wait()
}
