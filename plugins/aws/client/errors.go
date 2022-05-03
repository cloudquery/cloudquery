package client

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/ratelimit"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	client := meta.(*Client)

	return classifyError(err, diag.RESOLVING, client.Accounts, diag.WithResourceName(resourceName), includeResourceIdWithAccount(client, err))
}

func classifyError(err error, fallbackType diag.Type, accounts []Account, opts ...diag.BaseErrorOption) diag.Diagnostics {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation", "AuthorizationError", "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return diag.Diagnostics{
				RedactError(accounts, diag.NewBaseError(err,
					diag.ACCESS,
					append(opts,
						diag.WithType(diag.ACCESS),
						diag.WithSeverity(diag.WARNING),
						ParseSummaryMessage(err),
						diag.WithDetails("%s", errorCodeDescriptions[ae.ErrorCode()]),
					)...),
				),
			}
		case "InvalidAction":
			return diag.Diagnostics{
				RedactError(accounts, diag.NewBaseError(err,
					diag.RESOLVING,
					append(opts,
						diag.WithType(diag.RESOLVING),
						diag.WithSeverity(diag.IGNORE),
						ParseSummaryMessage(err),
						diag.WithDetails("The action is invalid for the service."),
					)...),
				),
			}
		}
	}
	if IsErrorThrottle(err) {
		return diag.Diagnostics{
			RedactError(accounts, diag.NewBaseError(err,
				diag.THROTTLE,
				append(opts,
					diag.WithType(diag.THROTTLE),
					diag.WithSeverity(diag.WARNING),
					ParseSummaryMessage(err),
					diag.WithDetails("CloudQuery AWS provider has been throttled, increase max_retries in provider configuration."),
				)...),
			),
		}
	}

	// Take over from SDK and always return diagnostics, redacting PII
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(accounts, diag.NewBaseError(d, d.Type(), opts...)),
		}
	}

	return diag.Diagnostics{
		RedactError(accounts, diag.NewBaseError(err, fallbackType, opts...)),
	}
}

func ParseSummaryMessage(err error) diag.BaseErrorOption {
	var (
		ae     smithy.APIError
		errMsg string
	)
	if errors.As(err, &ae) {
		errMsg = ae.ErrorMessage()
	}

	for {
		if op, ok := err.(*smithy.OperationError); ok {
			if errMsg == "" {
				if op.Err != nil {
					errMsg = op.Err.Error()
				} else {
					errMsg = err.Error()
				}
			}
			return diag.WithError(fmt.Errorf("%s: %s - %s", op.Service(), op.Operation(), errMsg))
		}
		if err2 := errors.Unwrap(err); err2 != nil {
			err = err2
			continue
		}

		if errMsg == "" {
			errMsg = err.Error()
		}
		return diag.WithError(errors.New(errMsg))
	}
}

// RedactError redacts a given diagnostic and returns a RedactedDiagnostic containing both original and redacted versions
func RedactError(aa []Account, e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		nil,
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
	if errors.As(err, &ae) && isCodeThrottle(ae.ErrorCode()) {
		return true
	}
	var qe ratelimit.QuotaExceededError
	return errors.As(err, &qe)
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
	requestIdRegex         = regexp.MustCompile(`\s([Rr]equest[ _]{0,1}(ID|Id|id):)\s[A-Za-z0-9-]+`)
	hostIdRegex            = regexp.MustCompile(`\sHostID: [A-Za-z0-9+/_=-]+`)
	arnIdRegex             = regexp.MustCompile(`(\s)(arn:aws[A-Za-z0-9-]*:)[^ \.\(\)\[\]\{\}\;\,]+(\s?)`)
	urlRegex               = regexp.MustCompile(`([\s"])http(s?):\/\/[a-z0-9_\-\./]+([":\s]?)`)
	lookupRegex            = regexp.MustCompile(`(\slookup\s)[-A-Za-z0-9\.]+\son\s([0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}:[0-9]{1,5})(:.+?)([0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}:[0-9]{1,5})->([0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}:[0-9]{1,5})(:.*)`)
	dialRegex              = regexp.MustCompile(`(\sdial\s)(tcp|udp)(\s)([0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}:[0-9]{1,5})(:.+?)`)
	encAuthRegex           = regexp.MustCompile(`(\s)(Encoded authorization failure message:)\s[A-Za-z0-9_-]+`)
	userRegex              = regexp.MustCompile(`(\s)(is not authorized to perform: .+ on resource:\s)(user)\s.+`)
	s3Regex                = regexp.MustCompile(`(\s)(S3(Key|Bucket))=(.+?)([,;\s])`)
	resourceNotExistsRegex = regexp.MustCompile(`(\sThe )([A-Za-z0-9 -]+ )'([A-Za-z0-9-]+?)'( does not exist)`)
	resourceNotFoundRegex  = regexp.MustCompile(`([A-Za-z0-9 -]+)( name not found - Could not find )([A-Za-z0-9 -]+)( named )'([A-Za-z0-9-]+?)'`)
)

func removePII(aa []Account, msg string) string {
	for i := range aa {
		msg = strings.ReplaceAll(msg, " AccountID "+aa[i].ID, " AccountID xxxx")
	}
	msg = requestIdRegex.ReplaceAllString(msg, " ${1} xxxx")
	msg = hostIdRegex.ReplaceAllString(msg, " HostID: xxxx")
	msg = arnIdRegex.ReplaceAllString(msg, "${1}${2}xxxx${3}")
	msg = urlRegex.ReplaceAllString(msg, "${1}http${2}://xxxx${3}")
	msg = lookupRegex.ReplaceAllString(msg, "${1}xxxx${3}xxxx->xxxx${6}")
	msg = dialRegex.ReplaceAllString(msg, "${1}${2}${3}xxxx${5}")
	msg = encAuthRegex.ReplaceAllString(msg, "${1}${2} xxxx")
	msg = userRegex.ReplaceAllString(msg, "${1}${2}${3} xxxx")
	msg = s3Regex.ReplaceAllString(msg, "${1}${2}=xxxx${5}")
	msg = resourceNotExistsRegex.ReplaceAllString(msg, "${1}${2}'xxxx'${4}")
	msg = resourceNotFoundRegex.ReplaceAllString(msg, "${1}${2}${3}${4}'xxxx'")
	msg = accountObfusactor(aa, msg)

	return msg
}

func includeResourceIdWithAccount(client *Client, err error) diag.BaseErrorOption {
	d, ok := err.(diag.Diagnostic)
	if !ok || len(d.Description().ResourceID) == 0 {
		return func(_ *diag.BaseError) {} // no-op option
	}

	resIdList := []string{
		client.AccountID,
		client.Region,
	}

	// remove accountID and region from PK list as we always prepend them
	for _, val := range d.Description().ResourceID {
		if val != client.AccountID && val != client.Region {
			resIdList = append(resIdList, val)
		}
	}

	return diag.WithResourceId(resIdList)
}
