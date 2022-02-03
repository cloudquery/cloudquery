package telemetry

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/jackc/pgconn"
	"github.com/lib/pq"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecordError should be called on a span to mark it as errored. Returns true if err was recorded.
func RecordError(span trace.Span, err error, opts ...trace.EventOption) bool {
	if err == nil {
		return false
	}

	if rd, ok := err.(diag.Redactable); ok {
		if r := rd.Redacted(); r != nil {
			err = r
		}
	}

	if cls := classifyError(err); cls != errNoClass {
		span.SetStatus(codes.Error, string(cls))
		return false
	}

	span.RecordError(err, opts...)
	span.SetStatus(codes.Error, err.Error())
	return true
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

	if errors.Is(err, context.Canceled) {
		return errCancellation
	}

	{
		var ope *net.OpError
		if errors.As(err, &ope) && ope.Op == "dial" {
			return errConn
		}
	}

	{
		var (
			pgCode string
			pqe    *pq.Error
			pge    *pgconn.PgError
		)
		if errors.As(err, &pqe) {
			pgCode = string(pqe.Code)
		} else if errors.As(err, &pge) {
			pgCode = pge.Code
		}
		if len(pgCode) >= 2 {
			switch pgCode[0:2] {
			// Class 28 - Invalid Authorization Specification
			// Class 3D - Invalid Catalog Name
			// Class 57 - Operator Intervention
			case "28", "3D", "57":
				return errDatabase
			}
		}
	}

	if errors.Is(err, pq.ErrSSLNotSupported) ||
		errors.Is(err, pq.ErrSSLKeyHasWorldPermissions) {
		return errDatabase
	}

	return errNoClass
}
