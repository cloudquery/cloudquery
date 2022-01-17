package client

import (
	"errors"
	"fmt"

	"github.com/aws/smithy-go"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/provider/schema/diag"
)

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) []diag.Diagnostic {
	client := meta.(*Client)
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation", "AuthorizationError":
			return []diag.Diagnostic{
				diag.FromError(err, diag.WARNING, diag.ACCESS, resourceName, ParseSummaryMessage(client.Accounts, err, ae), errorCodeDescriptions[ae.ErrorCode()]),
			}
		case "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return []diag.Diagnostic{
				diag.FromError(err, diag.WARNING, diag.ACCESS, resourceName, ParseSummaryMessage(client.Accounts, err, ae), errorCodeDescriptions[ae.ErrorCode()]),
			}
		case "InvalidAction":
			return []diag.Diagnostic{
				diag.FromError(err, diag.IGNORE, diag.RESOLVING, resourceName, ParseSummaryMessage(client.Accounts, err, ae),
					"The action is invalid for the service."),
			}
		}
	}
	if IsErrorThrottle(err) {
		return []diag.Diagnostic{
			diag.FromError(err, diag.WARNING, diag.THROTTLE, resourceName, ParseSummaryMessage(client.Accounts, err, ae),
				"CloudQuery AWS provider has been throttled, increase max_retries/retry_timeout in provider configuration."),
		}
	}

	return nil
}

func ParseSummaryMessage(aa []Account, err error, apiErr smithy.APIError) string {
	for {
		if op, ok := err.(*smithy.OperationError); ok {
			return fmt.Sprintf("%s: %s - %s", op.Service(), op.Operation(), accountObfusactor(aa, apiErr.ErrorMessage()))
		}
		if err = errors.Unwrap(err); err == nil {
			return accountObfusactor(aa, apiErr.ErrorMessage())
		}
	}
}

// IsErrorThrottle returns whether the error is to be throttled based on its code.
// Returns false if error is nil.
func IsErrorThrottle(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		return isCodeThrottle(ae.ErrorCode())
	}
	return false
}

var errorCodeDescriptions = map[string]string{
	"InvalidClientTokenId":          "The X.509 certificate or AWS access key ID provided does not exist in our records.",
	"SubscriptionRequiredException": "When you created your AWS account, all available services at that time were activated. However, as new services are released, they aren't automatically put into an active state without your permission. You must subscribe to each service individually as they are released.",
	"OptInRequired":                 "You are not authorized to use the requested service. Ensure that you have subscribed to the service you are trying to use. If you are new to AWS, your account might take some time to be activated while your credit card details are being verified.",
	"UnauthorizedOperation":         "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
	"AccessDeniedException":         "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
	"AccessDenied":                  "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
	"AuthorizationError":            "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
}

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

func isCodeThrottle(code string) bool {
	_, ok := throttleCodes[code]
	return ok
}
