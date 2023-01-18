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

	"github.com/rs/zerolog"
)

type Client struct {
	logger  zerolog.Logger
	hc      HTTPDoer
	baseURL string
	token   string
	teamID  string

	maxRetries int64
	maxWait    int64 // in seconds
	pageSize   int64
}

type Paginator struct {
	Count int64  `json:"count"`
	Next  *int64 `json:"next"`
	Prev  *int64 `json:"prev"`
}

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(logger zerolog.Logger, hc HTTPDoer, baseURL, token, teamID string, maxRetries, maxWait, pageSize int64) *Client {
	return &Client{
		logger:  logger,
		hc:      hc,
		baseURL: baseURL,
		token:   token,
		teamID:  teamID,

		maxRetries: maxRetries,
		maxWait:    maxWait,
		pageSize:   pageSize,
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
	uv.Set("limit", strconv.FormatInt(v.pageSize, 10))
	if until != nil {
		uv.Set("until", strconv.FormatInt(*until, 10))
	}
	if v.teamID != "" {
		uv.Set("teamId", v.teamID)
	}

	retries := int64(0)
	for {
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

		if res.StatusCode == http.StatusOK {
			return res.Body, nil
		}

		_ = res.Body.Close()
		retries++
		if retries > v.maxRetries {
			break
		}

		val := res.Header.Get("X-Ratelimit-Reset")
		rateErr := fmt.Errorf("request to %s failed: %s", path, res.Status)
		if val == "" || res.StatusCode != http.StatusTooManyRequests {
			return nil, rateErr
		}

		t, err := strconv.ParseInt(val, 10, 64)
		if err == nil && t > 0 {
			ts := time.Unix(t, 0)
			secsLeft := time.Until(ts).Round(time.Second)
			if secsLeft > 0 && secsLeft < time.Duration(v.maxWait)*time.Second {
				v.logger.Info().Dur("wait", secsLeft).Msg("waiting for rate limit reset")
				select {
				case <-ctx.Done():
					return nil, rateErr
				case <-time.After(secsLeft):
				}
				continue // next retry
			}

			val = ts.Format(time.RFC3339) + fmt.Sprintf(" (in %s)", secsLeft)
		}

		return nil, fmt.Errorf("request to %s failed: %s. Rate limit will reset at: %s", path, res.Status, val)
	}

	return nil, fmt.Errorf("exceeded max retries")
}
