package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDirectConnectLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeLocationsInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeLocations(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.Locations
	return nil
}
