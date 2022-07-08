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

const ssoInvalidOrExpired = "failed to refresh cached credentials, the SSO session has expired or is invalid"

var (
	ipv4Regex          = `\d+\.\d+\.\d+\.\d+`
	ipv6Regex          = `(?:(?:[a-fA-F0-9]{0,4}:)){2,7}[a-fA-F0-9]{0,4}`
	bracketedIpv6Regex = fmt.Sprintf(`\[%s\]`, ipv6Regex)

	requestIdRegex = regexp.MustCompile(`\s([Rr]equest[ _]{0,1}(ID|Id|id):)\s[A-Za-z0-9-]+`)
	hostIdRegex    = regexp.MustCompile(`\sHostID: [A-Za-z0-9+/_=-]+`)
	arnIdRegex     = regexp.MustCompile(`(\s)(arn:aws[A-Za-z0-9-]*:)[^ \.\(\)\[\]\{\}\;\,]+(\s?)`)
	urlRegex       = regexp.MustCompile(`([\s"])http(s?):\/\/[a-z0-9_\-\./]+([":\s]?)`)
	lookupRegex    = regexp.MustCompile(
		`\blookup\s[-A-Za-z0-9\.]+\s` + // " lookup host.name "
			`on\s\S+:\d+`, // "on 123.123.123.123:53"
	)
	readXonYRegex = regexp.MustCompile(
		`\bread\s(udp|tcp)\s` + // "read udp "
			`\S+:\d+->\S+:\d+`, // "192.168.1.2:5353->192.168.1.1:53"
	)
	dialRegex = regexp.MustCompile(
		`\bdial\s(tcp|udp)\s` + // "dial tcp "
			fmt.Sprintf(`(?:%s|%s):\d+`, ipv4Regex, bracketedIpv6Regex), // "192.168.1.2:123" or "[::1]:123"
	)
	ec2ImageIdRegex          = regexp.MustCompile(`The image ID '[^']+'`)
	encAuthRegex             = regexp.MustCompile(`(\s)(Encoded authorization failure message:)\s[A-Za-z0-9_-]+`)
	userRegex                = regexp.MustCompile(`(\s)(is not authorized to perform: .+ on resource:\s)(user)\s.+`)
	s3Regex                  = regexp.MustCompile(`(\s)(S3(Key|Bucket))=(.+?)([,;\s])`)
	resourceNotExistsRegex   = regexp.MustCompile(`(\sThe )([A-Za-z0-9 -]+ )'([A-Za-z0-9-]+?)'( does not exist)`)
	resourceNotFoundRegex    = regexp.MustCompile(`([A-Za-z0-9 -]+)( name not found - Could not find )([A-Za-z0-9 -]+)( named )'([A-Za-z0-9-]+?)'`)
	autoscalingGroupNotFound = regexp.MustCompile(`(ValidationError: Group ).+( not found)`)
)

var errorCodeDescriptions = map[string]string{
	"InvalidClientTokenId":          "The X.509 certificate or AWS access key ID provided does not exist in our records.",
	"SubscriptionRequiredException": "When you created your AWS account, all available services at that time were activated. However, as new services are released, they aren't automatically put into an active state without your permission. You must subscribe to each service individually as they are released.",
	"OptInRequired":                 "You are not authorized to use the requested service. Ensure that you have subscribed to the service you are trying to use. If you are new to AWS, your account might take some time to be activated while your credit card details are being verified.",
	"UnauthorizedOperation":         "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
	"AccessDeniedException":         "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
	"AccessDenied":                  "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
	"AuthorizationError":            "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
	"AuthFailure":                   "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys.",
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

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	client := meta.(*Client)

	return classifyError(err, diag.RESOLVING, client.ServicesManager.services.Accounts(), diag.WithResourceName(resourceName), includeResourceIdWithAccount(client, err))
}

