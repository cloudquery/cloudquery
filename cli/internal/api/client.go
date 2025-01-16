package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
)

const (
	defaultAPIURL = "https://api.cloudquery.io"
	envAPIURL     = "CLOUDQUERY_API_URL"
	envCLIAPIURL  = "CLOUDQUERY_CLI_API_URL"
	envCLIKey     = "CLOUDQUERY_CLI_KEY"
)

// NewClient creates a new client with the given token.
func NewClient(token string) (*cloudquery_api.ClientWithResponses, error) {
	return newClient(env.GetEnvOrDefault(envAPIURL, defaultAPIURL), token)
}

// NewAnonymousClient creates a client that doesn't require authentication.
func NewAnonymousClient() (*cloudquery_api.ClientWithResponses, error) {
	return NewClient("")
}

// NewLocalClient creates a client that connects to the local API if possible. If not, it falls back to the regular API using AuthClient.
func NewLocalClient(acceptableKeyTypes ...auth.TokenType) (*cloudquery_api.ClientWithResponses, error) {
	var (
		tokenType  auth.TokenType
		tokenValue string
	)

	apiURL, isLocal := getAPIURL(true)
	if !isLocal {
		token, err := auth.NewTokenClient().GetToken()
		if err != nil {
			return nil, err
		}
		tokenValue, tokenType = token.Value, token.Type
	} else {
		tokenValue = env.GetEnvOrDefault(envCLIKey, "")
		tokenType = tokenTypeFromValue(tokenValue)
	}

	if len(acceptableKeyTypes) > 0 {
		if !slices.Contains(acceptableKeyTypes, tokenType) {
			return nil, nil
		}
	}

	return newClient(apiURL, tokenValue)
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

func getAPIURL(preferLocal bool) (apiURL string, isLocal bool) {
	regularAPI := env.GetEnvOrDefault(envAPIURL, defaultAPIURL)
	if !preferLocal {
		return regularAPI, false
	}

	isCloud, _ := strconv.ParseBool(os.Getenv("CQ_CLOUD"))
	if !isCloud {
		return regularAPI, false
	}

	val := env.GetEnvOrDefault(envCLIAPIURL, regularAPI)
	return val, val != regularAPI
}

func newClient(endpoint, token string) (*cloudquery_api.ClientWithResponses, error) {
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

func tokenTypeFromValue(token string) auth.TokenType {
	switch {
	case strings.HasPrefix(token, "cqsr_"):
		return auth.SyncRunAPIKey
	case strings.HasPrefix(token, "cqstc_"):
		return auth.SyncTestConnectionAPIKey
	case token != "":
		return auth.APIKey
	default:
		return auth.Undefined
	}
}
