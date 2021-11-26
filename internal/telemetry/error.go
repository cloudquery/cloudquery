package telemetry

import (
	"fmt"

	"go.opentelemetry.io/otel/codes"
	otrace "go.opentelemetry.io/otel/trace"
)

// RecordError should be called on a span to mark it as errored. By default error values are not recorded, unless the debug flag is set.
func RecordError(span otrace.Span, err error, opts ...otrace.EventOption) {
	if err == nil {
		return
	}

	anon := true
	if d, ok := span.(debugModer); ok && d.DebugMode() {
		anon = false
	}

	//  TODO for fetch get table name / error type

	if anon {
		span.RecordError(fmt.Errorf("error"), opts...)
		span.SetStatus(codes.Error, "error")
		return
	}

	span.RecordError(err, opts...)
	span.SetStatus(codes.Error, err.Error())
}
