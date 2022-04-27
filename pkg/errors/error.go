package errors

import (
	"context"
	"errors"
	"net"
	"regexp"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/jackc/pgconn"
	"github.com/lib/pq"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var sqlStateRegex = regexp.MustCompile(`\(SQLSTATE ([0-9A-Z]{5})\)`)

// ShouldIgnoreDiag checks the wire-transferred diagnostic against errors we don't want to process.
func ShouldIgnoreDiag(d diag.Diagnostic) bool {
	if d.Severity() == diag.IGNORE || (d.Severity() == diag.WARNING && d.Type() == diag.ACCESS) {
		return true
	}
	if d.Type() == diag.DATABASE {
		ret := sqlStateRegex.FindStringSubmatch(d.Error())
		if len(ret) > 1 && shouldIgnorePgCode(ret[1]) {
			return true
		}
	}
	return false
}

type errClass string

const (
	errNoClass      = errClass("")
	errCancellation = errClass("cancelled")
	errAuth         = errClass("auth")
	errConn         = errClass("connection")
	errDatabase     = errClass("database")
)

// TODO: decide what to do with this method
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
		// Class 28 - Invalid Authorization Specification
		// Class 3D - Invalid Catalog Name
		// Class 57 - Operator Intervention
		case "28", "3D", "57":
			return true
		}
	}
	return false
}
