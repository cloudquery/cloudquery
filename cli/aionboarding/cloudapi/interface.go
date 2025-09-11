package cloudapi

import "context"

type CloudAPI interface {
	ListAllPlugins(ctx context.Context) ([]string, error)
	GetPluginData(ctx context.Context, pluginName, pluginKind string) (string, string, error)
	GetSourcePluginTables(ctx context.Context, pluginName, version, tablesRegex string) ([]string, error)

	GetConversationID(ctx context.Context, userID string) (string, error)
	SetConversationID(ctx context.Context, userID string, conversationID string) error
	GetTokenUsage(ctx context.Context, teamName string) (int, error)
	SetTokenUsage(ctx context.Context, teamName string, tokenUsage int) error
}
