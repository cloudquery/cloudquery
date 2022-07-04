package ses

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type Template struct {
	CreatedTimestamp *time.Time
	*types.Template
}
