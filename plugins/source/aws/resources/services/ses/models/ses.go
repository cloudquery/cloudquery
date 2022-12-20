package models

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type EmailIdentity struct {
	IdentityName   *string
	SendingEnabled bool

	*sesv2.GetEmailIdentityOutput
}

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