func classifyError(err error, fallbackType diag.Type, accounts []string, opts ...diag.BaseErrorOption) diag.Diagnostics {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation", "AuthorizationError", "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId", "AuthFailure", "ExpiredToken", "ExpiredTokenException", "FailedResourceAccessException":
			return diag.Diagnostics{
				RedactError(accounts, diag.NewBaseError(err,
					diag.ACCESS,
					append(opts,
						diag.WithType(diag.ACCESS),
						diag.WithSeverity(diag.WARNING),
						ParseSummaryMessage(err),
						diag.WithDetails(errorCodeDescriptions[ae.ErrorCode()]),
					)...),
				),
			}
		case "UnrecognizedClientException":
			if strings.Contains(ae.Error(), "The security token included in the request is invalid") {
				return diag.Diagnostics{
					RedactError(accounts, diag.NewBaseError(err,
						diag.ACCESS,
						append(opts,
							diag.WithType(diag.ACCESS),
							diag.WithSeverity(diag.WARNING),
							ParseSummaryMessage(err),
							diag.WithDetails("Something is wrong with your credentials. Either the accessKeyId or secretAccessKey (or both) are wrong."),
						)...),
					),
				}
			}
		case "MetadataException":
			if strings.Contains(ae.Error(), "is not authorized to perform") {
				return diag.Diagnostics{
					RedactError(accounts, diag.NewBaseError(err,
						diag.ACCESS,
						append(opts,
							diag.WithType(diag.ACCESS),
							diag.WithSeverity(diag.WARNING),
							ParseSummaryMessage(err),
							diag.WithDetails("Something is wrong with your credentials. Ensure you have access to the specified resource."),
						)...),
					),
				}
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
		case "UnsupportedOperation":
			if strings.Contains(ae.ErrorMessage(), "The functionality you requested is not available in this region.") {
				return diag.Diagnostics{
					RedactError(accounts, diag.NewBaseError(err,
						diag.RESOLVING,
						append(opts,
							diag.WithType(diag.RESOLVING),
							diag.WithSeverity(diag.IGNORE),
							ParseSummaryMessage(err),
							diag.WithDetails("The action is not available in selected region"))...),
					),
				}
			}
		}
		if ae.ErrorMessage() == ssoInvalidOrExpired {
			return diag.Diagnostics{
				RedactError(accounts, diag.NewBaseError(err,
					diag.ACCESS,
					append(opts,
						diag.WithType(diag.ACCESS),
						diag.WithSeverity(diag.WARNING),
						ParseSummaryMessage(err),
						diag.WithDetails(errorCodeDescriptions[ae.ErrorCode()]),
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
					diag.WithDetails("CloudQuery AWS provider has been throttled. This is unexpected - you can open an issue on github (https://github.com/cloudquery/cq-provider-aws/issues) or contact us on discord (https://cloudquery.io/discord)"),
				)...),
			),
		}
	}
	if strings.Contains(err.Error(), "socket: too many open files") {
		return diag.Diagnostics{
			RedactError(accounts, diag.NewBaseError(err,
				diag.THROTTLE,
				append(opts,
					diag.WithType(diag.THROTTLE),
					diag.WithSeverity(diag.WARNING),
					ParseSummaryMessage(err),
					diag.WithDetails("CloudQuery AWS provider has been throttled. Too many open files, try to increase your max file descriptors in your system or contact us on discord (https://cloudquery.io/discord)"),
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
func RedactError(aa []string, e diag.Diagnostic) diag.Diagnostic {
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

func isCodeThrottle(code string) bool {
	_, ok := throttleCodes[code]
	return ok
}

func removePII(aa []string, msg string) string {
	for _, i := range aa {
		msg = strings.ReplaceAll(msg, " AccountID "+i, " AccountID xxxx")
	}
	msg = requestIdRegex.ReplaceAllString(msg, " ${1} xxxx")
	msg = hostIdRegex.ReplaceAllString(msg, " HostID: xxxx")
	msg = arnIdRegex.ReplaceAllString(msg, "${1}${2}xxxx${3}")
	msg = urlRegex.ReplaceAllString(msg, "${1}http${2}://xxxx${3}")
	msg = lookupRegex.ReplaceAllString(msg, "lookup xxxx on xxxx:xx")
	msg = readXonYRegex.ReplaceAllString(msg, "read $1 xxxx:xx->xxxx:xx")
	msg = dialRegex.ReplaceAllString(msg, "dial $1 xxxx:xx")
	msg = encAuthRegex.ReplaceAllString(msg, "${1}${2} xxxx")
	msg = userRegex.ReplaceAllString(msg, "${1}${2}${3} xxxx")
	msg = s3Regex.ReplaceAllString(msg, "${1}${2}=xxxx${5}")
	msg = resourceNotExistsRegex.ReplaceAllString(msg, "${1}${2}'xxxx'${4}")
	msg = resourceNotFoundRegex.ReplaceAllString(msg, "${1}${2}${3}${4}'xxxx'")
	msg = ec2ImageIdRegex.ReplaceAllString(msg, "The image ID 'xxxx'")
	msg = autoscalingGroupNotFound.ReplaceAllString(msg, "${1}xxxx${2}")
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
