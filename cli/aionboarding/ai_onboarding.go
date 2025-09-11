package aionboarding

import (
	"context"
	"errors"
	"os"

	"github.com/cloudquery/cloudquery/cli/v6/aionboarding/cloudapi"
	"github.com/cloudquery/cloudquery/cli/v6/aionboarding/openaiapi"
	"github.com/cloudquery/cloudquery/cli/v6/aionboarding/posthogapi"
	"github.com/rs/zerolog"
	"github.com/samber/lo"

	_ "embed"
)

//go:embed system_prompt.txt
var systemPrompt string

var DefaultPerTeamDayTokenLimit = 10000

type AIOnboarding struct {
	OpenAIAPIKey         string
	PosthogAPIKey        string
	PerTeamDayTokenLimit int
	ForceNewConversation bool
	Debug                bool
	Tools                map[string]openaiapi.Tool

	PosthogAPI posthogapi.PosthogAPI
	CloudAPI   cloudapi.CloudAPI
	OpenAIAPI  openaiapi.OpenAIAPI
}

type Option func(*AIOnboarding)

func WithPerTeamDayTokenLimit(limit int) Option {
	return func(a *AIOnboarding) {
		a.PerTeamDayTokenLimit = limit
	}
}

func WithPosthogAPIKey(apiKey string) Option {
	return func(a *AIOnboarding) {
		a.PosthogAPIKey = apiKey
	}
}

func WithDebug(debug bool) Option {
	return func(a *AIOnboarding) {
		a.Debug = debug
	}
}

func WithOpenAIAPIKey(apiKey string) Option {
	return func(a *AIOnboarding) {
		a.OpenAIAPIKey = apiKey
	}
}

func WithForceNewConversation(forceNewConversation bool) Option {
	return func(a *AIOnboarding) {
		a.ForceNewConversation = forceNewConversation
	}
}

func NewAIOnboarding(ctx context.Context, opts ...Option) (*AIOnboarding, error) {
	a := &AIOnboarding{
		PerTeamDayTokenLimit: DefaultPerTeamDayTokenLimit,
		ForceNewConversation: false,
	}

	for _, opt := range opts {
		opt(a)
	}

	if err := a.validate(); err != nil {
		return nil, err
	}

	a.Tools = a.buildTools()

	a.OpenAIAPI = openaiapi.NewOpenAIAPI(a.OpenAIAPIKey, a.Debug, openaiapi.WithTools(a.Tools))
	a.PosthogAPI = posthogapi.NewPosthogAPI(a.PosthogAPIKey)
	a.CloudAPI = cloudapi.NewCloudAPI(zerolog.New(os.Stdout).Level(zerolog.InfoLevel))

	return a, nil
}

func (a *AIOnboarding) validate() error {
	if a.OpenAIAPIKey == "" {
		return ErrorOpenAIAPIKeyRequired
	}
	if a.PosthogAPIKey == "" {
		return ErrorPosthogAPIKeyRequired
	}
	return nil
}

type ChatResponse struct {
	Message               string
	FunctionCall          string
	FunctionCallArguments map[string]any
	FunctionCallID        string
}

func (c *ChatResponse) HasFunctionCalls() bool {
	return c.FunctionCall != ""
}

func WithInstructions(instructions string) func(*openaiapi.ChatOptions) {
	return func(o *openaiapi.ChatOptions) {
		o.Instructions = instructions
	}
}

