package homebrew

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const defaultURL = "https://formulae.brew.sh"

var defaultHTTPClient = http.DefaultClient

type Client struct {
	baseURL string
	client  *http.Client
}

type Option func(*Client)

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL: defaultURL,
		client:  defaultHTTPClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type HTTPError struct {
	Code int
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("status %d (%v)", e.Code, http.StatusText(e.Code))
}

func (c *Client) get(ctx context.Context, path string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("while creating request: %w", err)
	}
	r, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		r.Body.Close()
		return nil, HTTPError{Code: r.StatusCode}
	}
	return r.Body, nil
}
