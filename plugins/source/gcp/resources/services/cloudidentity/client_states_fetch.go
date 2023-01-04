package cloudidentity

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/cloudidentity/v1"
)

func fetchClientStates(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	identityClient, err := cloudidentity.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	nextPageToken := ""
	for {
		output, err := identityClient.Devices.DeviceUsers.ClientStates.List("devices/-/deviceUsers/-").PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		res <- output.ClientStates

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
