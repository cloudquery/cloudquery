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

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	client := meta.(*Client)

	var (
		detailedError autorest.DetailedError
		reqError      azure.RequestError
	)
	if errors.As(err, &detailedError) {
		if errors.As(detailedError.Original, &reqError) && reqError.ServiceError != nil && reqError.ServiceError.Code == "DisallowedOperation" {
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.ACCESS, diag.WithType(diag.ACCESS), diag.WithSeverity(diag.WARNING), diag.WithResourceName(resourceName), ParseSummaryMessage(client.SubscriptionId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))),
			}
		}

		switch detailedError.StatusCode {
		case http.StatusNotFound:
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.RESOLVING, diag.WithType(diag.RESOLVING), diag.WithSeverity(diag.IGNORE), diag.WithResourceName(resourceName), ParseSummaryMessage(client.SubscriptionId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))),
			}
		case http.StatusBadRequest:
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.RESOLVING, diag.WithType(diag.RESOLVING), diag.WithSeverity(diag.WARNING), diag.WithResourceName(resourceName), ParseSummaryMessage(client.SubscriptionId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))),
			}
		case http.StatusForbidden:
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.ACCESS, diag.WithType(diag.ACCESS), diag.WithSeverity(diag.WARNING), diag.WithResourceName(resourceName), ParseSummaryMessage(client.SubscriptionId, err, detailedError), diag.WithDetails("%s", errorCodeDescriptions[detailedError.StatusCode]))),
			}
		default:
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.RESOLVING, ParseSummaryMessage(client.SubscriptionId, err, detailedError))),
			}
		}
	}

	// Take over from SDK and always return diagnostics, redacting PII
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(client.SubscriptionId, d),
		}
	}

	return diag.Diagnostics{
		RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.RESOLVING, diag.WithResourceName(resourceName))),
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

var (
	uuidRegex = regexp.MustCompile(`(\W)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}(\W)`)
)

func removePII(subId string, msg string) string {
	msg = strings.ReplaceAll(msg, subId, "xxxx")
	msg = uuidRegex.ReplaceAllString(msg, "${1}xxxx${2}")
	return msg
}
