package client

import (
	"errors"
	"net/http"

	"github.com/google/go-github/v45/github"
)

func IgnoreError(err error) bool {
	var er *github.ErrorResponse
	if errors.As(err, &er) {
		if er.Response.StatusCode == http.StatusNotFound {
			return true
		}
		if er.Response.StatusCode == http.StatusForbidden {
			return true
		}
	}
	return false
}
