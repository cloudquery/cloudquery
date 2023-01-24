package mixpanel

import (
	"context"
	"math"
	"net/http"
	"net/url"
)

type EngageProfileList struct {
	CommonResponse
	EngagePaginator

	Data []EngageProfile `json:"-"`
}

type EngageProfileInResponse struct {
	DistinctID string         `json:"$distinct_id"`
	Properties map[string]any `json:"$properties"`
}

type EngageProfile struct {
	DistinctID string         `json:"distinct_id"`
	Properties map[string]any `json:"properties"`
}

type EngagePaginator struct {
	Page      int64  `json:"page"`
	SessionID string `json:"session_id"`
	PageSize  int64  `json:"page_size"`
	Total     int64  `json:"total"`

	TotalPages int64 `json:"-"`
}

func (c *Client) EngageProfiles(ctx context.Context, qp url.Values) (*EngageProfileList, error) {
	var d struct {
		EngageProfileList
		Results []EngageProfileInResponse `json:"results"`
	}
	err := c.Request(ctx, http.MethodPost, "/api/2.0/engage", qp, &d)
	if err != nil {
		return nil, err
	}

	d.Data = make([]EngageProfile, len(d.Results))
	for i := range d.Results {
		d.Data[i] = EngageProfile{
			DistinctID: d.Results[i].DistinctID,
			Properties: d.Results[i].Properties,
		}
	}
	d.Results = nil

	if d.PageSize > 0 && d.Total > 0 {
		d.TotalPages = int64(math.Ceil(float64(d.Total) / float64(d.PageSize)))
	}

	return &d.EngageProfileList, nil
}
