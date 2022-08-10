package client

import (
	"errors"
	"net/http"

	"github.com/google/go-github/v45/github"
)

const EnterpriseOnly = "This organization is not part of externally managed enterprise."

func IgnoreError(err error) bool {
	var er *github.ErrorResponse
	if errors.As(err, &er) {
		if er.Response.StatusCode == http.StatusNotFound {
			return true
		}
		if er.Message == EnterpriseOnly {
			return true
		}
	}
	return false
}
