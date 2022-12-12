package vercel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	hc      http.Client
	baseURL string
	token   string
	teamID  string
}

type Paginator struct {
	Count int64  `json:"count"`
	Next  *int64 `json:"next"`
	Prev  *int64 `json:"prev"`
}

func New(hc http.Client, baseURL, token, teamID string) *Client {
	return &Client{
		hc:      hc,
		baseURL: baseURL,
		token:   token,
		teamID:  teamID,
	}
}

func (v *Client) Request(ctx context.Context, path string, until *int64, fill interface{}) error {
	body, err := v.request(ctx, path, until)
	if err != nil {
		return err
	}
	defer body.Close()

	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
	return json.Unmarshal(b, &fill)
	//return json.NewDecoder(body).Decode(&fill)
}

func (v *Client) request(ctx context.Context, path string, until *int64) (io.ReadCloser, error) {
	u := v.baseURL + path
	uv := url.Values{}
	uv.Set("limit", "100") // Maximum limit
	if until != nil {
		uv.Set("until", strconv.FormatInt(*until, 10))
	}
	if v.teamID != "" {
		uv.Set("teamId", v.teamID)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u+"?"+uv.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+v.token)
	res, err := v.hc.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("request failed: %s", res.Status)
	}

	return res.Body, nil
}
