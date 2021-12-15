package telemetry

import (
	"context"

	otrace "go.opentelemetry.io/otel/trace"
)

type tracerContextKeyType int

const currentTracerKey tracerContextKeyType = iota

type Tracer interface {
	otrace.Tracer
}

// TracerFromContext returns the current Tracer from ctx.
//
// If no Tracer is currently set in ctx, a noopTracer is returned.
func TracerFromContext(ctx context.Context) Tracer {
	if ctx != nil {
		if tracer, ok := ctx.Value(currentTracerKey).(Tracer); ok {
			return tracer
		}
	}

	return &wrappedTracer{
		Tracer: otrace.NewNoopTracerProvider().Tracer(""),
	}
}

// ContextWithTracer returns a copy of given context with tracer set as the current Tracer.
func ContextWithTracer(parent context.Context, tracer Tracer) context.Context {
	return context.WithValue(parent, currentTracerKey, tracer)
}

// SpanCloser gets an error and options and ends the span its attached to
type SpanCloser func(error, ...otrace.SpanEndOption)

// StartSpanFromContext starts the span from a given context with the given options, and returns a new context with span attached, as well as a closer fn.
// Returned SpanCloser should be called when done with span. To catch panics, it should be used under a defer.
func StartSpanFromContext(ctx context.Context, spanName string, opts ...otrace.SpanStartOption) (context.Context, SpanCloser) {
	ktx, span := TracerFromContext(ctx).Start(ctx, spanName, opts...)
	return ktx, func(err error, opts ...otrace.SpanEndOption) {
		RecordError(span, err)
		span.End(opts...)
	}
}

// wrappedTracer is a standard tracer for CQ
type wrappedTracer struct {
	otrace.Tracer
}

// Start calls the parent Start method, then wraps the Span with the debug flag
func (t *wrappedTracer) Start(ctx context.Context, spanName string, opts ...otrace.SpanStartOption) (context.Context, otrace.Span) {
	ctx, s := t.Tracer.Start(ctx, spanName, opts...)
	return ctx, &wrappedSpan{Span: s}
}

var _ Tracer = (*wrappedTracer)(nil)

type wrappedSpan struct {
	otrace.Span
}
