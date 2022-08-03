package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

// TODO - delete

func ErrorClassifier(_ schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	return classifyError(err, diag.RESOLVING, diag.WithResourceName(resourceName))
}

func classifyError(err error, fallbackType diag.Type, opts ...diag.BaseErrorOption) diag.Diagnostics {
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(diag.NewBaseError(d, d.Type(), opts...)),
		}
	}

	return diag.Diagnostics{
		RedactError(diag.NewBaseError(err, fallbackType, opts...)),
	}
}

// RedactError redacts a given diagnostic and returns a RedactedDiagnostic containing both original and redacted versions
func RedactError(e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		nil,
		e.Type(),
		diag.WithSeverity(e.Severity()),
		diag.WithResourceName(e.Description().Resource),
		diag.WithSummary("%s", removePII(e.Description().Summary)),
		diag.WithDetails("%s", removePII(e.Description().Detail)),
	)
	return diag.NewRedactedDiagnostic(e, r)
}

func removePII(msg string) string {
	return msg
}
