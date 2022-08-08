package errors

import (
	"context"
	stderrors "errors"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsCancelation(err error) bool {
	if stderrors.Is(err, context.Canceled) || stderrors.Is(err, context.DeadlineExceeded) {
		return true
	}

	if st, ok := status.FromError(err); ok && (st.Code() == gcodes.Canceled || st.Code() == gcodes.DeadlineExceeded) {
		return true
	}

	return false
}

func CancelationDiag(err error) diag.Diagnostics {
	return diag.Diagnostics{diag.NewBaseError(err, diag.USER, diag.WithSummary("operation was canceled by user"))}
}
