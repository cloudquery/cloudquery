package mixpanel

import (
	"context"
	"net/http"
	"strconv"
)

type Annotation struct {
	Date        Time   `json:"date"`
	Description string `json:"description"`
	ID          int64  `json:"id"`
	User        any    `json:"user"`
}

type CommonResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type AnnotationList struct {
	CommonResponse
	Results []Annotation `json:"results"`
}

func (c *Client) ListAnnotations(ctx context.Context) (*AnnotationList, error) {
	var l AnnotationList

	if err := c.Request(ctx, http.MethodGet, "/api/app/projects/"+strconv.FormatInt(c.opts.ProjectID, 10)+"/annotations", nil, &l); err != nil {
		return nil, err
	}
	return &l, nil
}
