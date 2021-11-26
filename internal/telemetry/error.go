package telemetry

import (
	"fmt"

	"go.opentelemetry.io/otel/codes"
	otrace "go.opentelemetry.io/otel/trace"
)

// RecordAnonError should be called on a span to mark it as errored. Errors values are not recorded.
func RecordAnonError(span otrace.Span, err error, opts ...otrace.EventOption) {
	if err == nil {
		return
	}

	//  TODO for fetch get table name / error type

	span.RecordError(fmt.Errorf("error"), opts...)
	span.SetStatus(codes.Error, "error")
}

// RecordError should be called on a span to mark it as errored. Errors values are recorded.
func RecordError(span otrace.Span, err error, opts ...otrace.EventOption) {
	if err == nil {
		return
	}

	//  TODO for fetch get table name / error type

	span.RecordError(err, opts...)
	span.SetStatus(codes.Error, err.Error())
}
