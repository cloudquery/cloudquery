package client

import (
	"errors"
	heroku "github.com/heroku/heroku-go/v5"
)

// IgnoreError returns true if the given error can be ignored, and false
// otherwise. We currently don't ignore any errors, so this function
// always return false. It exists only to explicitly document the types of
// errors we return to the user.
func IgnoreError(err error) bool {
	var er *heroku.Error
	if errors.As(err, &er) {
		// Heroku API errors are not currently ignored
		if er.StatusCode >= 500 {
			return false
		}
		// These are all the errors that can be returned by the Heroku API,
		// according to the docs: https://devcenter.heroku.com/articles/platform-api-reference#client-error-responses
		// We don't currently ignore any of them.
		switch er.ID {
		case "bad_request", "unauthorized", "delinquent", "forbidden",
			"suspended", "not_found", "not_acceptable", "conflict", "gone",
			"requested_range_not_satisfiable", "invalid_params",
			"verification_needed", "rate_limit":
			return false
		}
	}
	// Default: return false
	return false
}
