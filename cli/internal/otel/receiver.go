package otel

import (
	"context"
	"fmt"
	"net"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
)

type pluginMetric struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Unit        string         `json:"unit"`
	Value       int64          `json:"value"`
	Attributes  map[string]any `json:"attributes"`
}

type tableMetric struct {
	Table string `json:"table"`

	// Deprecated: ClientId is deprecated.
	ClientId string `json:"client_id"`
	// Deprecated: StartTime is deprecated, use Duration instead
	StartTime *time.Time `json:"start_time"`
	// Deprecated: EndTime is deprecated, use Duration instead
	EndTime *time.Time `json:"end_time"`

	Duration  *int64 `json:"duration"`
	Resources int64  `json:"resources"`
	Errors    int64  `json:"errors"`
	Panics    int64  `json:"panics"`
}

type Consumer struct {
	quit            chan any
	metricsFilename string
	metricsFile     *os.File
	consumeMetric   func(context.Context, pluginMetric)
	wg              *sync.WaitGroup
}

func (c Consumer) Shutdown(ctx context.Context) {
	close(c.quit)
	c.wg.Wait()
}

func newMetricConsumer(metricsFile *os.File, quit chan any, wg *sync.WaitGroup) func(context.Context, pluginMetric) {
	tableLock := sync.Mutex{}
	metricsMap := make(map[string]*tableMetric)
	ticker := time.NewTicker(20 * time.Second)

	renderTable := func() {
		tableLock.Lock()
		metrics := maps.Values(metricsMap)
		tableLock.Unlock()
		_, err := metricsFile.Seek(0, 0)
		if err != nil {
			return
		}
		t := table.NewWriter()
		t.SetOutputMirror(metricsFile)
		t.AppendHeader(table.Row{"Table", "Duration", "Resources", "Errors", "Panics"})
		sort.SliceStable(metrics, func(i, j int) bool {
			m1 := metrics[i]
			m2 := metrics[j]

			if m1.EndTime == nil && m2.EndTime != nil {
				return true
			}

			if m1.EndTime != nil && m2.EndTime == nil {
				return false
			}

			return m1.Table+m1.ClientId < m2.Table+m2.ClientId
		})
		for _, metrics := range metrics {
			var duration time.Duration
			switch {
			case metrics.Duration != nil:
				duration = time.Duration(*metrics.Duration * int64(time.Millisecond))
			case metrics.StartTime != nil && metrics.EndTime != nil:
				duration = metrics.EndTime.Sub(*metrics.StartTime)
			case metrics.StartTime != nil:
				duration = time.Since(*metrics.StartTime)
			}
			row := table.Row{
				metrics.Table,
				duration,
				metrics.Resources,
				metrics.Errors,
				metrics.Panics,
			}
			if duration == 0 {
				row[4] = "N/A"
			}
			t.AppendRow(row)
		}
		t.Render()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				renderTable()
			case <-quit:
				renderTable()
				ticker.Stop()
				_ = metricsFile.Close()
				return
			}
		}
	}()

	return func(ctx context.Context, metric pluginMetric) {
		table := metric.Attributes["sync.table.name"].(string)
		clientId, _ := metric.Attributes["sync.table.client_id"].(string)

		tableLock.Lock()
		defer tableLock.Unlock()
		key := table + clientId
		metrics, ok := metricsMap[key]
		if !ok {
			metrics = &tableMetric{
				Table:    table,
				ClientId: clientId,
			}
			metricsMap[key] = metrics
		}

		switch metric.Name {
		case "sync.table.start_time":
			startTime := time.Unix(0, metric.Value)
			metrics.StartTime = &startTime
		case "sync.table.end_time":
			endTime := time.Unix(0, metric.Value)
			metrics.EndTime = &endTime
		case "sync.table.duration":
			metrics.Duration = &metric.Value
		case "sync.table.resources":
			metrics.Resources = metric.Value
		case "sync.table.errors":
			metrics.Errors = metric.Value
		case "sync.table.panics":
			metrics.Panics = metric.Value
		}
	}
}

// Capabilities implements consumer.Traces.
func (Consumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

// ConsumeTraces implements consumer.Traces.
func (Consumer) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	// Do nothing, the CLI only needs metrics to print the table metrics file
	return nil
}

// ConsumeLogs implements consumer.Logs.
func (Consumer) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	// Do nothing, the CLI only needs metrics to print the table metrics file
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
					c.consumeMetric(ctx, pluginMetric{
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

func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

type OtelReceiverOption func(*Consumer)

func WithMetricsFilename(filename string) OtelReceiverOption {
	return func(c *Consumer) {
		c.metricsFilename = filename
	}
}

func StartOtelReceiver(ctx context.Context, opts ...OtelReceiverOption) (*OtelReceiver, error) {
	c := Consumer{
		wg: &sync.WaitGroup{},
	}
	for _, opt := range opts {
		opt(&c)
	}

	metricsFile, err := os.Create(c.metricsFilename)
	if err != nil {
		return nil, err
	}
	quit := make(chan any)
	c.consumeMetric = newMetricConsumer(metricsFile, quit, c.wg)
	c.metricsFile = metricsFile
	c.quit = quit

	factory := otlpreceiver.NewFactory()
	config := factory.CreateDefaultConfig().(*otlpreceiver.Config)
	config.GRPC = configoptional.None[configgrpc.ServerConfig]()
	port, err := getFreePort()
	if err != nil {
		return nil, err
	}
	// See https://github.com/open-telemetry/opentelemetry-collector/blob/b63f70d5ea20e6c664991a535c4fae43ea526491/receiver/otlpreceiver/factory.go#L51
	// `CreateDefaultConfig` creates a default config for HTTP server config, but it's not accessible via config.HTTP.Get()
	// Also see https://github.com/open-telemetry/opentelemetry-collector/issues/13160
	httpConfig := confighttp.NewDefaultServerConfig()
	httpConfig.Endpoint = fmt.Sprintf("localhost:%d", port)
	httpConfig.TLS = configoptional.None[configtls.ServerConfig]()
	config.HTTP = configoptional.Some(otlpreceiver.HTTPConfig{
		ServerConfig:   httpConfig,
		TracesURLPath:  "/v1/traces",
		MetricsURLPath: "/v1/metrics",
		LogsURLPath:    "/v1/logs",
	})

	settings := receiver.Settings{
		ID: component.NewID(component.MustNewType("otlp")),
		TelemetrySettings: component.TelemetrySettings{
			Logger:         zap.NewNop(),
			MeterProvider:  noopmetric.NewMeterProvider(),
			TracerProvider: nooptrace.NewTracerProvider(),
		},
	}

	var components []component.Component
	traces, err := factory.CreateTraces(ctx, settings, config, c)
	if err != nil {
		return nil, err
	}
	components = append(components, traces)

	metrics, err := factory.CreateMetrics(ctx, settings, config, c)
	if err != nil {
		return nil, err
	}
	components = append(components, metrics)

	logs, err := factory.CreateLogs(ctx, settings, config, c)
	if err != nil {
		return nil, err
	}
	components = append(components, logs)

	for _, c := range components {
		if err := c.Start(ctx, nil); err != nil {
			return nil, err
		}
	}

	return &OtelReceiver{
		consumer:   c,
		components: components,
		Endpoint:   httpConfig.Endpoint,
	}, nil
}

func (r *OtelReceiver) Shutdown(ctx context.Context) {
	if r == nil {
		return
	}

	r.consumer.Shutdown(ctx)
	for _, c := range r.components {
		_ = c.Shutdown(ctx)
	}
}
