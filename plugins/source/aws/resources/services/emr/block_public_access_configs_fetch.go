package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEmrBlockPublicAccessConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Emr
	out, err := svc.GetBlockPublicAccessConfiguration(ctx, &emr.GetBlockPublicAccessConfigurationInput{})
	if err != nil {
		if client.IgnoreNotAvailableRegion(err) {
			return nil
		}
		return err
	}
	res <- out
	return nil
}
