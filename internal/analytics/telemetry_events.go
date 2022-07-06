package analytics

import "github.com/cloudquery/cq-provider-sdk/provider/diag"

// TelemetryEvent holds data of a telemetry event collected from a Diagnostic.
type TelemetryEvent struct {
	Error    string
	Resource string
	Summary  string
	Category string
}

func (e TelemetryEvent) Properties() map[string]interface{} {
	return map[string]interface{}{
		"error":    e.Error,
		"resource": e.Resource,
		"summary":  e.Summary,
	}
}

// TelemetryFromDiagnostic converts a Diagnostic to a TelemetryEvent. Be sure to pass only Diagnostics of type TELEMETRY.
func TelemetryFromDiagnostic(d diag.Diagnostic) TelemetryEvent {
	desc := d.Description()
	return TelemetryEvent{
		Error:    d.Error(),
		Resource: desc.Resource,
		Summary:  desc.Summary,
		Category: desc.Detail,
	}
}

// FilterTelemetryEvents splits a list of supplied Diagnostics into a list of TelemetryEvents and all other ordinary Diagnostics.
func FilterTelemetryEvents(diags diag.Diagnostics) ([]TelemetryEvent, diag.Diagnostics) {
	filtered := make(diag.Diagnostics, 0, len(diags))
	var events []TelemetryEvent
	for _, d := range diags {
		if d.Type() == diag.TELEMETRY {
			events = append(events, TelemetryFromDiagnostic(d))
		} else {
			filtered = append(filtered, d)
		}
	}
	return events, filtered
}
