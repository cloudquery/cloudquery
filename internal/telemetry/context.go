package telemetry

import (
	"context"

	otrace "go.opentelemetry.io/otel/trace"
)

type tracerContextKeyType int

const currentTracerKey tracerContextKeyType = iota

// TracerFromContext returns the current Tracer from ctx.
//
// If no Tracer is currently set in ctx, a noopTracer is returned.
func TracerFromContext(ctx context.Context) otrace.Tracer {
	if ctx == nil {
		return otrace.NewNoopTracerProvider().Tracer("")
	}
	if tracer, ok := ctx.Value(currentTracerKey).(otrace.Tracer); ok {
		return tracer
	}
	return otrace.NewNoopTracerProvider().Tracer("")

}

// ContextWithTracer returns a copy of parent with tracer set as the current Tracer.
func ContextWithTracer(parent context.Context, tracer otrace.Tracer) context.Context {
	return context.WithValue(parent, currentTracerKey, tracer)
}
