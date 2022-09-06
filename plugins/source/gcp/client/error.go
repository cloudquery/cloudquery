package client

import (
	"errors"
	"net/http"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	// ""
)

func ClassifyError(err error) (bool, string) {
	var gerr *googleapi.Error
	var apierr *apierror.APIError
	if ok := errors.As(err, &gerr); ok {
		if gerr.Code == http.StatusForbidden && len(gerr.Errors) > 0 {
			switch gerr.Errors[0].Reason {
			case "accessNotConfigured", "forbidden", "SERVICE_DISABLED":
				return true, gerr.Errors[0].Reason
			}
		}
		if gerr.Code == http.StatusNotFound && len(gerr.Errors) > 0 {
			return true, "notFound"
		}
	} else if ok := errors.As(err, &apierr); ok {
		if apierr.GRPCStatus().Code() == codes.PermissionDenied {
			return true, "permission_denied"
		}
	}
	return false, ""
}
