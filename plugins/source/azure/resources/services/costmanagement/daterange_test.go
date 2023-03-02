package costmanagement

import (
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/stretchr/testify/assert"
)

func TestDateRangeToTimeframe(t *testing.T) {
	cases := []struct {
		rangeType string
		now       time.Time
		wantTT    *armcostmanagement.TimeframeType
		wantQP    *armcostmanagement.QueryTimePeriod
	}{
		{
			"Last7Days",
			time.Date(2023, 3, 2, 15, 59, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, 2, 23, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			"Last30Days",
			time.Date(2023, 3, 2, 15, 59, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			"ThisMonth",
			time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeMonthToDate),
			nil,
		},
		{
			"LastMonth",
			time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeTheLastMonth),
			nil,
		},
		{
			"Last3Months",
			time.Date(2023, 3, 2, 15, 59, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, 2, 28, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"Last6Months",
			time.Date(2023, 3, 2, 15, 59, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, 2, 28, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"Last12Months",
			time.Date(2023, 3, 2, 15, 59, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, 2, 28, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"ThisQuarter",
			time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, time.March, 31, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"ThisQuarter",
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, time.June, 30, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"ThisQuarter",
			time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, time.September, 30, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"ThisQuarter",
			time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, time.October, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"LastQuarter",
			time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, time.September, 30, 23, 59, 59, 0, time.UTC)),
			},
		},
		{
			"ThisYear",
			time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),
			ttPtr(armcostmanagement.TimeframeTypeCustom),
			&armcostmanagement.QueryTimePeriod{
				From: timePtr(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)),
			},
		},
	}

	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("#%d: %s", i, tc.rangeType), func(t *testing.T) {
			tt, qp := dateRangeToTimeFrame(tc.now, tc.rangeType)
			assert.Equal(t, tc.wantTT, tt)
			assert.Equal(t, tc.wantQP, qp)
		})
	}
}
