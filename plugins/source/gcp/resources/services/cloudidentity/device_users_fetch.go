package cloudidentity

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/cloudidentity/v1"
)

func fetchDeviceUsers(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	identityClient, err := cloudidentity.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	nextPageToken := ""
	for {
		output, err := identityClient.Devices.DeviceUsers.List("devices/-").PageSize(20).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		res <- output.DeviceUsers

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
