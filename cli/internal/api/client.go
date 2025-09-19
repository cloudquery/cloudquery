package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
)

const (
	defaultAPIURL = "https://api.cloudquery.io"
	envAPIURL     = "CLOUDQUERY_API_URL"
)

var (
	ErrDisabled = errors.New("AI onboarding is disabled")
)

func NewClient(token string) (*cloudquery_api.ClientWithResponses, error) {
	c, err := cloudquery_api.NewClientWithResponses(env.GetEnvOrDefault(envAPIURL, defaultAPIURL),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return c, nil
}

func NewAnonymousClient() (*cloudquery_api.ClientWithResponses, error) {
	c, err := cloudquery_api.NewClientWithResponses(env.GetEnvOrDefault(envAPIURL, defaultAPIURL))
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return c, nil
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

func Chat(ctx context.Context, cl *cloudquery_api.ClientWithResponses, teamName string, message *string, functionCallOutputs *[]FunctionCallOutput) (*ChatResponse, error) {
	imsg := any(message)
	ifcos := any(functionCallOutputs)

	requestBody := cloudquery_api.AIOnboardingChatJSONRequestBody{
		Message:             &imsg,
		FunctionCallOutputs: &ifcos,
	}

	resp, err := cl.AIOnboardingChatWithResponse(ctx, teamName, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to chat: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to chat with code %d: %s", resp.StatusCode(), resp.Status())
	}
	if resp.JSON200 == nil {
		return nil, fmt.Errorf("failed to chat with code %d: no response data", resp.StatusCode())
	}

	// Extract values from the interface{} fields
	chatResp := &ChatResponse{}

	// Extract message
	if resp.JSON200.Message != nil {
		if msg, ok := resp.JSON200.Message.(string); ok {
			chatResp.Message = &msg
		}
	}

	// Extract function call
	if resp.JSON200.FunctionCall != nil {
		if fc, ok := (*resp.JSON200.FunctionCall).(string); ok {
			chatResp.FunctionCall = &fc
		}
	}

	// Extract function call ID
	if resp.JSON200.FunctionCallID != nil {
		if fcid, ok := (*resp.JSON200.FunctionCallID).(string); ok {
			chatResp.FunctionCallID = fcid
		}
	}

	// Extract function call arguments
	if resp.JSON200.FunctionCallArguments != nil {
		if fca, ok := (*resp.JSON200.FunctionCallArguments).(map[string]any); ok {
			chatResp.FunctionCallArguments = fca
		}
	}

	return chatResp, nil
}

func NewConversation(ctx context.Context, cl *cloudquery_api.ClientWithResponses, teamName string, resumeConversation bool) error {
	requestBody := cloudquery_api.AIOnboardingNewConversationJSONRequestBody{}
	if resumeConversation {
		tryResume := any(true)
		requestBody.TryResume = &tryResume
	}
	resp, err := cl.AIOnboardingNewConversationWithResponse(ctx, teamName, requestBody)
	if err != nil {
		return fmt.Errorf("failed to start new conversation: %w", err)
	}
	if resp.StatusCode() == http.StatusNotFound {
		return ErrDisabled
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to start new conversation: %s", resp.Status())
	}
	return nil
}

func EndConversation(ctx context.Context, cl *cloudquery_api.ClientWithResponses, teamName string) {
	_, _ = cl.AIOnboardingEndConversationWithResponse(ctx, teamName)
	// This call is best-effort, so we don't return an error
}
