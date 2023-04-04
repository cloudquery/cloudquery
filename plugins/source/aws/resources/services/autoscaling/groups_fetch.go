package autoscaling

import (
	"errors"
	"regexp"

	"github.com/aws/smithy-go"
)

var groupNotFoundRegex = regexp.MustCompile(`AutoScalingGroup name not found|Group .* not found`)

func isAutoScalingGroupNotExistsError(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "ValidationError" && groupNotFoundRegex.MatchString(ae.ErrorMessage()) {
			return true
		}
	}
	return false
}
