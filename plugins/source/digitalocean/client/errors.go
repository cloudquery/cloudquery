package client

import (
	"regexp"

	"github.com/digitalocean/godo"
	"github.com/pkg/errors"
)

var (
	requestIdRegex = regexp.MustCompile(`([Rr]equest[: ]+)\"[A-Za-z0-9-]+\"`)
)

func ClassifyError(err error) (bool, string) {
	var ae *godo.ErrorResponse

	if errors.As(err, &ae) {
		switch ae.Message {
		case "permission denied", "Unable to authenticate you", "API Rate limit exceeded.", "Too many requests":
			return true, removePII(ae.Message)
		}
	}
	return false, ""
}

func removePII(msg string) string {
	msg = requestIdRegex.ReplaceAllString(msg, "${1} xxxx")
	return msg
}

func IsErrorMessage(err error, message string) bool {
	var ae *godo.ErrorResponse
	if !errors.As(err, &ae) {
		return false
	}
	if message == ae.Message {
		return true
	}
	return false
}
