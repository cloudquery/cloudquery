package ses

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type Template struct {
	TemplateName *string
	*types.EmailTemplateContent
	CreatedTimestamp *time.Time
}
