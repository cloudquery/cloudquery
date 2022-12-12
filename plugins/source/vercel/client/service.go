package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type ServiceClient struct {
	http.Client
	baseURL string
	token   string
	teamID  string
}

type VercelPaginator struct {
	Count int64  `json:"count"`
	Next  *int64 `json:"next"`
	Prev  *int64 `json:"prev"`
}

func (s *ServiceClient) Request(ctx context.Context, path string, until *int64, fill interface{}) error {
	body, err := s.request(ctx, path, until)
	if err != nil {
		return err
	}
	defer body.Close()

	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
	return json.Unmarshal(b, &fill)
	//return json.NewDecoder(body).Decode(&fill)
}

func (s *ServiceClient) request(ctx context.Context, path string, until *int64) (io.ReadCloser, error) {
	u := s.baseURL + path
	uv := url.Values{}
	uv.Set("limit", "100") // Maximum limit
	if until != nil {
		uv.Set("until", strconv.FormatInt(*until, 10))
	}
	if s.teamID != "" {
		uv.Set("teamId", s.teamID)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u+"?"+uv.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)
	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("request failed: %s", res.Status)
	}

	return res.Body, nil
}
