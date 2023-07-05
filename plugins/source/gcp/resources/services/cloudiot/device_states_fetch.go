package cloudiot

import (
	"context"

	iot "cloud.google.com/go/iot/apiv1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchDeviceStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListDeviceStatesRequest{
		Name: parent.Item.(*pb.Device).Name,
	}
	gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	resp, err := gcpClient.ListDeviceStates(ctx, req, c.CallOptions...)
	if err != nil {
		return err
	}
	res <- resp.DeviceStates
	return nil
}
