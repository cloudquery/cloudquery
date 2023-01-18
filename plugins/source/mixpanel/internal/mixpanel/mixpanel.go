package mixpanel

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
	logger                 zerolog.Logger
	hc                     HTTPDoer
	baseURL                string
	apiUser, apiSecret     string
	projectID, workspaceID int64

	maxRetries int64
	pageSize   int64
}

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(logger zerolog.Logger, hc HTTPDoer, baseURL, apiUser, apiSecret string, projectID, workspaceID, maxRetries, pageSize int64) *Client {
	if baseURL == "" {
		baseURL = "https://mixpanel.com"
	}

	return &Client{
		logger:      logger,
		hc:          hc,
		baseURL:     baseURL,
		apiUser:     apiUser,
		apiSecret:   apiSecret,
		projectID:   projectID,
		workspaceID: workspaceID,

		maxRetries: maxRetries,
		pageSize:   pageSize,
	}
}

func (v *Client) Request(ctx context.Context, path string, qp url.Values, fill any) error {
	body, err := v.request(ctx, path, qp)
	if err != nil {
		return err
	}
	defer body.Close()

	//b, _ := io.ReadAll(body)
	//fmt.Println(path, "\n", string(b), "\n")
	//return json.Unmarshal(b, &fill)
	return json.NewDecoder(body).Decode(&fill)
}

func (v *Client) request(ctx context.Context, path string, qp url.Values) (io.ReadCloser, error) {
	u := v.baseURL + path

	if v.projectID > 0 {
		qp.Set("project_id", strconv.FormatInt(v.projectID, 10))
	}
	if v.workspaceID > 0 {
		qp.Set("workspace_id", strconv.FormatInt(v.workspaceID, 10))
	}

	retries := int64(0)
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u+"?"+qp.Encode(), nil)
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth(v.apiUser, v.apiSecret)
		req.Header.Set("Content-Type", "application/json")
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

		rateErr := fmt.Errorf("request to %s failed: %s", path, res.Status)
		if res.StatusCode != http.StatusTooManyRequests {
			return nil, rateErr
		}

		backoff := time.Duration(retries) * 1500 * time.Millisecond
		v.logger.Info().Dur("wait", backoff).Msg("waiting for rate limit reset")
		select {
		case <-ctx.Done():
			return nil, rateErr
		case <-time.After(backoff):
		}
	}

	return nil, fmt.Errorf("exceeded max retries")
}
