package securitycenter

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
)

func fetchFindings(parent string) func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)
		req := &pb.ListFindingsRequest{
			Parent:   parent,
			PageSize: 1000,
		}
		gcpClient, err := securitycenter.NewClient(ctx, c.ClientOptions...)
		if err != nil {
			return err
		}
		it := gcpClient.ListFindings(ctx, req, c.CallOptions...)
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
}
