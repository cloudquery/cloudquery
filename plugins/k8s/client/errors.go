package client

import (
	"errors"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	k8s "k8s.io/apimachinery/pkg/api/errors"
)

func ErrorClassifier(_ schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	return classifyError(err, diag.RESOLVING, diag.WithResourceName(resourceName))
}

func classifyError(err error, fallbackType diag.Type, opts ...diag.BaseErrorOption) diag.Diagnostics {
	ie := errors.Unwrap(err)
	if se, ok := ie.(k8s.APIStatus); ok {
		if se.Status().Code == 403 {
			return diag.FromError(ie, diag.ACCESS, diag.WithSeverity(diag.WARNING), diag.WithDetails(se.Status().Details.String()))
		}
	}
	return diag.Diagnostics{diag.NewBaseError(err, fallbackType, opts...)}
}

func IgnoreForbiddenNotFound(err error) bool {
	statusError, ok := err.(k8s.APIStatus)
	if !ok {
		return false
	}
	if statusError.Status().Code == 403 || statusError.Status().Code == 404 {
		return true
	}
	return false
}
