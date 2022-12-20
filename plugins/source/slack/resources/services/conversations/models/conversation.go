package models

import "github.com/slack-go/slack"

// Conversation is a custom model because the SDK-provided struct is wrapped in several levels of embeds
type Conversation struct {
	ID                 string         `json:"id"`
	Created            slack.JSONTime `json:"created"`
	IsOpen             bool           `json:"is_open"`
	LastRead           string         `json:"last_read,omitempty"`
	UnreadCount        int            `json:"unread_count,omitempty"`
	UnreadCountDisplay int            `json:"unread_count_display,omitempty"`
	IsGroup            bool           `json:"is_group"`
	IsShared           bool           `json:"is_shared"`
	IsIM               bool           `json:"is_im"`
	IsExtShared        bool           `json:"is_ext_shared"`
	IsOrgShared        bool           `json:"is_org_shared"`
	IsPendingExtShared bool           `json:"is_pending_ext_shared"`
	IsPrivate          bool           `json:"is_private"`
	IsMpIM             bool           `json:"is_mpim"`
	Unlinked           int            `json:"unlinked"`
	NameNormalized     string         `json:"name_normalized"`
	NumMembers         int            `json:"num_members"`
	Priority           float64        `json:"priority"`
	User               string         `json:"user"`
	Name               string         `json:"name"`
	Creator            string         `json:"creator"`
	IsArchived         bool           `json:"is_archived"`
	Members            []string       `json:"members"`
	Topic              slack.Topic    `json:"topic"`
	Purpose            slack.Purpose  `json:"purpose"`
	IsChannel          bool           `json:"is_channel"`
	IsGeneral          bool           `json:"is_general"`
	IsMember           bool           `json:"is_member"`
	Locale             string         `json:"locale"`
	// Latest             *slack.Message `json:"latest,omitempty"`
	// TODO support pending_shared
	// TODO support previous_names
}

type ConversationMember struct {
	UserID    string `json:"user_id"`
	ChannelID string `json:"channel_id"`
}
