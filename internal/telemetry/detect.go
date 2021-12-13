package telemetry

import (
	"context"
	"net"
	"os"
	"sort"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// osInfo will extract OS info using the opentelemetry helpers but remove the hostname from "uname"
func osInfo() []attribute.KeyValue {
	r, err := resource.New(context.Background(), resource.WithOS())
	if err != nil {
		return nil
	}

	hn, err := os.Hostname()
	if err != nil || hn == "" {
		return nil
	}

	attrs := r.Attributes()

	ret := make([]attribute.KeyValue, 0, len(attrs))
	for _, a := range attrs {
		switch a.Key {
		case semconv.OSDescriptionKey:
			parts := strings.SplitN(a.Value.AsString(), " ", 6)
			if len(parts) < 5 || parts[4] != hn {
				continue // skip attribute
			}

			parts[4] = "host"
			ret = append(ret, a.Key.String(strings.Join(parts, " ")))
		case semconv.OSNameKey, semconv.OSVersionKey, semconv.OSTypeKey:
			a := a
			ret = append(ret, a)
		default:
			// skip attribute
		}
	}

	return ret
}

// macHost will extract MAC addresses, add the hostname and return a hash
func macHost() []attribute.KeyValue {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil
	}

	as := make([]string, 0, len(ifas)+1)
	for _, ifa := range ifas {
		if a := ifa.HardwareAddr.String(); a != "" {
			as = append(as, a)
		}
	}

	sort.Strings(as)

	if hn, err := os.Hostname(); err == nil && hn != "" {
		as = append(as, hn)
	}

	return []attribute.KeyValue{
		attribute.String("cq.machost", hashAttribute(strings.Join(as, ","))),
	}
}
