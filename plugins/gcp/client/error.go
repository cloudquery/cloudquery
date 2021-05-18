package client

import (
	"google.golang.org/api/googleapi"
)

const QuotaExceeded = 429
const Forbidden = 403

func IgnoreErrorHandler(err error) bool {
	if e, ok := err.(*googleapi.Error); ok {
		if e.Code == Forbidden && len(e.Errors) > 0 && e.Errors[0].Reason == "accessNotConfigured" {
			return true
		} else if e.Code == Forbidden && len(e.Errors) > 0 && e.Errors[0].Reason == "forbidden" {
			return true
		}
	}
	return false
}
