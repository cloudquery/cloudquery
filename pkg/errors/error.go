package errors

import (
	"context"
	"errors"
	"net"
	"regexp"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/lib/pq"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errClass string

type sentryDiag interface {
	IsSentryDiagnostic() (bool, map[string]string, bool)
}

const (
	errNoClass      = errClass("")
	errCancellation = errClass("cancelled")
	errAuth         = errClass("auth")
	errConn         = errClass("connection")
	errDatabase     = errClass("database")
)

var sqlStateRegex = regexp.MustCompile(`\(SQLSTATE ([0-9A-Z]{5})\)`)

// classifyError classifies given error by type and internals. Successfully classified (not errNoClass) errors don't get reported to sentry.
func classifyError(err error) errClass {
	if st, ok := status.FromError(err); ok {
		if st.Code() == gcodes.Canceled {
			return errCancellation
		}
		if strings.HasPrefix(st.Message(), `failed to create aws client for`) ||
			strings.HasPrefix(st.Message(), `failed to retrieve credentials for`) {
			return errAuth
		}
	}

	if errors.Is(err, context.Canceled) || strings.Contains(err.Error(), context.Canceled.Error()) {
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
		if shouldIgnorePgCode(pgCode) {
			return errDatabase
		}
	}

	if errors.Is(err, pq.ErrSSLNotSupported) ||
		errors.Is(err, pq.ErrSSLKeyHasWorldPermissions) {
		return errDatabase
	}

	return errNoClass
}

func shouldIgnorePgCode(code string) bool {
	if len(code) >= 2 {
		switch code[0:2] {
		// https://www.postgresql.org/docs/9.3/errcodes-appendix.html
		// Class 08 - Connection Exception
		// Class 28 - Invalid Authorization Specification
		// Class 3D - Invalid Catalog Name
		// Class 3F - Invalid Schema Name
		// Class 53 - Insufficient Resources
		// Class 57 - Operator Intervention
		case "08", "28", "3D", "3F", "53", "57":
			return true
		}
	}
	return false
}
