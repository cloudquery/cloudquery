package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cloudquery_platform_api "github.com/cloudquery/cloudquery-platform-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
)

const (
	defaultAPIURL = "https://api.cloudquery.io"
	envAPIURL     = "CLOUDQUERY_API_URL"
)

var (
	ErrDisabled = errors.New("AI onboarding is disabled")
)

func NewClient(token string, opts ...cloudquery_platform_api.ClientOption) (*cloudquery_platform_api.ClientWithResponses, error) {
	requestEditorOpt := cloudquery_platform_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	})

	clientOpts := append([]cloudquery_platform_api.ClientOption{requestEditorOpt}, opts...)

	c, err := cloudquery_platform_api.NewClientWithResponses(
		env.GetEnvOrDefault(envAPIURL, defaultAPIURL),
		clientOpts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return c, nil
}

func NewAnonymousClient() (*cloudquery_platform_api.ClientWithResponses, error) {
	c, err := cloudquery_platform_api.NewClientWithResponses(env.GetEnvOrDefault(envAPIURL, defaultAPIURL))
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return c, nil
}

func ListAllPlugins(cl *cloudquery_platform_api.ClientWithResponses) ([]cloudquery_platform_api.ListPlugin, error) {
	page := cloudquery_platform_api.Page(1)
	perPage := cloudquery_platform_api.PerPage(100)
	plugins := make([]cloudquery_platform_api.ListPlugin, 0)
	for {
		resp, err := cl.ListPluginsWithResponse(context.Background(), &cloudquery_platform_api.ListPluginsParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
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

func GetPluginVersion(cl *cloudquery_platform_api.ClientWithResponses, teamName string, kind cloudquery_platform_api.PluginKind, pluginName, pluginVersion string) (*cloudquery_platform_api.PluginVersionDetails, error) {
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

type FunctionCallOutput struct {
	// Arguments The arguments passed to the function
	Arguments map[string]any `json:"arguments"`

	// CallID The unique identifier for this function call
	CallID string `json:"call_id"`

	// Name The name of the function that was called
	Name string `json:"name"`

	// Output The output/result from the function call
	Output string `json:"output"`
}

type ChatResponse struct {
	Message               *string
	FunctionCall          *string
	FunctionCallID        string
	FunctionCallArguments map[string]any
}

func Chat(ctx context.Context, cl *cloudquery_platform_api.ClientWithResponses, teamName string, message *string, functionCallOutputs *[]FunctionCallOutput) (*ChatResponse, error) {
	// Platform API does not have equivalent AI onboarding chat endpoint
	return nil, fmt.Errorf("AI onboarding chat is not available in platform API")
}

func NewConversation(ctx context.Context, cl *cloudquery_platform_api.ClientWithResponses, teamName string, resumeConversation bool) error {
	// Platform API does not have equivalent AI onboarding conversation endpoint
	return fmt.Errorf("AI onboarding conversation is not available in platform API")
}

func EndConversation(ctx context.Context, cl *cloudquery_platform_api.ClientWithResponses, teamName string) {
	// Platform API does not have equivalent AI onboarding conversation endpoint
	// This call is best-effort, so we don't return an error
}
