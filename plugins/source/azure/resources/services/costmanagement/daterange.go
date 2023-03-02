package costmanagement

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

func timePtr(t time.Time) *time.Time {
	return &t
}

func ttPtr(v armcostmanagement.TimeframeType) *armcostmanagement.TimeframeType {
	return &v
}

func dateRangeToTimeFrame(now time.Time, dr string) (*armcostmanagement.TimeframeType, *armcostmanagement.QueryTimePeriod) {
	quarterStart := func(t time.Time) time.Time {
		var m time.Month
		switch t.Month() {
		case time.January, time.February, time.March:
			m = time.January
		case time.April, time.May, time.June:
			m = time.April
		case time.July, time.August, time.September:
			m = time.July
		case time.October, time.November, time.December:
			m = time.October
		}
		return time.Date(t.Year(), m, 1, 0, 0, 0, 0, t.Location())
	}

	now = now.Truncate(24 * time.Hour)

	switch dr {
	case "Last7Days":
		now = now.AddDate(0, 0, -1) // ends yesterday
		return ttPtr(armcostmanagement.TimeframeTypeCustom), &armcostmanagement.QueryTimePeriod{
			From: timePtr(now.AddDate(0, 0, -6)),
			To:   timePtr(now),
		}
	case "Last30Days":
		now = now.AddDate(0, 0, -1) // ends yesterday
		return ttPtr(armcostmanagement.TimeframeTypeCustom), &armcostmanagement.QueryTimePeriod{
			From: timePtr(now.AddDate(0, 0, -29)),
			To:   timePtr(now),
		}
	case "Last3Months", "Last6Months", "Last12Months":
		e := now.AddDate(0, 0, -now.Day()+1).Add(-time.Second) // last second of last month
		s := e.AddDate(0, map[string]int{
			"Last3Months":  -2,
			"Last6Months":  -5,
			"Last12Months": -11,
		}[dr], 0)
		s = s.AddDate(0, 0, -s.Day()+1).Truncate(24 * time.Hour) // first day
		return ttPtr(armcostmanagement.TimeframeTypeCustom), &armcostmanagement.QueryTimePeriod{
			From: &s,
			To:   &e,
		}
	case "ThisMonth":
		return ttPtr(armcostmanagement.TimeframeTypeMonthToDate), nil
	case "LastMonth":
		return ttPtr(armcostmanagement.TimeframeTypeTheLastMonth), nil
	case "CurrentBillingPeriod":
		return ttPtr(armcostmanagement.TimeframeTypeBillingMonthToDate), nil
	case "LastBillingPeriod":
		return ttPtr(armcostmanagement.TimeframeTypeTheLastBillingMonth), nil
	case "ThisQuarter":
		qs := quarterStart(now)
		next := qs.AddDate(0, 3, 15)               // 15 days into next quarter
		qe := quarterStart(next).Add(-time.Second) // one second before first day of this quarter
		return ttPtr(armcostmanagement.TimeframeTypeCustom), &armcostmanagement.QueryTimePeriod{
			From: &qs,
			To:   &qe,
		}
	case "LastQuarter":
		qe := quarterStart(now).Add(-time.Second) // one second before first day of this quarter
		qs := quarterStart(qe)
		return ttPtr(armcostmanagement.TimeframeTypeCustom), &armcostmanagement.QueryTimePeriod{
			From: &qs,
			To:   &qe,
		}
	case "ThisYear":
		s := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
		e := time.Date(now.Year(), time.December, 31, 23, 59, 59, 0, now.Location())
		return ttPtr(armcostmanagement.TimeframeTypeCustom), &armcostmanagement.QueryTimePeriod{
			From: &s,
			To:   &e,
		}
	default:
		return nil, nil
	}
}
