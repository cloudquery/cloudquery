package mixpanel

import (
	"context"
	"net/http"
	"time"
)

type Cohort struct {
	Count       int64     `json:"count"`
	IsVisible   bool      `json:"is_visible"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	ProjectID   int64     `json:"project_id"`
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
}

func (c *Client) ListCohorts(ctx context.Context) ([]Cohort, error) {
	var l []Cohort

	err := c.Request(ctx, http.MethodGet, "/api/2.0/cohorts/list", nil, &l)
	return l, err
}
