package tenable

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const defaultBaseURL = "https://cloud.tenable.com"

type Client struct {
	baseURL    string
	httpClient *http.Client
	accessKey  string
	secretKey  string
}

type Option func(*Client)

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func NewClient(accessKey, secretKey string, opts ...Option) *Client {
	c := &Client{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{},
		accessKey:  accessKey,
		secretKey:  secretKey,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) newRequest(ctx context.Context, method, path string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s; secretKey=%s", c.accessKey, c.secretKey))
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ServerStatus represents the Tenable.io server status response.
type ServerStatus struct {
	Status string `json:"status"`
}

// GetServerStatus calls the /server/status endpoint to verify connectivity and credentials.
func (c *Client) GetServerStatus(ctx context.Context) (*ServerStatus, error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/server/status")
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// ok
	case http.StatusUnauthorized, http.StatusForbidden:
		return nil, ErrUnauthorized
	default:
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var status ServerStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &status, nil
}
