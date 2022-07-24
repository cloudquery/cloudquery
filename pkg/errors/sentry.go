package errors

import (
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
