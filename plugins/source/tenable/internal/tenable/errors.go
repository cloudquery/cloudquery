package tenable

import "errors"

var ErrUnauthorized = errors.New("unauthorized: invalid or missing API keys")
