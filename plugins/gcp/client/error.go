package client

import (
	"errors"

	"google.golang.org/api/googleapi"
)

const QuotaExceeded = 429
const Forbidden = 403

func IgnoreErrorHandler(err error) bool {
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if gerr.Code == Forbidden && len(gerr.Errors) > 0 && gerr.Errors[0].Reason == "accessNotConfigured" {
			return true
		}
		if gerr.Code == Forbidden && len(gerr.Errors) > 0 && gerr.Errors[0].Reason == "forbidden" {
			return true
		}
	}
	return false
}