func (a *AIOnboarding) Chat(ctx context.Context, userID, teamName string, userMessage *string, opts ...func(*openaiapi.ChatOptions)) (ChatResponse, error) {
	if userID == "" {
		return ChatResponse{}, ErrorUserIDRequired
	}
	if teamName == "" {
		return ChatResponse{}, ErrorTeamNameRequired
	}
	conversationID, err := a.resolveConversation(ctx, userID, opts...)
	if err != nil {
		return ChatResponse{}, err
	}

	opts = append(opts, openaiapi.WithConversationID(conversationID))
	opts = append(opts, WithInstructions(systemPrompt))

	if err := a.validateTokenUsage(ctx, teamName); err != nil {
		return ChatResponse{}, err
	}

	var doChat func(ctx context.Context, userMessage *string, opts ...func(*openaiapi.ChatOptions)) (ChatResponse, error)
	doChat = func(ctx context.Context, userMessage *string, opts ...func(*openaiapi.ChatOptions)) (ChatResponse, error) {
		responses, err := a.OpenAIAPI.Chat(
			ctx,
			userMessage,
			opts...,
		)
		if err != nil {
			return ChatResponse{}, err
		}
		fnResponses := functionCallResponses(responses)
		for _, fnResponse := range fnResponses {
			tool, ok := a.Tools[fnResponse.FunctionCall]
			if !ok {
				return ChatResponse{}, errors.New("tool not found")
			}
			if tool.IsExternal() {
				return ChatResponse{
					Message:               "",
					FunctionCall:          fnResponse.FunctionCall,
					FunctionCallArguments: fnResponse.FunctionCallArguments,
					FunctionCallID:        fnResponse.FunctionCallID,
				}, nil
			}
			output, err := tool.Fn(ctx, fnResponse.FunctionCallArguments)
			if err != nil {
				return ChatResponse{}, err
			}
			opts = append(opts, openaiapi.WithFunctionCallOutput(fnResponse.FunctionCall, fnResponse.FunctionCallArguments, fnResponse.FunctionCallID, output))
		}
		if len(fnResponses) > 0 {
			return doChat(ctx, userMessage, opts...)
		}
		messageResponses := messageResponses(responses)
		if len(messageResponses) > 0 {
			return ChatResponse{
				Message: messageResponses[0].Message,
			}, nil
		}

		return ChatResponse{}, errors.New("didn't get any response from LLM")
	}

	return doChat(ctx, userMessage, opts...)
}

func functionCallResponses(responses []openaiapi.ChatResponse) []openaiapi.ChatResponse {
	return lo.Filter(responses, func(response openaiapi.ChatResponse, _ int) bool {
		return response.IsFunctionCall()
	})
}

func messageResponses(responses []openaiapi.ChatResponse) []openaiapi.ChatResponse {
	return lo.Filter(responses, func(response openaiapi.ChatResponse, _ int) bool {
		return !response.IsFunctionCall()
	})
}

func (a *AIOnboarding) validateTokenUsage(ctx context.Context, teamName string) error {
	tokenUsage, err := a.CloudAPI.GetTokenUsage(ctx, teamName)
	if err != nil {
		return err
	}
	if tokenUsage >= a.PerTeamDayTokenLimit {
		return ErrorTokenUsageLimitReached
	}
	return nil
}

func (a *AIOnboarding) resolveConversation(ctx context.Context, userID string, opts ...func(*openaiapi.ChatOptions)) (string, error) {
	chatOptions := &openaiapi.ChatOptions{}
	for _, opt := range opts {
		opt(chatOptions)
	}

	if chatOptions.ForceNewConversation {
		if chatOptions.ConversationID != "" {
			if err := a.OpenAIAPI.DeleteConversation(ctx, chatOptions.ConversationID); err != nil {
				return "", err
			}
		}
		if err := a.CloudAPI.SetConversationID(ctx, userID, ""); err != nil {
			return "", err
		}
	}
	conversationID, err := a.CloudAPI.GetConversationID(ctx, userID)
	if err != nil {
		return "", err
	}
	if conversationID == "" {
		conversation, err := a.OpenAIAPI.NewConversation(ctx)
		if err != nil {
			return "", err
		}
		conversationID = conversation.ID
	}
	if err := a.CloudAPI.SetConversationID(ctx, userID, conversationID); err != nil {
		return "", err
	}
	return conversationID, nil
}

var (
	ErrorTokenUsageLimitReached = errors.New("Token usage limit reached")
	ErrorOpenAIAPIKeyRequired   = errors.New("OpenAI API key is required")
	ErrorPosthogAPIKeyRequired  = errors.New("Posthog API key is required")
	ErrorUserIDRequired         = errors.New("User ID is required")
	ErrorTeamNameRequired       = errors.New("Team name is required")
)
