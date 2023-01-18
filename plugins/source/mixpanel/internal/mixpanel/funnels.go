package mixpanel

import "context"

type Funnel struct {
	FunnelID int64  `json:"funnnel_id"`
	Name     string `json:"name"`
}

func (c *Client) ListFunnels(ctx context.Context) ([]Funnel, error) {
	var l []Funnel

	err := c.Request(ctx, "/api/2.0/funnels/list", nil, &l)
	return l, err
}
