package legacy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const defaultURL = "https://api.snyk.io"

var defaultHTTPClient = http.DefaultClient

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseURL string
	client  HTTPClient
	token   string
}

type Option func(*Client)

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithHTTPClient(client HTTPClient) Option {
	return func(c *Client) {
		c.client = client
	}
}

type bearerAuthClient struct {
	Token  string
	Client HTTPClient
}

func (c *bearerAuthClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "token "+c.Token)
	return c.Client.Do(req)
}

func NewClient(token string, opts ...Option) *Client {
	c := &Client{
		baseURL: defaultURL,
		client:  defaultHTTPClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	// wrap the client with a bearer auth client
	c.client = &bearerAuthClient{
		Token:  token,
		Client: c.client,
	}
	return c
}

type HTTPError struct {
	Body string
	Code int
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("status %d (%v)", e.Code, http.StatusText(e.Code))
}

func (c *Client) post(ctx context.Context, path string, query url.Values) ([]byte, error) {
	uri := fmt.Sprintf("%s?%s", c.baseURL+path, query.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("while creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	r, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		return nil, HTTPError{Body: string(b), Code: r.StatusCode}
	}
	return b, err
}
