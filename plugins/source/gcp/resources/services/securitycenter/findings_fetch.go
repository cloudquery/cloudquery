package securitycenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/protobuf/types/known/timestamppb"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
)

const pageSize = 1000

func getRequest(ctx context.Context, c *client.Client, table string, parent string) (*pb.ListFindingsRequest, error) {
	filter, err := c.Backend.GetKey(ctx, table+c.ID())
	if err != nil {
		return nil, fmt.Errorf("failed to get filter state %w for table %q", err, table)
	}
	req := &pb.ListFindingsRequest{
		Parent:   parent,
		OrderBy:  "event_time",
		PageSize: pageSize,
	}
	if filter != "" {
		req.Filter = filter
	}
	return req, nil
}

func setBackendState(ctx context.Context, c *client.Client, table string, lastEventTime *timestamppb.Timestamp) error {
	err := c.Backend.SetKey(ctx, table+c.ID(), fmt.Sprintf(`event_time >= "%s"`, lastEventTime.AsTime().Format(time.RFC3339)))
	if err != nil {
		return fmt.Errorf("failed to set filter state %w for table %q", err, table)
	}
	return nil
}

func fetchFindings(table string, parent string) func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)
		req, err := getRequest(ctx, c, table, parent)
		if err != nil {
			return err
		}
		gcpClient, err := securitycenter.NewClient(ctx, c.ClientOptions...)
		if err != nil {
			return err
		}
		it := gcpClient.ListFindings(ctx, req, c.CallOptions...)
		itemInPage := 0
		var lastEventTime *timestamppb.Timestamp
		for {
			resp, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			lastEventTime = resp.Finding.EventTime
			if itemInPage >= pageSize {
				// When paginating over a huge result set, we might error out before getting all the results, so we need to set the filter periodically
				err = setBackendState(ctx, c, table, lastEventTime)
				if err != nil {
					return err
				}
				itemInPage = 0
			}
			itemInPage++
			res <- resp
		}
		if lastEventTime != nil {
			err = setBackendState(ctx, c, table, lastEventTime)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
