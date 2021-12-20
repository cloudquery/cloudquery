package telemetry

import (
	"errors"
	"net"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/lib/pq"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecordError should be called on a span to mark it as errored. By default error values are not recorded, unless the debug flag is set.
func RecordError(span trace.Span, err error, opts ...trace.EventOption) {
	if err == nil {
		return
	}

	if cls := classifyError(err); cls != errNoClass {
		span.SetStatus(codes.Error, string(cls))
		return
	}

	sentry.CaptureException(err)

	span.RecordError(err, opts...)
	span.SetStatus(codes.Error, err.Error())
}

type errClass string

const (
	errNoClass      = errClass("")
	errCancellation = errClass("cancelled")
	errAuth         = errClass("auth")
	errConn         = errClass("connection")
	errDatabase     = errClass("database")
)

func classifyError(err error) errClass {
	if st, ok := status.FromError(err); ok {
		if st.Code() == gcodes.Canceled {
			return errCancellation
		}
		// if strings.Contains(st.Message(), `AWS Error: operation error STS: GetCallerIdentity,`) {
		if strings.HasPrefix(st.Message(), `failed to create aws client for`) ||
			strings.HasPrefix(st.Message(), `failed to retrieve credentials for`) {
			return errAuth
		}
	}

	{
		var ope *net.OpError
		if errors.As(err, &ope) && ope.Op == "dial" {
			return errConn
		}
	}

	{
		var pge *pq.Error
		if errors.As(err, &pge) {
			switch pge.Code.Class() {
			// Class 28 - Invalid Authorization Specification
			// Class 3D - Invalid Catalog Name
			case "28", "3D":
				return errDatabase
			}
		}
	}

	return errNoClass
}
