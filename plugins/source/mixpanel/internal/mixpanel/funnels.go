package mixpanel

import (
	"context"
	"net/url"
	"strconv"
)

type Funnel struct {
	FunnelID int64  `json:"funnnel_id"`
	Name     string `json:"name"`
}

type FunnelData struct {
	Meta any `json:"meta"`
	Data any `json:"data"`
}

func (c *Client) ListFunnels(ctx context.Context) ([]Funnel, error) {
	var l []Funnel

	err := c.Request(ctx, "/api/2.0/funnels/list", nil, &l)
	return l, err
}

func (c *Client) QueryFunnel(ctx context.Context, id int64, startDate, endDate string) (*FunnelData, error) {
	qp := url.Values{}
	qp.Set("funnel_id", strconv.FormatInt(id, 10))
	qp.Set("from_date", startDate)
	qp.Set("to_date", endDate)

	var d FunnelData
	err := c.Request(ctx, "/api/2.0/funnels", qp, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
