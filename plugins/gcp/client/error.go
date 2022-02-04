package client

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IgnoreErrorHandler(err error) bool {
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if gerr.Code == http.StatusForbidden && len(gerr.Errors) > 0 {
			switch gerr.Errors[0].Reason {
			case "accessNotConfigured", "forbidden", "SERVICE_DISABLED":
				return true
			}
		}
	}
	return false
}

type diagValue struct {
	severity diag.Severity
	typ      diag.DiagnosticType
	summary  string
}

var grpcCodeToDiag = map[codes.Code]diagValue{
	codes.PermissionDenied:  {diag.WARNING, diag.ACCESS, "Access denied"},
	codes.Unauthenticated:   {diag.WARNING, diag.ACCESS, "Authentication failure"},
	codes.ResourceExhausted: {diag.WARNING, diag.THROTTLE, "Resource exhausted (quota etc)"},
	codes.Unimplemented:     {diag.IGNORE, diag.RESOLVING, "Operation not implemented or not supported"},
}

var httpCodeToGRPCCode = map[int]codes.Code{
	http.StatusForbidden:       codes.PermissionDenied,
	http.StatusUnauthorized:    codes.Unauthenticated,
	http.StatusTooManyRequests: codes.ResourceExhausted,
	http.StatusNotImplemented:  codes.Unimplemented,
}

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	// https://pkg.go.dev/cloud.google.com/go#hdr-Inspecting_errors:
	// Most of the errors returned by the generated clients can be converted into a `grpc.Status`
	if err == nil {
		return nil
	}
	client := meta.(*Client)

	// Don't override if already a diagnostic, just redact
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(client.projects, d),
		}
	}

	if s, ok := status.FromError(err); ok {
		if v, ok := grpcCodeToDiag[s.Code()]; ok {
			return diag.Diagnostics{
				RedactError(client.projects, diag.NewBaseError(err, v.severity, v.typ, resourceName, v.summary, s.Message())),
			}
		}
	}

	// as a fallback, try to convert the error to *googleapi.Error
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if grpcCode, ok := httpCodeToGRPCCode[gerr.Code]; ok {
			if v, ok := grpcCodeToDiag[grpcCode]; ok {
				return diag.Diagnostics{
					RedactError(client.projects, diag.NewBaseError(err, v.severity, v.typ, resourceName, v.summary, "")),
				}
			}
		}
	}

	// Take over from SDK and always return diagnostics, redacting PII
	return diag.Diagnostics{
		RedactError(client.projects, diag.NewBaseError(err, diag.ERROR, diag.RESOLVING, resourceName, err.Error(), "")),
	}
}

// RedactError redacts a given diagnostic and returns a RedactedDiagnostic containing both original and redacted versions
func RedactError(projects []string, e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		errors.New(removePII(projects, e.Error())),
		e.Severity(),
		e.Type(),
		e.Description().Resource,
		removePII(projects, e.Description().Summary),
		removePII(projects, e.Description().Detail),
	)
	return diag.NewRedactedDiagnostic(e, r)
}

var (
	codeRegex      = regexp.MustCompile(`\(Code: '[A-Z0-9\.]+'\)`)
	projectIdRegex = regexp.MustCompile(`project(_number|s)?(\W)[0-9]+(\W)`)
	userIdRegex    = regexp.MustCompile(`(\W)[^@ ]+@(.+?)\.iam\.gserviceaccount\.com`)
)

func removePII(projects []string, msg string) string {
	for i := range projects {
		if projects[i] != "" {
			msg = strings.ReplaceAll(msg, projects[i], "xxxx")
		}
	}

	msg = userIdRegex.ReplaceAllString(msg, `${1}xxxx@xxxx.iam.gserviceaccount.com`)
	msg = codeRegex.ReplaceAllLiteralString(msg, `(Code: 'xxxx')`)
	msg = projectIdRegex.ReplaceAllString(msg, `project${1}${2}xxxx${3}`)
	return msg
}
