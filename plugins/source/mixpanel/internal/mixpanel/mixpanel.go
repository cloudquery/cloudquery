package mixpanel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/sync/semaphore"
)

type Client struct {
	opts ClientOptions

	concurrentLimit *semaphore.Weighted
}

const StatusOK = "ok"

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientOptions struct {
	Logger                 zerolog.Logger
	HC                     HTTPDoer
	Region                 Region
	BaseURL                string
	APIUser, APISecret     string
	ProjectID, WorkspaceID int64
	MaxRetries             int64
}

func New(opts ClientOptions) *Client {
	return &Client{
		opts: opts,

		// https://help.mixpanel.com/hc/en-us/articles/115004602563-Rate-Limits-for-Export-API-Endpoints
		concurrentLimit: semaphore.NewWeighted(5),
	}
}

func (v *Client) Request(ctx context.Context, method, path string, qp url.Values, fill any) error {
	body, err := v.RequestWithReader(ctx, method, path, qp)
	if err != nil {
		return err
	}
	defer body.Close()

	return json.NewDecoder(body).Decode(&fill)
}

func (v *Client) RequestWithReader(ctx context.Context, method, path string, qp url.Values) (io.ReadCloser, error) {
	if qp == nil {
		qp = url.Values{}
	}

	isExport := strings.HasPrefix(path, `/api/2.0/export`)
	uri := v.getBaseURL(isExport) + path

	body, err := v.request(ctx, method, uri, qp)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (v *Client) request(ctx context.Context, method, uri string, qp url.Values) (io.ReadCloser, error) {
	if v.opts.ProjectID > 0 {
		qp.Set("project_id", strconv.FormatInt(v.opts.ProjectID, 10))
	}
	if v.opts.WorkspaceID > 0 {
		qp.Set("workspace_id", strconv.FormatInt(v.opts.WorkspaceID, 10))
	}

	retries := int64(0)
	for {
		if err := v.concurrentLimit.Acquire(ctx, 1); err != nil {
			return nil, err
		}

		res, err := func() (*http.Response, error) {
			defer v.concurrentLimit.Release(1)

			req, err := http.NewRequestWithContext(ctx, method, uri+"?"+qp.Encode(), nil)
			if err != nil {
				return nil, err
			}
			req.SetBasicAuth(v.opts.APIUser, v.opts.APISecret)
			req.Header.Set("Content-Type", "application/json")
			return v.opts.HC.Do(req)
		}()

		if err != nil {
			return nil, err
		}

		if res.StatusCode == http.StatusOK {
			return res.Body, nil
		}

		respText, _ := io.ReadAll(res.Body)

		_ = res.Body.Close()
		retries++
		if retries > v.opts.MaxRetries {
			break
		}

		rateErr := fmt.Errorf("request to %s failed: %s %s", uri, res.Status, string(respText))
		if res.StatusCode != http.StatusTooManyRequests {
			return nil, rateErr
		}

		saneMsg := string(respText)
		{
			var v map[string]any
			if err := json.Unmarshal(respText, &v); err == nil {
				if msg, ok := v["error"].(string); ok {
					saneMsg = msg
				}
			}
		}

		backoff := time.Duration(retries) * 1500 * time.Millisecond
		v.opts.Logger.Info().Str("url", uri).Interface("query_params", qp).Dur("wait", backoff).Str("api_message", saneMsg).Msg("waiting for rate limit reset")
		select {
		case <-ctx.Done():
			return nil, rateErr
		case <-time.After(backoff):
		}
	}

	return nil, fmt.Errorf("exceeded max retries")
}

func (v *Client) getBaseURL(isExport bool) string {
	if v.opts.BaseURL != "" {
		return v.opts.BaseURL
	}

	switch v.opts.Region {
	case RegionEU:
		if isExport {
			return "https://data-eu.mixpanel.com"
		}
		return "https://eu.mixpanel.com"
	default:
		if isExport {
			return "https://data.mixpanel.com"
		}
		return "https://mixpanel.com"
	}
}
