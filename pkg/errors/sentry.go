package errors

import (
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/getsentry/sentry-go"
)

func CaptureError(err error, tags map[string]string) {
	if err == nil {
		return
	}
	if classifyError(err) != errNoClass {
		return
	}
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetTags(tags)
		sentry.CaptureException(err)
	})
}

func CaptureDiagnostics(dd diag.Diagnostics, tags map[string]string) {
	for _, d := range dd.Squash().Redacted() {
		if ShouldIgnoreDiag(d) {
			continue
		}

		if classifyError(d) != errNoClass {
			continue
		}
		sentry.WithScope(func(scope *sentry.Scope) {
			if ok, tags, ignore := isSentryDiagnostic(d); ok {
				if ignore {
					return
				}
				scope.SetTags(tags)
			}
			// set any extra tags to this scope
			scope.SetTags(tags)
			scope.SetTags(map[string]string{"diag_type": d.Type().String()})
			scope.SetExtra("detail", d.Description().Detail)

			switch d.Severity() {
			case diag.IGNORE:
				scope.SetLevel(sentry.LevelDebug)
			case diag.WARNING:
				scope.SetLevel(sentry.LevelWarning)
			case diag.PANIC:
				scope.SetLevel(sentry.LevelFatal)
			}
			sentry.CaptureException(d)
		})
	}
}
