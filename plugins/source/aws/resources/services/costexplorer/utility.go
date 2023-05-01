package costexplorer

import "time"

func beginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func endOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day()+1)
}

func forecastEndOfMonth(now time.Time) time.Time {
	if now.Format("2006-01-02") == endOfMonth(now).Format("2006-01-02") {
		return now.AddDate(0, 0, 1)
	}
	return endOfMonth(now).AddDate(0, 0, 1)
}
