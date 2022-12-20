package conversations

import "strings"

// isNotInChannel is an error returned by Slack if the bot token doesn't have
// access to the channel because it's not added to it. We don't consider this
// an error.
func isNotInChannel(err error) bool {
	return strings.Contains(err.Error(), "not_in_channel")
}
