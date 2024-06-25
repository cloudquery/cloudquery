package otel

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"
)

type Span struct {
	TraceID        string         `json:"trace_id"`
	SpanID         string         `json:"span_id"`
	TraceState     string         `json:"trace_state"`
	ParentSpanID   string         `json:"parent_span_id"`
	Name           string         `json:"name"`
	StartTimestamp time.Time      `json:"start_timestamp"`
	EndTimestamp   time.Time      `json:"end_timestamp"`
	Attributes     map[string]any `json:"attributes"`
}

type Metric struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Unit        string         `json:"unit"`
	Value       int64          `json:"value"`
	Attributes  map[string]any `json:"attributes"`
}

type ConsumeSpan = func(context.Context, Span)

type ConsumeMetric = func(context.Context, Metric)

type Consumer struct {
	tracesFile    *os.File
	consumeSpan   ConsumeSpan
	metricsFile   *os.File
	consumeMetric ConsumeMetric
}

func (c *Consumer) Shutdown(ctx context.Context) error {
	var err error
	if c.tracesFile != nil {
		err = errors.Join(err, c.tracesFile.Close())
	}
	if c.metricsFile != nil {
		err = errors.Join(err, c.metricsFile.Close())
	}
	return err
}

func newDefaultSpanConsumer(tracesFile *os.File) ConsumeSpan {
	return func(ctx context.Context, span Span) {
		asJson, err := json.Marshal(span)
		if err != nil {
			return
		}
		_, _ = tracesFile.Write(append(asJson, '\n'))
	}
}

func newDefaultMetricConsumer(metricsFile *os.File) ConsumeMetric {
	return func(ctx context.Context, metric Metric) {
		// write metric to file as json
		asJson, err := json.Marshal(metric)
		if err != nil {
			return
		}
		_, _ = metricsFile.Write(append(asJson, '\n'))
	}
}

// Capabilities implements consumer.Traces.
func (Consumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

// ConsumeTraces implements consumer.Traces.
func (c Consumer) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	resourceSpans := td.ResourceSpans()
	for i := 0; i < resourceSpans.Len(); i++ {
		resourceSpan := resourceSpans.At(i)
		scopedSpans := resourceSpan.ScopeSpans()
		for j := 0; j < scopedSpans.Len(); j++ {
			spans := scopedSpans.At(j).Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				traceID := span.TraceID()
				spanID := span.SpanID()
				parentSpanID := span.ParentSpanID()
				c.consumeSpan(ctx, Span{
					TraceID:        hex.EncodeToString(traceID[:]),
					SpanID:         hex.EncodeToString(spanID[:]),
					ParentSpanID:   hex.EncodeToString(parentSpanID[:]),
					TraceState:     span.TraceState().AsRaw(),
					Name:           span.Name(),
					StartTimestamp: span.StartTimestamp().AsTime(),
					EndTimestamp:   span.EndTimestamp().AsTime(),
					Attributes:     span.Attributes().AsRaw(),
				})
			}
		}
	}
	return nil
}

// ConsumeMetrics implements consumer.Metrics.
func (c Consumer) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	resourceMetrics := md.ResourceMetrics()
	for i := 0; i < resourceMetrics.Len(); i++ {
		resourceMetric := resourceMetrics.At(i)
		scopeMetrics := resourceMetric.ScopeMetrics()
		for j := 0; j < scopeMetrics.Len(); j++ {
			metrics := scopeMetrics.At(j).Metrics()
			for k := 0; k < metrics.Len(); k++ {
				metric := metrics.At(k)
				if metric.Type() != pmetric.MetricTypeSum {
					continue
				}
				dataPoints := metric.Sum().DataPoints()
				for l := 0; l < dataPoints.Len(); l++ {
					dataPoint := dataPoints.At(l)
					c.consumeMetric(ctx, Metric{
						Name:        metric.Name(),
						Description: metric.Description(),
						Unit:        metric.Unit(),
						Value:       dataPoint.IntValue(),
						Attributes:  dataPoint.Attributes().AsRaw(),
					})
				}
			}
		}
	}
	return nil
}

type OtelReceiver struct {
	consumer   Consumer
	components []component.Component
	Endpoint   string
}

func getPort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

type OtelReceiverOption func(*Consumer)

func WithSpanConsumer(consumerSpan ConsumeSpan) OtelReceiverOption {
	return func(c *Consumer) {
		c.consumeSpan = consumerSpan
	}
}

func WithMetricConsumer(consumerMetric ConsumeMetric) OtelReceiverOption {
	return func(c *Consumer) {
		c.consumeMetric = consumerMetric
	}
}

func StartOtelReceiver(ctx context.Context, opts ...OtelReceiverOption) (*OtelReceiver, error) {
	c := Consumer{}
	for _, opt := range opts {
		opt(&c)
	}

	if c.consumeSpan == nil {
		tracesFile, err := os.OpenFile("traces.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		c.consumeSpan = newDefaultSpanConsumer(tracesFile)
		c.tracesFile = tracesFile
	}

	if c.consumeMetric == nil {
		metricsFile, err := os.OpenFile("metrics.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		c.consumeMetric = newDefaultMetricConsumer(metricsFile)
		c.metricsFile = metricsFile
	}

	factory := otlpreceiver.NewFactory()
	config := factory.CreateDefaultConfig().(*otlpreceiver.Config)
	config.GRPC = nil
	// This will list on the first available port
	port, err := getPort()
	if err != nil {
		return nil, err
	}
	config.HTTP.Endpoint = fmt.Sprintf("localhost:%d", port)

	settings := receiver.Settings{
		TelemetrySettings: component.TelemetrySettings{
			Logger:         zap.NewNop(),
			MeterProvider:  noopmetric.NewMeterProvider(),
			TracerProvider: nooptrace.NewTracerProvider(),
			ReportStatus:   func(*component.StatusEvent) {},
		},
	}

	var components []component.Component
	if c.consumeSpan != nil {
		traces, err := factory.CreateTracesReceiver(ctx, settings, config, c)
		if err != nil {
			return nil, err
		}
		components = append(components, traces)
	}

	if c.consumeMetric != nil {
		metrics, err := factory.CreateMetricsReceiver(ctx, settings, config, c)
		if err != nil {
			return nil, err
		}
		components = append(components, metrics)
	}

	for _, c := range components {
		if err := c.Start(ctx, nil); err != nil {
			return nil, err
		}
	}

	return &OtelReceiver{
		consumer:   c,
		components: components,
		Endpoint:   config.HTTP.Endpoint,
	}, nil
}

func (r *OtelReceiver) Shutdown(ctx context.Context) {
	if r == nil {
		return
	}

	_ = r.consumer.Shutdown(ctx)
	for _, c := range r.components {
		_ = c.Shutdown(ctx)
	}
}
