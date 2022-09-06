package client

import (
	"errors"
	"net/http"

	"github.com/google/go-github/v45/github"
)

const EnterpriseOnly = "This organization is not part of externally managed enterprise."

func IgnoreError(err error) (bool, string) {
	var er *github.ErrorResponse
	if errors.As(err, &er) {
		if er.Response.StatusCode == http.StatusNotFound {
			return true, "not_found"
		}
		if er.Message == EnterpriseOnly {
			return true, "permission_denied"
		}
		if er.Response.StatusCode == http.StatusForbidden {
			return true, "permission_denied"
		}
	}
	return false, ""
}
