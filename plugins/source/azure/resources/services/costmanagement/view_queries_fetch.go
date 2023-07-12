package costmanagement

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchViewQueries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	item := parent.Item.(*armcostmanagement.View)
	if item.Properties == nil {
		return nil
	}

	svc, err := armcostmanagement.NewQueryClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}

	b, err := json.Marshal(item.Properties.Query)
	if err != nil {
		return err
	}
	var qd armcostmanagement.QueryDefinition
	if err := json.Unmarshal(b, &qd); err != nil {
		return err
	}

	if (qd.Timeframe == nil || *qd.Timeframe != armcostmanagement.TimeframeTypeCustom) && qd.TimePeriod == nil && item.Properties.DateRange != nil {
		qd.Timeframe, qd.TimePeriod = dateRangeToTimeFrame(time.Now().UTC(), *item.Properties.DateRange)
		if qd.Timeframe != nil {
			cl.Logger().Debug().Any("report_name", item.Name).Str("date_range", *item.Properties.DateRange).Any("resulting_time_period", qd.TimePeriod).Any("resulting_timeframe", qd.Timeframe).Msg("setting time period from date range")
		}
	}

	if qd.Timeframe == nil && item.Properties.DateRange != nil {
		return fmt.Errorf("could not convert date range %q to time period", *item.Properties.DateRange)
	}

	data, err := svc.Usage(ctx, *item.Properties.Scope, qd, nil)
	if err != nil {
		return err
	}

	res <- data
	return nil
}
