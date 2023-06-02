package client

import (
	"errors"
	"fmt"

	"github.com/okta/okta-sdk-golang/v3/okta"
)

func ProcessOktaAPIError(err error) error {
	var genericOpenAPIErr *okta.GenericOpenAPIError
	if errors.As(err, &genericOpenAPIErr) {
		return fmt.Errorf(`received error from Okta API: err: %s, body: %q model: %v`, genericOpenAPIErr.Error(), genericOpenAPIErr.Body(), genericOpenAPIErr.Model())
	}

	return err
}
