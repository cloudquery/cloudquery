package telemetry

import (
	"github.com/getsentry/sentry-go"
	"go.opentelemetry.io/otel/codes"
	otrace "go.opentelemetry.io/otel/trace"
)

// RecordError should be called on a span to mark it as errored. By default error values are not recorded, unless the debug flag is set.
func RecordError(span otrace.Span, err error, opts ...otrace.EventOption) {
	if err == nil {
		return
	}

	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetFingerprint([]string{span.SpanContext().TraceID().String()})
		sentry.CaptureException(err)
	})

	span.RecordError(err, opts...)
	span.SetStatus(codes.Error, err.Error())
}
