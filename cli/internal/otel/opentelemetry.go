package otel

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/rs/zerolog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	otellog "go.opentelemetry.io/otel/log"
	logglobal "go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	def := resource.Default()
	schemaURL := def.SchemaURL()
	if schemaURL == "" {
		schemaURL = semconv.SchemaURL
	}
	r, err := resource.Merge(
		def,
		resource.NewWithAttributes(
			schemaURL,
			semconv.ServiceName("cloudquery-cli"),
		),
	)
	if err != nil {
		panic(err)
	}
	return r
}

type otelConfig struct {
	endpoint string
}

func getLogsProcessor(ctx context.Context, opts otelConfig) (*log.BatchProcessor, error) {
	if opts.endpoint == "" {
		return nil, nil
	}

	logOptions := []otlploghttp.Option{
		otlploghttp.WithEndpoint(opts.endpoint),
		otlploghttp.WithInsecure(),
		otlploghttp.WithCompression(otlploghttp.GzipCompression),
	}

	exporter, err := otlploghttp.New(ctx, logOptions...)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP log exporter: %w", err)
	}

	processor := log.NewBatchProcessor(exporter)
	return processor, nil
}

func SetupOtel(ctx context.Context, logger zerolog.Logger, otelEndpoint string) (shutdown func(), err error) {
	opts := otelConfig{
		endpoint: otelEndpoint,
	}

	logsProcessor, err := getLogsProcessor(ctx, opts)
	if err != nil {
		return nil, err
	}

	pluginResource := newResource()
	lp := log.NewLoggerProvider(
		log.WithProcessor(logsProcessor),
		log.WithResource(pluginResource),
	)

	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		logger.Warn().Err(err).Msg("otel error")
	}))

	logglobal.SetLoggerProvider(lp)

	shutdown = func() {
		if err := lp.Shutdown(context.Background()); err != nil {
			logger.Error().Err(err).Msg("failed to shutdown OTLP logger provider")
		}
	}

	return shutdown, nil
}

type otelLoggerHook struct {
	otellog.Logger
	ctx          context.Context
	invocationId string
}

func (h *otelLoggerHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	record := otellog.Record{}
	record.SetTimestamp(time.Now().UTC())
	record.SetSeverity(otellogSeverity(level))
	record.SetSeverityText(level.String())
	record.SetBody(otellog.StringValue(message))
	// See https://github.com/rs/zerolog/issues/493, this is ugly but it works
	// At the moment there's no way to get the log fields from the event, so we use reflection to get the buffer and parse it
	// TODO: Remove this if https://github.com/rs/zerolog/pull/682 is merged
	logData := make(map[string]any)
	eventBuffer := fmt.Sprintf("%s}", reflect.ValueOf(e).Elem().FieldByName("buf"))
	err := json.Unmarshal([]byte(eventBuffer), &logData)
	if err == nil {
		recordAttributes := make([]otellog.KeyValue, 0, len(logData))
		for k, v := range logData {
			if k == "message" {
				record.SetBody(otellog.StringValue(fmt.Sprintf("%v", v)))
				continue
			}

			if k == "level" {
				continue
			}
			if k == "time" {
				eventTimestamp, ok := v.(string)
				if !ok {
					continue
				}
				t, err := time.Parse(time.RFC3339Nano, eventTimestamp)
				if err == nil {
					record.SetTimestamp(t)
					continue
				}
			}
			var attributeValue otellog.Value
			switch v := v.(type) {
			case string:
				attributeValue = otellog.StringValue(v)
			case int:
				attributeValue = otellog.IntValue(v)
			case int64:
				attributeValue = otellog.Int64Value(v)
			case float64:
				attributeValue = otellog.Float64Value(v)
			case bool:
				attributeValue = otellog.BoolValue(v)
			case []byte:
				attributeValue = otellog.BytesValue(v)
			default:
				attributeValue = otellog.StringValue(fmt.Sprintf("%v", v))
			}
			recordAttributes = append(recordAttributes, otellog.KeyValue{
				Key:   k,
				Value: attributeValue,
			})
		}
		recordAttributes = append(recordAttributes, otellog.KeyValue{
			Key:   "invocation_id",
			Value: otellog.StringValue(h.invocationId)},
		)
		record.AddAttributes(recordAttributes...)
	}

	h.Emit(h.ctx, record)
}

func otellogSeverity(level zerolog.Level) otellog.Severity {
	switch level {
	case zerolog.DebugLevel:
		return otellog.SeverityDebug
	case zerolog.InfoLevel:
		return otellog.SeverityInfo
	case zerolog.WarnLevel:
		return otellog.SeverityWarn
	case zerolog.ErrorLevel:
		return otellog.SeverityError
	case zerolog.FatalLevel:
		return otellog.SeverityFatal2
	case zerolog.PanicLevel:
		return otellog.SeverityFatal1
	case zerolog.TraceLevel:
		return otellog.SeverityTrace
	default:
		return otellog.SeverityUndefined
	}
}

func NewOTELLoggerHook(invocationId string) zerolog.Hook {
	return &otelLoggerHook{logglobal.GetLoggerProvider().Logger("cloudquery"), context.Background(), invocationId}
}
