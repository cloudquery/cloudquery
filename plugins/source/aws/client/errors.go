package client

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws/ratelimit"
	"github.com/aws/smithy-go"
)

var throttleCodes = map[string]struct{}{
	"ProvisionedThroughputExceededException": {},
	"Throttling":                             {},
	"ThrottlingException":                    {},
	"RequestLimitExceeded":                   {},
	"RequestThrottled":                       {},
	"RequestThrottledException":              {},
	"TooManyRequestsException":               {}, // Lambda functions
	"PriorRequestNotComplete":                {}, // Route53
}

// IsErrorThrottle returns whether the error is to be throttled based on its code.
// Returns false if error is nil.
func IsErrorThrottle(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) && isCodeThrottle(ae.ErrorCode()) {
		return true
	}
	var qe ratelimit.QuotaExceededError
	return errors.As(err, &qe)
}

func isCodeThrottle(code string) bool {
	_, ok := throttleCodes[code]
	return ok
}
