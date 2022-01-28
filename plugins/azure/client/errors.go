package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/provider/schema/diag"
)

var errorCodeDescriptions = map[interface{}]string{
	http.StatusNotFound:   "The requested resource could not be found.",
	http.StatusBadRequest: "Bad request",
	http.StatusForbidden:  "You are not authorized to perform this operation.",
}

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) []diag.Diagnostic {
	var detailedError autorest.DetailedError
	if !errors.As(err, &detailedError) {
		return nil
	}
	client := meta.(*Client)
	switch detailedError.StatusCode {
	case http.StatusNotFound:
		return []diag.Diagnostic{
			diag.FromError(err, diag.IGNORE, diag.RESOLVING, resourceName, ParseSummaryMessage(client.SubscriptionId, err, detailedError), errorCodeDescriptions[detailedError.StatusCode]),
		}
	case http.StatusBadRequest:
		return []diag.Diagnostic{
			diag.FromError(err, diag.WARNING, diag.RESOLVING, resourceName, ParseSummaryMessage(client.SubscriptionId, err, detailedError), errorCodeDescriptions[detailedError.StatusCode]),
		}
	case http.StatusForbidden:
		return []diag.Diagnostic{
			diag.FromError(err, diag.WARNING, diag.ACCESS, resourceName, ParseSummaryMessage(client.SubscriptionId, err, detailedError), errorCodeDescriptions[detailedError.StatusCode]),
		}
	}
	return nil
}

func ParseSummaryMessage(subscriptionId string, err error, detailedError autorest.DetailedError) string {
	for {
		if de, ok := err.(autorest.DetailedError); ok {
			return fmt.Sprintf("%s: %s - %s", de.Method, de.PackageType, detailedError.Error())
		}
		if err = errors.Unwrap(err); err == nil {
			return detailedError.Error()
		}
	}
}
