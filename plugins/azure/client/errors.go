package client

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var errorCodeDescriptions = map[interface{}]string{
	http.StatusNotFound:   "The requested resource could not be found.",
	http.StatusBadRequest: "Bad request",
	http.StatusForbidden:  "You are not authorized to perform this operation.",
}

var (
	uuidRegex = regexp.MustCompile(`(\W)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}(\W)`)
)

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	client := meta.(*Client)
	return classifyError(err, diag.RESOLVING, client.SubscriptionId, diag.WithResourceName(resourceName))
}

func classifyError(err error, fallbackType diag.Type, subId string, opts ...diag.BaseErrorOption) diag.Diagnostics {
	var (
		detailedError autorest.DetailedError
		reqError      *azure.RequestError
	)
	if errors.As(err, &detailedError) {
		if errors.As(detailedError.Original, &reqError) && reqError.ServiceError != nil {
			switch reqError.ServiceError.Code {
			case "DisallowedOperation":
				return diag.Diagnostics{
					RedactError(subId, diag.NewBaseError(err, diag.ACCESS, append(opts, diag.WithType(diag.ACCESS), diag.WithSeverity(diag.WARNING), ParseSummaryMessage(subId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))...)),
				}
			case "SubscriptionNotRegistered":
				return diag.Diagnostics{
					RedactError(subId, diag.NewBaseError(err, diag.ACCESS, append(opts, diag.WithType(diag.ACCESS), diag.WithSeverity(diag.WARNING), ParseSummaryMessage(subId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))...)),
				}
			}
		}

		switch detailedError.StatusCode {
		case http.StatusNotFound:
			return diag.Diagnostics{
				RedactError(subId, diag.NewBaseError(err, diag.RESOLVING, append(opts, diag.WithType(diag.RESOLVING), diag.WithSeverity(diag.IGNORE), ParseSummaryMessage(subId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))...)),
			}
		case http.StatusBadRequest:
			return diag.Diagnostics{
				RedactError(subId, diag.NewBaseError(err, diag.RESOLVING, append(opts, diag.WithType(diag.RESOLVING), diag.WithSeverity(diag.WARNING), ParseSummaryMessage(subId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))...)),
			}
		case http.StatusForbidden:
			return diag.Diagnostics{
				RedactError(subId, diag.NewBaseError(err, diag.ACCESS, append(opts, diag.WithType(diag.ACCESS), diag.WithSeverity(diag.WARNING), ParseSummaryMessage(subId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))...)),
			}
		default:
			return diag.Diagnostics{
				RedactError(subId, diag.NewBaseError(err, fallbackType, append(opts, ParseSummaryMessage(subId, err, detailedError))...)),
			}
		}
	}

	// Take over from SDK and always return diagnostics, redacting PII
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(subId, d),
		}
	}

	return diag.Diagnostics{
		RedactError(subId, diag.NewBaseError(err, fallbackType, opts...)),
	}
}

func ParseSummaryMessage(subscriptionId string, err error, detailedError autorest.DetailedError) diag.BaseErrorOption {
	for {
		if de, ok := err.(autorest.DetailedError); ok {
			return diag.WithError(fmt.Errorf("%s: %s - %s", de.Method, de.PackageType, detailedError.Error()))
		}
		if err = errors.Unwrap(err); err == nil {
			return diag.WithError(errors.New(detailedError.Error()))
		}
	}
}

// RedactError redacts a given diagnostic and returns a RedactedDiagnostic containing both original and redacted versions
func RedactError(subId string, e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		errors.New(removePII(subId, e.Error())),
		e.Type(),
		diag.WithSeverity(e.Severity()),
		diag.WithResourceName(e.Description().Resource),
		diag.WithSummary("%s", removePII(subId, e.Description().Summary)),
		diag.WithDetails("%s", removePII(subId, e.Description().Detail)),
	)
	return diag.NewRedactedDiagnostic(e, r)
}

func removePII(subId string, msg string) string {
	if subId != "" {
		msg = strings.ReplaceAll(msg, subId, "xxxx")
	}
	msg = uuidRegex.ReplaceAllString(msg, "${1}xxxx${2}")
	return msg
}
