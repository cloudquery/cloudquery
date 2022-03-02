package client

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/smithy-go"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	client := meta.(*Client)

	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation", "AuthorizationError", "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return diag.Diagnostics{
				RedactError(client.Accounts, diag.NewBaseError(err, diag.ACCESS, diag.WithType(diag.ACCESS), ParseSummaryMessage(err, ae),
					diag.WithDetails("%s", errorCodeDescriptions[ae.ErrorCode()]), diag.WithNoOverwrite(), diag.WithSeverity(diag.WARNING))),
			}
		case "InvalidAction":
			return diag.Diagnostics{
				RedactError(client.Accounts, diag.NewBaseError(err, diag.RESOLVING, diag.WithType(diag.RESOLVING), diag.WithSeverity(diag.IGNORE), ParseSummaryMessage(err, ae),
					diag.WithDetails("The action is invalid for the service."))),
			}
		}
	}
	if IsErrorThrottle(err) {
		return diag.Diagnostics{
			RedactError(client.Accounts, diag.NewBaseError(err, diag.THROTTLE, diag.WithType(diag.THROTTLE), diag.WithSeverity(diag.WARNING), ParseSummaryMessage(err, ae),
				diag.WithDetails("CloudQuery AWS provider has been throttled, increase max_retries/retry_timeout in provider configuration."))),
		}
	}

	// Take over from SDK and always return diagnostics, redacting PII
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(client.Accounts, d),
		}
	}

	return diag.Diagnostics{
		RedactError(client.Accounts, diag.NewBaseError(err, diag.RESOLVING, diag.WithResourceName(resourceName))),
	}
}

func ParseSummaryMessage(err error, apiErr smithy.APIError) diag.BaseErrorOption {
	for {
		if op, ok := err.(*smithy.OperationError); ok {
			return diag.WithError(fmt.Errorf("%s: %s - %s", op.Service(), op.Operation(), apiErr.ErrorMessage()))
		}
		if err = errors.Unwrap(err); err == nil {
			return diag.WithError(errors.New(apiErr.ErrorMessage()))
		}
	}
}

// RedactError redacts a given diagnostic and returns a RedactedDiagnostic containing both original and redacted versions
func RedactError(aa []Account, e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		errors.New(removePII(aa, e.Error())),
		e.Type(),
		diag.WithSeverity(e.Severity()),
		diag.WithResourceName(e.Description().Resource),
		diag.WithSummary("%s", removePII(aa, e.Description().Summary)),
		diag.WithDetails("%s", removePII(aa, e.Description().Detail)),
	)
	return diag.NewRedactedDiagnostic(e, r)
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

var (
	requestIdRegex = regexp.MustCompile(`\sRequestID: [A-Za-z0-9-]+`)
	hostIdRegex    = regexp.MustCompile(`\sHostID: [A-Za-z0-9+/_=-]+`)
	arnIdRegex     = regexp.MustCompile(`\sarn:aws[A-Za-z0-9-]*:.+?\s`)
)

func removePII(aa []Account, msg string) string {
	for i := range aa {
		msg = strings.ReplaceAll(msg, " AccountID "+aa[i].ID, " AccountID xxxx")
	}
	msg = requestIdRegex.ReplaceAllString(msg, " RequestID: xxxx")
	msg = hostIdRegex.ReplaceAllString(msg, " HostID: xxxx")
	msg = arnIdRegex.ReplaceAllString(msg, " arn:xxxx ")
	msg = accountObfusactor(aa, msg)

	return msg
}
