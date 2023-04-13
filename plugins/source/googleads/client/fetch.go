package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/services"
)

func Fetcher[RESOURCE any](
	extract RowExtractor[*RESOURCE],
	queryOptions ...*gaql.Options,
) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		c := meta.(*Client)
		ctx = c.OutgoingContext(ctx)

		req := &services.SearchGoogleAdsStreamRequest{
			CustomerId:        c.CustomerID,
			Query:             gaql.Query(new(RESOURCE), parent, queryOptions...),
			SummaryRowSetting: enums.SummaryRowSettingEnum_NO_SUMMARY_ROW,
		}

		resp, err := c.GoogleAdsClient.SearchStream(ctx, req)
		if err != nil {
			return err
		}

		return ReceiveStream(resp.Recv, extract, res)
	}
}
