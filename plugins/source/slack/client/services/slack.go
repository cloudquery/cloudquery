package services

import (
	"context"
	"io"

	"github.com/slack-go/slack"
)

//go:generate mockgen -package=mocks -destination=../mocks/slack.go -source=slack.go SlackClient
type SlackClient interface {
	GetAccessLogsContext(context.Context, slack.AccessLogParameters) ([]slack.Login, *slack.Paging, error)
	GetAuditLogsContext(context.Context, slack.AuditLogParameters) ([]slack.AuditEntry, string, error)
	GetBillableInfoContext(context.Context, string) (map[string]slack.BillingActive, error)
	GetBillableInfoForTeamContext(context.Context) (map[string]slack.BillingActive, error)
	GetBotInfoContext(context.Context, string) (*slack.Bot, error)
	GetConversationHistoryContext(context.Context, *slack.GetConversationHistoryParameters) (*slack.GetConversationHistoryResponse, error)
	GetConversationInfoContext(context.Context, *slack.GetConversationInfoInput) (*slack.Channel, error)
	GetConversationRepliesContext(context.Context, *slack.GetConversationRepliesParameters) ([]slack.Message, bool, string, error)
	GetConversationsContext(context.Context, *slack.GetConversationsParameters) ([]slack.Channel, string, error)
	GetConversationsForUserContext(context.Context, *slack.GetConversationsForUserParameters) ([]slack.Channel, string, error)
	GetDNDInfoContext(context.Context, *string) (*slack.DNDStatus, error)
	GetDNDTeamInfoContext(context.Context, []string) (map[string]slack.DNDStatus, error)
	GetEmojiContext(context.Context) (map[string]string, error)
	GetFileContext(context.Context, string, io.Writer) error
	GetFileInfoContext(context.Context, string, int, int) (*slack.File, []slack.Comment, *slack.Paging, error)
	GetFilesContext(context.Context, slack.GetFilesParameters) ([]slack.File, *slack.Paging, error)
	GetOtherTeamInfoContext(context.Context, string) (*slack.TeamInfo, error)
	GetPermalinkContext(context.Context, *slack.PermalinkParameters) (string, error)
	GetReactionsContext(context.Context, slack.ItemRef, slack.GetReactionsParameters) ([]slack.ItemReaction, error)
	GetRemoteFileInfoContext(context.Context, string, string) (*slack.RemoteFile, error)
	GetScheduledMessagesContext(context.Context, *slack.GetScheduledMessagesParameters) ([]slack.ScheduledMessage, string, error)
	GetStarredContext(context.Context, slack.StarsParameters) ([]slack.StarredItem, *slack.Paging, error)
	GetTeamInfoContext(context.Context) (*slack.TeamInfo, error)
	GetTeamProfileContext(context.Context) (*slack.TeamProfile, error)
	GetUserByEmailContext(context.Context, string) (*slack.User, error)
	GetUserGroupMembersContext(context.Context, string) ([]string, error)
	GetUserGroupsContext(context.Context, ...slack.GetUserGroupsOption) ([]slack.UserGroup, error)
	GetUserIdentityContext(context.Context) (*slack.UserIdentityResponse, error)
	GetUserInfoContext(context.Context, string) (*slack.User, error)
	GetUserPrefsContext(context.Context) (*slack.UserPrefsCarrier, error)
	GetUserPresenceContext(context.Context, string) (*slack.UserPresence, error)
	GetUserProfileContext(context.Context, *slack.GetUserProfileParameters) (*slack.UserProfile, error)
	GetUsersContext(context.Context, ...slack.GetUsersOption) ([]slack.User, error)
	GetUsersInConversationContext(context.Context, *slack.GetUsersInConversationParameters) ([]string, string, error)
	GetUsersInfoContext(context.Context, ...string) (*[]slack.User, error)
	ListAllStarsContext(context.Context) ([]slack.Item, error)
	ListBookmarksContext(context.Context, string) ([]slack.Bookmark, error)
	ListEventAuthorizationsContext(context.Context, string) ([]slack.EventAuthorization, error)
	ListFilesContext(context.Context, slack.ListFilesParameters) ([]slack.File, *slack.ListFilesParameters, error)
	ListPinsContext(context.Context, string) ([]slack.Item, *slack.Paging, error)
	ListReactionsContext(context.Context, slack.ListReactionsParameters) ([]slack.ReactedItem, *slack.Paging, error)
	ListRemindersContext(context.Context) ([]*slack.Reminder, error)
	ListRemoteFilesContext(context.Context, slack.ListRemoteFilesParameters) ([]slack.RemoteFile, error)
	ListStarsContext(context.Context, slack.StarsParameters) ([]slack.Item, *slack.Paging, error)
	ListTeamsContext(context.Context, slack.ListTeamsParameters) ([]slack.Team, string, error)
}
