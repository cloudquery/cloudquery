package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
)

const (
	defaultAPIURL = "https://api.cloudquery.io"
	envAPIURL     = "CLOUDQUERY_API_URL"
	envCLIAPIURL  = "CLOUDQUERY_CLI_API_URL"
	envCLIToken   = "CLOUDQUERY_CLI_TOKEN"
)

func NewClient(token string) (*cloudquery_api.ClientWithResponses, error) {
	return newClient(token, false)
}

func NewAnonymousClient() (*cloudquery_api.ClientWithResponses, error) {
	return NewClient("")
}

func NewLocalGroupClient(token string) (*cloudquery_api.ClientWithResponses, error) {
	return newClient(token, true)
}

func ListAllPlugins(cl *cloudquery_api.ClientWithResponses) ([]cloudquery_api.ListPlugin, error) {
	page := cloudquery_api.Page(1)
	perPage := cloudquery_api.PerPage(100)
	plugins := make([]cloudquery_api.ListPlugin, 0)
	for {
		resp, err := cl.ListPluginsWithResponse(context.Background(), &cloudquery_api.ListPluginsParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to list plugins: %w", err)
		}
		if resp.JSON200 == nil {
			return nil, fmt.Errorf("failed to list plugins: %w", err)
		}
		if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
			return nil, fmt.Errorf("failed to list plugins: %s", resp.Status())
		}
		plugins = append(plugins, resp.JSON200.Items...)
		if resp.JSON200.Metadata.LastPage == nil || *resp.JSON200.Metadata.LastPage <= int(page) {
			break
		}
		page++
	}
	return plugins, nil
}

func GetPluginVersion(cl *cloudquery_api.ClientWithResponses, teamName string, kind cloudquery_api.PluginKind, pluginName, pluginVersion string) (*cloudquery_api.PluginVersionDetails, error) {
	resp, err := cl.GetPluginVersionWithResponse(context.Background(), teamName, kind, pluginName, pluginVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to get plugin version %s/%s@%s: %w", teamName, pluginName, pluginVersion, err)
	}
	if resp.JSON200 == nil {
		return nil, fmt.Errorf("failed to get plugin version %s/%s@%s: %w", teamName, pluginName, pluginVersion, err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get plugin version %s/%s@%s: %s", teamName, pluginName, pluginVersion, resp.Status())
	}
	return resp.JSON200, nil
}

func getAPIURL(preferLocalGroup bool) (apiURL string, isLocalGroup bool) {
	regularAPI := env.GetEnvOrDefault(envAPIURL, defaultAPIURL)
	if !preferLocalGroup {
		return regularAPI, false
	}

	isCloud, _ := strconv.ParseBool(os.Getenv("CQ_CLOUD"))
	if !isCloud {
		return regularAPI, false
	}

	val := env.GetEnvOrDefault(envCLIAPIURL, regularAPI)
	return val, val != regularAPI
}

func overrideToken(token string, getLocalGroup bool) string {
	if !getLocalGroup {
		return token
	}
	return env.GetEnvOrDefault(envCLIToken, "")
}

func newClient(token string, localGroup bool) (*cloudquery_api.ClientWithResponses, error) {
	endpoint, isLocalGroup := getAPIURL(localGroup)
	token = overrideToken(token, isLocalGroup)

	var opts []cloudquery_api.ClientOption
	if token != "" {
		opts = append(opts, cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	}

	c, err := cloudquery_api.NewClientWithResponses(endpoint, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return c, nil
}
