package mixpanel

import (
	"context"
	"math"
	"net/http"
	"net/url"
)

type EngageProfileList struct {
	CommonResponse

	Results   []EngageProfile `json:"results"`
	Page      int64           `json:"page"`
	SessionID string          `json:"session_id"`
	PageSize  int64           `json:"page_size"`
	Total     int64           `json:"total"`
	//ComputedAt time.Time `json:"computed_at"`

	TotalPages int64 `json:"-"`
}

type EngageProfile map[string]any

func (c *Client) EngageProfiles(ctx context.Context, qp url.Values) (*EngageProfileList, error) {
	var d EngageProfileList
	err := c.Request(ctx, http.MethodPost, "/api/2.0/engage", qp, &d)
	if err != nil {
		return nil, err
	}

	if d.PageSize > 0 {
		d.TotalPages = int64(math.Ceil(float64(d.Total) / float64(d.PageSize)))
	}

	return &d, nil
}
