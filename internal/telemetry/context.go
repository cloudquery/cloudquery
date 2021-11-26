package telemetry

import (
	"context"

	otrace "go.opentelemetry.io/otel/trace"
)

type tracerContextKeyType int

const currentTracerKey tracerContextKeyType = iota

type Tracer interface {
	otrace.Tracer
	DebugMode() bool
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

// wrappedTracer is a standard tracer with the debug flag persisted from telemetry.Client
type wrappedTracer struct {
	otrace.Tracer
	debug bool
}

func (t *wrappedTracer) DebugMode() bool {
	return t.debug
}

// Start calls the parent Start method, then wraps the Span with the debug flag
func (t *wrappedTracer) Start(ctx context.Context, spanName string, opts ...otrace.SpanStartOption) (context.Context, otrace.Span) {
	ctx, s := t.Tracer.Start(ctx, spanName, opts...)
	return ctx, &wrappedSpan{Span: s, debug: t.debug}
}

var _ Tracer = (*wrappedTracer)(nil)

type wrappedSpan struct {
	otrace.Span
	debug bool
}

func (s *wrappedSpan) DebugMode() bool {
	return s.debug
}

type debugModer interface {
	DebugMode() bool
}

var (
	_ debugModer = (*wrappedSpan)(nil)
	_ debugModer = (*wrappedTracer)(nil)
)
