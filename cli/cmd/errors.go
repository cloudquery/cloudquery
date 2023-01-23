package cmd

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// isUnimplemented returns true if an error indicates that the underlying grpc call
// was unimplemented on the server side.
func isUnimplemented(err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	if ok && st.Code() == codes.Unimplemented {
		return true
	}
	err = errors.Unwrap(err)
	return isUnimplemented(err)
}
