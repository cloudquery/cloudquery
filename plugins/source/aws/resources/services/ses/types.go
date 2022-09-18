package ses

import (
	"time"
)

type Template struct {
	TemplateName *string

	// The HTML body of the email.
	Html *string

	// The subject line of the email.
	Subject *string

	// The email body that will be visible to recipients whose email clients do not
	// display HTML.
	Text *string

	CreatedTimestamp *time.Time
}
