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

func NewClient(token string, opts ...cloudquery_api.ClientOption) (*cloudquery_api.ClientWithResponses, error) {
	requestEditorOpt := cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	})

	clientOpts := append([]cloudquery_api.ClientOption{requestEditorOpt}, opts...)

	c, err := cloudquery_api.NewClientWithResponses(
		env.GetEnvOrDefault(envAPIURL, defaultAPIURL),
		clientOpts...,
	)
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

type FunctionCallOutput = cloudquery_api.FunctionCallOutput

type ChatResponse struct {
	Message               *string
	FunctionCall          *string
	FunctionCallID        string
	FunctionCallArguments map[string]any
}

func Chat(ctx context.Context, cl *cloudquery_api.ClientWithResponses, teamName string, message *string, functionCallOutputs *[]FunctionCallOutput) (*ChatResponse, error) {
	resp, err := cl.AIOnboardingChatWithResponse(ctx, teamName, cloudquery_api.AIOnboardingChatJSONRequestBody{
		Message:             message,
		FunctionCallOutputs: functionCallOutputs,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to chat: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to chat with code %d: %s", resp.StatusCode(), resp.Status())
	}
	if resp.JSON200 == nil {
		return nil, fmt.Errorf("failed to chat with code %d: no response data", resp.StatusCode())
	}

	chatResp := &ChatResponse{}
	if resp.JSON200.Message != "" {
		m := resp.JSON200.Message
		chatResp.Message = &m
	}
	if resp.JSON200.FunctionCall != nil {
		fc := *resp.JSON200.FunctionCall
		chatResp.FunctionCall = &fc
	}
	if resp.JSON200.FunctionCallID != nil {
		chatResp.FunctionCallID = *resp.JSON200.FunctionCallID
	}
	if resp.JSON200.FunctionCallArguments != nil {
		chatResp.FunctionCallArguments = *resp.JSON200.FunctionCallArguments
	}
	return chatResp, nil
}

func NewConversation(ctx context.Context, cl *cloudquery_api.ClientWithResponses, teamName string, resumeConversation bool) error {
	requestBody := cloudquery_api.AIOnboardingNewConversationJSONRequestBody{}
	if resumeConversation {
		tryResume := true
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
