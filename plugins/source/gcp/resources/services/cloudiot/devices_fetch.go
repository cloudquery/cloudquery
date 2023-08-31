package cloudiot

import (
	"context"
	"strings"

	"google.golang.org/api/iterator"

	iot "cloud.google.com/go/iot/apiv1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func getFieldsMask() []string {
	fieldMaskPaths := []string{}
	for _, col := range Devices().Columns {
		if col.Name == "project_id" || strings.HasPrefix(col.Name, "_") {
			continue
		}
		fieldMaskPaths = append(fieldMaskPaths, col.Name)
	}
	return fieldMaskPaths
}

func fetchDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	req := &pb.ListDevicesRequest{
		Parent:    parent.Item.(*pb.DeviceRegistry).Name,
		FieldMask: &fieldmaskpb.FieldMask{Paths: getFieldsMask()},
	}
	gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListDevices(ctx, req, c.CallOptions...)
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
}
