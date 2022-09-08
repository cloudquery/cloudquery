package versions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type manifestResponse struct {
	Latest string `json:"latest"`
}

type githubLatestResponse struct {
	TagName string `json:"tag_name"`
	// other fields are ignored
}

// Client interacts with repositories to fetch version information.
// It relies on convention to determine the URL format to use when fetching.
// Official CloudQuery plugin versions are fetched from release manifest files,
// while community plugins are fetched using GithubLatestURL.
type Client struct {
	cloudQueryBaseURL string
	githubBaseURL     string
	httpClient        *http.Client
}

const (
	CloudQueryOrg     = "cloudquery"
	GithubBaseURL     = "https://github.com"
	CloudQueryBaseURL = "https://versions.cloudquery.io"
)

// NewClient returns a new client for fetching plugin versions.
func NewClient() *Client {
	return &Client{
		cloudQueryBaseURL: CloudQueryBaseURL,
		githubBaseURL:     GithubBaseURL,
		httpClient:        http.DefaultClient,
	}
}

// GetLatestCLIRelease returns the latest release version string for CloudQuery CLI
func (c *Client) GetLatestCLIRelease(ctx context.Context) (string, error) {
	return c.readCLIManifest(ctx)
}

// GetLatestPluginRelease returns the latest release version string for the given organization, plugin type
// and plugin.
func (c *Client) GetLatestPluginRelease(ctx context.Context, org, pluginType, pluginName string) (string, error) {
	if org == CloudQueryOrg {
		return c.readManifest(ctx, pluginName)
	}
	return c.readGithubLatest(ctx, org, pluginType, pluginName)
}

func (c *Client) readCLIManifest(ctx context.Context) (string, error) {
	url := fmt.Sprintf(c.cloudQueryBaseURL + "/v1/cli.json")
	b, err := c.doRequest(ctx, url)
	if err != nil {
		return "", fmt.Errorf("reading manifest for cli: %w", err)
	}
	mr := &manifestResponse{}
	err = json.Unmarshal(b, mr)
	if err != nil {
		return "", fmt.Errorf("unmarshaling manifest response: %w", err)
	}
	return extractVersionFromTag(mr.Latest), nil
}

func (c *Client) readManifest(ctx context.Context, name string) (string, error) {
	url := fmt.Sprintf(c.cloudQueryBaseURL+"/v1/%s-%s.json", "source", name)
	b, err := c.doRequest(ctx, url)
	if err != nil {
		return "", fmt.Errorf("reading manifest for %v: %w", name, err)
	}
	mr := &manifestResponse{}
	err = json.Unmarshal(b, mr)
	if err != nil {
		return "", fmt.Errorf("unmarshaling manifest response: %w", err)
	}
	return extractVersionFromTag(mr.Latest), nil
}

// extractVersionFromTag takes a tag of the form "plugins/source/test/v0.1.21" and returns
// the version, i.e. "v0.1.21"
func extractVersionFromTag(tag string) string {
	parts := strings.Split(tag, "/")
	return parts[len(parts)-1]
}

func (c *Client) readGithubLatest(ctx context.Context, org, pluginType, name string) (string, error) {
	url := fmt.Sprintf(c.githubBaseURL+"/%s/cq-%s-%s/releases/latest", org, pluginType, name)
	b, err := c.doRequest(ctx, url)
	if err != nil {
		return "", fmt.Errorf("reading %v: %w", url, err)
	}
	gr := &githubLatestResponse{}
	err = json.Unmarshal(b, gr)
	if err != nil {
		return "", fmt.Errorf("unmarshaling GitHub latest response %s: %w", url, err)
	}
	return gr.TagName, nil
}

func (c *Client) doRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %v (%v)", resp.StatusCode, resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	return b, nil
}
