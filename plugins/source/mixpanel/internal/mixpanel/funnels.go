package mixpanel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Funnel struct {
	FunnelID int64  `json:"funnel_id"`
	Name     string `json:"name"`
}

type FunnelData struct {
	Date     time.Time      `json:"date"`
	Steps    []any          `json:"steps"`
	Analysis map[string]any `json:"analysis"`
}

func (c *Client) ListFunnels(ctx context.Context) ([]Funnel, error) {
	var l []Funnel

	err := c.Request(ctx, http.MethodGet, "/api/2.0/funnels/list", nil, &l)
	return l, err
}

func (c *Client) QueryFunnel(ctx context.Context, id int64, startDate, endDate string) ([]FunnelData, error) {
	qp := url.Values{}
	qp.Set("funnel_id", strconv.FormatInt(id, 10))
	qp.Set("from_date", startDate)
	qp.Set("to_date", endDate)

	var d struct {
		Data map[string]FunnelData `json:"data"`
		//Meta any                   `json:"meta"`
	}
	err := c.Request(ctx, http.MethodGet, "/api/2.0/funnels", qp, &d)
	if err != nil {
		return nil, err
	}

	ret := make([]FunnelData, 0, len(d.Data))
	for k, v := range d.Data {
		v.Date, err = time.Parse("2006-01-02", k)
		if err != nil {
			return nil, fmt.Errorf("invalid date key in data: %q", k)
		}

		ret = append(ret, v)
	}

	return ret, nil
}
