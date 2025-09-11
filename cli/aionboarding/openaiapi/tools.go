package openaiapi

import (
	"context"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/responses"
)

type Tool struct {
	Name        string
	Description string
	Parameters  map[string]any

	// Fn can be nil if the tool is external to the API and must be run by the client
	Fn func(ctx context.Context, args map[string]any) (string, error)
}

func (t Tool) IsExternal() bool {
	return t.Fn == nil
}

func (t Tool) ToOpenAITool() responses.ToolUnionParam {
	return responses.ToolUnionParam{OfFunction: &responses.FunctionToolParam{
		Name:        t.Name,
		Description: openai.String(t.Description),
		Parameters:  t.Parameters,
	}}
}
