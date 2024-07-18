package client

import (
	"errors"
	"net/http"

	"google.golang.org/api/googleapi"
)

func isAPINotFoundError(err error) bool {
	var apiErr *googleapi.Error
	if !errors.As(err, &apiErr) {
		return false
	}
	return apiErr.Code == http.StatusNotFound
}

func isEntityTooLargeError(err error) bool {
	var apiErr *googleapi.Error
	if !errors.As(err, &apiErr) {
		return false
	}
	return apiErr.Code == http.StatusRequestEntityTooLarge
}
