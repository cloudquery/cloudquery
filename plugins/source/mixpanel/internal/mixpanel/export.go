package mixpanel

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ExportEvent struct {
	Event      string         `json:"event"`
	Properties map[string]any `json:"properties"`
}

func (c *Client) ExportData(ctx context.Context, startDate, endDate string) ([]ExportEvent, error) {
	qp := url.Values{}
	qp.Set("from_date", startDate)
	qp.Set("to_date", endDate)

	body, err := c.RequestWithReader(ctx, http.MethodGet, "/api/2.0/export", qp)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var data []ExportEvent
	s := bufio.NewScanner(body)
	line := 0
	for s.Scan() {
		var d ExportEvent
		if err := json.Unmarshal(s.Bytes(), &d); err != nil {
			return nil, fmt.Errorf("error parsing line %d: %w", line, err)
		}

		data = append(data, d)
		line++
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("scanner failed on line %d: %w", line, err)
	}

	return data, nil
}
