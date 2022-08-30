package autoscaling

import (
	"errors"
	"regexp"

	"github.com/aws/smithy-go"
)

var autoscalingGroupNotFoundRegex = regexp.MustCompile(`AutoScalingGroup name not found|Group .* not found`)

func IsGroupNotExistsError(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "ValidationError" && autoscalingGroupNotFoundRegex.MatchString(ae.ErrorMessage()) {
			return true
		}
	}
	return false
}
