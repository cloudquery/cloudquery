package client

import (
	"errors"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type diagValue struct {
	severity diag.Severity
	typ      diag.Type
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

var (
	codeRegex      = regexp.MustCompile(`\(Code: '[A-Z0-9\.]+'\)`)
	projectIdRegex = regexp.MustCompile(`project(_number|s)?(\W)[0-9]+(\W)`)
	userIdRegex    = regexp.MustCompile(`(\W)[^@ ]+@(.+?)\.iam\.gserviceaccount\.com`)
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
		if gerr.Code == http.StatusNotFound && len(gerr.Errors) > 0 {
			return true
		}
	}
	return false
}

func IgnoreErrorColumnHandler(err error) bool {
	// gcp API can return an empty string for some CIDR values, so we ignore those
	var parseError *net.ParseError
	if ok := errors.As(err, &parseError); ok {
		if parseError.Type == "CIDR address" {
			return true
		}
	}
	return false
}

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	client := meta.(*Client)

	return classifyError(err, diag.RESOLVING, client.projects, diag.WithResourceName(resourceName))
}

func classifyError(err error, fallbackType diag.Type, projects []string, opts ...diag.BaseErrorOption) diag.Diagnostics {
	// If the error is a diagnostic already then there is no need to classify it.
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(projects, d),
		}
	}

	// https://pkg.go.dev/cloud.google.com/go#hdr-Inspecting_errors:
	// Most of the errors returned by the generated clients can be converted into a `grpc.Status`
	if s, ok := statusFromError(err); ok {
		if v, ok := grpcCodeToDiag[s.Code()]; ok {
			return diag.Diagnostics{
				RedactError(projects,
					diag.NewBaseError(err, v.typ, append(opts,
						diag.WithType(v.typ),
						diag.WithSummary("%s", v.summary),
						diag.WithDetails("%s", s.Message()),
						diag.WithSeverity(v.severity),
					)...),
				),
			}
		}
	}

	// as a fallback, try to convert the error to *googleapi.Error
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if len(gerr.Errors) > 0 && gerr.Errors[0].Reason == "rateLimitExceeded" {
			return diag.Diagnostics{
				RedactError(projects, diag.NewBaseError(err, diag.THROTTLE, append(opts, diag.WithType(diag.THROTTLE), diag.WithSeverity(diag.WARNING), diag.WithError(errors.New(gerr.Message)))...)),
			}
		}

		if grpcCode, ok := httpCodeToGRPCCode[gerr.Code]; ok {
			if v, ok := grpcCodeToDiag[grpcCode]; ok {
				return diag.Diagnostics{
					RedactError(projects,
						diag.NewBaseError(err, v.typ, append(opts,
							diag.WithType(v.typ),
							diag.WithSummary("%s", v.summary),
							diag.WithError(errors.New(gerr.Message)),
							diag.WithSeverity(v.severity),
						)...),
					),
				}
			}
		}
	}

	if es := err.Error(); strings.HasPrefix(es, "google: error getting credentials") || strings.HasPrefix(es, "google: could not find default credentials") {
		return diag.Diagnostics{
			RedactError(projects,
				diag.NewBaseError(err, diag.ACCESS, append(opts, diag.WithDetails("Please see this document for GCP authentication: https://docs.cloudquery.io/docs/getting-started/getting-started-with-gcp#authenticate-with-gcp"))...),
			),
		}
	}

	return diag.Diagnostics{
		RedactError(projects, diag.NewBaseError(err, fallbackType, opts...)),
	}
}

// RedactError redacts a given diagnostic and returns a redacted diagnostic containing both original and redacted versions
func RedactError(projects []string, e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		errors.New(removePII(projects, e.Error())),
		e.Type(),
		diag.WithSeverity(e.Severity()),
		diag.WithResourceName(e.Description().Resource),
		diag.WithSummary("%s", removePII(projects, e.Description().Summary)),
		diag.WithDetails("%s", removePII(projects, e.Description().Detail)),
	)
	return diag.NewRedactedDiagnostic(e, r)
}

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

func statusFromError(err error) (*status.Status, bool) {
	if err == nil {
		return nil, false
	}
	var se interface {
		GRPCStatus() *status.Status
	}
	if errors.As(err, &se) {
		return se.GRPCStatus(), true
	}
	return nil, false
}
