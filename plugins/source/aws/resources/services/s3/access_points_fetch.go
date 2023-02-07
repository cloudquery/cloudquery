package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAccessPoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	svc := c.Services().S3control
	paginator := s3control.NewListAccessPointsPaginator(svc, &s3control.ListAccessPointsInput{})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.AccessPointList
	}

	return nil
}
