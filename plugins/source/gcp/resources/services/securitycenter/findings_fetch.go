package securitycenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/protobuf/types/known/timestamppb"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
)

const pageSize = 1000

func getRequest(ctx context.Context, c *client.Client, table string, parent string) *pb.ListFindingsRequest {
	filter, err := c.Backend.Get(ctx, table, c.ID())
	if err != nil {
		c.Logger().Warn().Str("table", table).Msgf("failed to get filter %s", err.Error())
	}
	var req *pb.ListFindingsRequest
	if filter == "" || err != nil {
		req = &pb.ListFindingsRequest{
			Parent:   parent,
			OrderBy:  "event_time",
			PageSize: pageSize,
		}
	} else {
		req = &pb.ListFindingsRequest{
			Parent:   parent,
			OrderBy:  "event_time",
			PageSize: pageSize,
			Filter:   filter,
		}
	}
	return req
}

func setBackendState(ctx context.Context, c *client.Client, table string, lastEventTime *timestamppb.Timestamp) {
	err := c.Backend.Set(ctx, table, c.ID(), fmt.Sprintf(`event_time >= "%s"`, lastEventTime.AsTime().Format(time.RFC3339)))
	if err != nil {
		c.Logger().Warn().Str("table", table).Msgf("failed to set filter %s", err.Error())
	}
}

func fetchFindings(table string, parent string) func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)
		req := getRequest(ctx, c, table, parent)
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
				setBackendState(ctx, c, table, lastEventTime)
				itemInPage = 0
			}
			itemInPage++
			res <- resp
		}
		if lastEventTime != nil {
			setBackendState(ctx, c, table, lastEventTime)
		}
		return nil
	}
}
