package analytics

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
