package cloudiot

import (
	"context"

	pb "cloud.google.com/go/iot/apiv1/iotpb"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	iot "cloud.google.com/go/iot/apiv1"
)

func fetchDeviceConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListDeviceConfigVersionsRequest{
		Name: parent.Item.(*pb.Device).Name,
	}
	gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	resp, err := gcpClient.ListDeviceConfigVersions(ctx, req, c.CallOptions...)
	if err != nil {
		return err
	}
	res <- resp.DeviceConfigs
	return nil
}
