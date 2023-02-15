package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

func fetch(tableName string, request *analyticsdata.RunReportRequest) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		c := meta.(*Client)
		logger := c.Logger().With().Str("table", tableName).Logger()

		req := c.service.Properties.RunReport(c.PropertyID, request).Context(ctx)

		dates, err := genDates(ctx, c, tableName)
		if err != nil {
			return err
		}

		for date := range dates {
			dateStr := date.Format(layout)
			request.DateRanges = []*analyticsdata.DateRange{{StartDate: dateStr, EndDate: dateStr}}

			var fetched int64
			var gotOther bool
			for {
				request.Offset = fetched

				resp, err := req.Do()
				if err != nil {
					return err
				}

				res <- convertRows(resp, date)

				fetched += int64(len(resp.Rows))

				if fetched >= resp.RowCount {
					break
				}

				gotOther = gotOther || resp.Metadata.DataLossFromOtherRow
			}

			if gotOther {
				logger.Warn().Str("date", dateStr).Msg("got (other) row, consider modifying report")
			}

			// We save current date here, even with data loss (as the report should be edited to get all the data)
			if err := c.backend.Set(ctx, tableName, c.ID(), dateStr); err != nil {
				logger.Err(err).Msg("failed to save state")
				return err
			}
		}
		return nil
	}
}
