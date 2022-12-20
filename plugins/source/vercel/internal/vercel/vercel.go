package vercel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	hc      HTTPDoer
	baseURL string
	token   string
	teamID  string
}

type Paginator struct {
	Count int64  `json:"count"`
	Next  *int64 `json:"next"`
	Prev  *int64 `json:"prev"`
}

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(hc HTTPDoer, baseURL, token, teamID string) *Client {
	return &Client{
		hc:      hc,
		baseURL: baseURL,
		token:   token,
		teamID:  teamID,
	}
}

func (v *Client) WithTeamID(teamID string) *Client {
	vv := *v
	vv.teamID = teamID
	return &vv
}

func (v *Client) Request(ctx context.Context, path string, until *int64, fill any) error {
	body, err := v.request(ctx, path, until)
	if err != nil {
		return err
	}
	defer body.Close()

	//b, _ := io.ReadAll(body)
	//fmt.Println(path, "\n", string(b), "\n")
	//return json.Unmarshal(b, &fill)
	return json.NewDecoder(body).Decode(&fill)
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
		defer res.Body.Close()

		if res.StatusCode == http.StatusTooManyRequests {
			val := res.Header.Get("X-Ratelimit-Reset")
			if val != "" {
				t, err := strconv.ParseInt(val, 10, 64)
				if err == nil && t > 0 {
					ts := time.Unix(t, 0)
					val = ts.Format(time.RFC3339) + fmt.Sprintf(" (in %s)", time.Until(ts).Round(time.Second))
				}

				return nil, fmt.Errorf("request to %s failed: %s. Rate limit will reset at: %s", path, res.Status, val)
			}
		}

		return nil, fmt.Errorf("request to %s failed: %s", path, res.Status)
	}

	return res.Body, nil
}
