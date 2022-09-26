package plugins

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

const (
	CloudQueryOrg     = "cloudquery"
	GithubBaseURL     = "https://github.com"
	CloudQueryBaseURL = "https://versions.cloudquery.io"
)


// GetLatestPluginRelease returns the latest release version string for the given organization, plugin type
// and plugin.
func GetLatestPluginRelease(ctx context.Context, org, name string, typ PluginType) (string, error) {
	if org == CloudQueryOrg {
		return getLatestCQPluginRelease(ctx, name, typ)
	}
	return getLatestCommunityPluginRelease(ctx, org, name, typ)
}

func GetLatestCLIRelease(ctx context.Context) (string, error) {
	url := fmt.Sprintf(CloudQueryBaseURL + "/v2/cli.json")
	b, err := doRequest(ctx, url)
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

func getLatestCQPluginRelease(ctx context.Context, name string, typ PluginType) (string, error) {
	url := fmt.Sprintf(CloudQueryBaseURL+"/v2/%s-%s.json", typ, name)
	b, err := doRequest(ctx, url)
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

// extractVersionFromTag takes a tag of the form "plugins-source-test-v0.1.21" and returns
// the version, i.e. "v0.1.21"
func extractVersionFromTag(tag string) string {
	parts := strings.Split(tag, "-")
	return parts[len(parts)-1]
}

func getLatestCommunityPluginRelease(ctx context.Context, org, name string, typ PluginType) (string, error) {
	url := fmt.Sprintf(GithubBaseURL+"/%s/cq-%s-%s/releases/latest", org, typ, name)
	b, err := doRequest(ctx, url)
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

func doRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
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
