package mixpanel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type EngageRevenue struct {
	Date      time.Time `json:"date"`
	Amount    float64   `json:"amount"`
	Count     int64     `json:"count"`
	PaidCount int64     `json:"paid_count"`
}

func (c *Client) ListEngageRevenues(ctx context.Context, startDate, endDate string) ([]EngageRevenue, error) {
	qp := url.Values{}
	qp.Set("from_date", startDate)
	qp.Set("to_date", endDate)

	var d struct {
		Data map[string]EngageRevenue `json:"data"`
	}
	err := c.Request(ctx, http.MethodPost, "/api/2.0/engage/revenue", qp, &d)
	if err != nil {
		return nil, err
	}

	ret := make([]EngageRevenue, 0, len(d.Data))
	for k, v := range d.Data {
		if k == "$overall" {
			continue
		}

		v.Date, err = time.Parse("2006-01-02", k)
		if err != nil {
			return nil, fmt.Errorf("invalid date key in data: %q", k)
		}

		ret = append(ret, v)
	}

	return ret, nil
}
