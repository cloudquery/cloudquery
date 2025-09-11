package openaiapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/conversations"
	"github.com/openai/openai-go/v2/option"
	"github.com/openai/openai-go/v2/responses"
)

type FunctionCallOutput struct {
	Name      string
	Arguments map[string]any
	CallID    string
	Output    string
}

type ChatOptions struct {
	ForceNewConversation bool
	ConversationID       string
	Model                string
	Tools                []responses.ToolUnionParam
	Instructions         string
	FunctionCallOutput   []FunctionCallOutput
}

type ChatResponse struct {
	Message               string
	FunctionCall          string
	FunctionCallArguments map[string]any
	FunctionCallID        string
}

func (c *ChatResponse) IsMessage() bool {
	return c.FunctionCall == ""
}

func (c *ChatResponse) IsFunctionCall() bool {
	return c.FunctionCall != ""
}

func WithForceNewConversation(forceNewConversation bool) func(*ChatOptions) {
	return func(o *ChatOptions) {
		o.ForceNewConversation = forceNewConversation
	}
}

func WithConversationID(conversationID string) func(*ChatOptions) {
	return func(o *ChatOptions) {
		o.ConversationID = conversationID
	}
}

func WithModel(model string) func(*ChatOptions) {
	return func(o *ChatOptions) {
		o.Model = model
	}
}

func WithTools(tools map[string]Tool) func(*openAIAPI) {
	return func(a *openAIAPI) {
		a.tools = make([]responses.ToolUnionParam, 0, len(tools))
		for _, tool := range tools {
			a.tools = append(a.tools, tool.ToOpenAITool())
		}
	}
}

func WithInstructions(instructions string) func(*ChatOptions) {
	return func(a *ChatOptions) {
		a.Instructions = instructions
	}
}

func WithFunctionCallOutput(name string, arguments map[string]any, functionCallID string, output string) func(*ChatOptions) {
	return func(a *ChatOptions) {
		a.FunctionCallOutput = append(a.FunctionCallOutput, FunctionCallOutput{
			Name:      name,
			Arguments: arguments,
			CallID:    functionCallID,
			Output:    output,
		})
	}
}

type OpenAIAPI interface {
	Chat(ctx context.Context, message *string, opts ...func(*ChatOptions)) ([]ChatResponse, error)
	NewConversation(ctx context.Context) (*conversations.Conversation, error)
	DeleteConversation(ctx context.Context, conversationID string) error
}

type openAIAPI struct {
	client openai.Client
	tools  []responses.ToolUnionParam
}

func NewOpenAIAPI(apiKey string, debug bool, opts ...func(*openAIAPI)) OpenAIAPI {
	os.Setenv("OPENAI_API_KEY", apiKey)

	client := openai.NewClient()
	if debug {
		client = openai.NewClient(option.WithMiddleware(createLoggingMiddleware()))
	}

	a := &openAIAPI{client: client}
	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *openAIAPI) Chat(ctx context.Context, userMessage *string, opts ...func(*ChatOptions)) ([]ChatResponse, error) {
	chatOptions := &ChatOptions{
		Model:        openai.ChatModelGPT4oMini,
		Tools:        []responses.ToolUnionParam{},
		Instructions: "",
	}
	for _, opt := range opts {
		opt(chatOptions)
	}

	inputs := []responses.ResponseInputItemUnionParam{}

	if userMessage != nil {
		inputs = append(inputs, responses.ResponseInputItemUnionParam{OfMessage: &responses.EasyInputMessageParam{
			Content: responses.EasyInputMessageContentUnionParam{OfString: openai.String(*userMessage)},
			Role:    responses.EasyInputMessageRoleUser,
		}})
	}

	for _, functionCallOutput := range chatOptions.FunctionCallOutput {
		arguments, err := json.Marshal(functionCallOutput.Arguments)
		if err != nil {
			return []ChatResponse{}, err
		}
		inputs = append(inputs, responses.ResponseInputItemUnionParam{
			OfFunctionCall: &responses.ResponseFunctionToolCallParam{
				Status:    "completed",
				Name:      functionCallOutput.Name,
				Arguments: string(arguments),
				CallID:    functionCallOutput.CallID,
			},
		})
		inputs = append(inputs, responses.ResponseInputItemUnionParam{
			OfFunctionCallOutput: &responses.ResponseInputItemFunctionCallOutputParam{
				Output: functionCallOutput.Output,
				Status: "completed",
				CallID: functionCallOutput.CallID,
			},
		})
	}

	response, err := a.client.Responses.New(ctx, responses.ResponseNewParams{
		Conversation: responses.ResponseNewParamsConversationUnion{OfString: openai.String(chatOptions.ConversationID)},
		Model:        chatOptions.Model,
		Instructions: openai.String(chatOptions.Instructions),
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: inputs,
		},
		Tools: a.tools,
	})
	if err != nil {
		return []ChatResponse{}, err
	}

	chatResponses := []ChatResponse{}

	for _, outputItem := range response.Output {
		switch outputItem.Type {
		case "function_call":
			functionCall := outputItem.AsFunctionCall()
			arguments := make(map[string]any)
			err := json.Unmarshal([]byte(functionCall.Arguments), &arguments)
			if err != nil {
				return []ChatResponse{}, err
			}
			chatResponse := ChatResponse{
				FunctionCall:          functionCall.Name,
				FunctionCallArguments: arguments,
				FunctionCallID:        functionCall.CallID,
			}
			chatResponses = append(chatResponses, chatResponse)
		case "message":
			for _, content := range outputItem.Content {
				switch variant := content.AsAny().(type) {
				case responses.ResponseOutputText:
					chatResponses = append(chatResponses, ChatResponse{
						Message: variant.Text,
					})
				case responses.ResponseOutputRefusal:
					chatResponses = append(chatResponses, ChatResponse{
						Message: variant.Refusal,
					})
				default:
					return []ChatResponse{}, errors.New("no variant present")
				}
			}
		default:
			fmt.Println("ignoring unsupported output item type", outputItem.Type)
		}
	}

	return chatResponses, nil
}

func (a *openAIAPI) NewConversation(ctx context.Context) (*conversations.Conversation, error) {
	return a.client.Conversations.New(ctx, conversations.ConversationNewParams{})
}

func (a *openAIAPI) DeleteConversation(ctx context.Context, conversationID string) error {
	_, err := a.client.Conversations.Delete(ctx, conversationID)
	return err
}
