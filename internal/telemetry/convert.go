package telemetry

import "go.opentelemetry.io/otel/attribute"

func MapToAttributes(m map[string]int64) []attribute.KeyValue {
	ret := make([]attribute.KeyValue, 0, len(m))
	for k, v := range m {
		ret = append(ret, attribute.Int64(k, v))
	}
	return ret
}
