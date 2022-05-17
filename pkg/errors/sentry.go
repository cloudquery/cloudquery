package errors

import (
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
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
	allowUnmanaged := core.Version == core.DevelopmentVersion && viper.GetBool("debug-sentry")
	for _, d := range dd.Squash().Redacted() {
		if ShouldIgnoreDiag(d) {
			continue
		}

		if classifyError(d) != errNoClass {
			continue
		}
		sentry.WithScope(func(scope *sentry.Scope) {
			if t, ok := d.(core.FetchDiagnostic); ok {
				if allowUnmanaged && t.Version == plugin.Unmanaged {
					return
				}
				scope.SetTags(map[string]string{"provider": t.Provider, "provider_version": t.Version, "resource": t.Description().Resource})
			}
			if isConfigureDiagnostic(d) {
				scope.SetTag("source", "configure")
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
