package telemetry

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/afero"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	otrace "go.opentelemetry.io/otel/trace"
)

type Client struct {
	ores     *resource.Resource
	tp       otrace.TracerProvider
	exporter trace.SpanExporter
	closer   io.Closer

	logger hclog.Logger
	fs     afero.Afero
	err    error

	disabled bool
}

type Option func(*Client)

func WithFS(fs afero.Fs) Option {
	return func(c *Client) {
		c.fs = afero.Afero{Fs: fs}
	}
}

func WithLogger(logger hclog.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithResource(res *resource.Resource) Option {
	return func(c *Client) {
		c.ores = res
	}
}

func WithExporter(w io.WriteCloser) Option {
	return func(c *Client) {
		exp, err := stdouttrace.New(
			stdouttrace.WithWriter(w),
			stdouttrace.WithPrettyPrint(),
			stdouttrace.WithoutTimestamps(),
		)
		c.setError(err)
		c.exporter = exp
		c.closer = w
	}
}

func WithDisabled() Option {
	return func(c *Client) {
		c.disabled = true
	}
}

func New(options ...Option) *Client {
	c := &Client{
		fs: afero.Afero{Fs: afero.NewOsFs()},
	}
	for _, opt := range options {
		opt(c)
	}

	if c.ores == nil {
		c.ores, c.err = c.defaultResource()
	}

	if c.logger == nil {
		c.logger = hclog.NewNullLogger()
	}

	opts := []trace.TracerProviderOption{
		trace.WithSampler(trace.AlwaysSample()),
	}
	if c.disabled {
		c.tp = otrace.NewNoopTracerProvider()
		return c
	}

	if c.ores != nil {
		opts = append(opts, trace.WithResource(c.ores))
	}
	if c.exporter != nil {
		opts = append(opts, trace.WithBatcher(c.exporter)) // could consider using trace.WithSyncer instead for sync (and slow) results
	} else {
		// TODO add a "real" networked exporter here
		_ = true
	}

	c.tp = trace.NewTracerProvider(opts...)

	return c
}

func (c *Client) Tracer() otrace.Tracer {
	return c.tp.Tracer("cloudquery.io/internal/telemetry")
}

func (c *Client) Shutdown() {
	if c.disabled {
		return
	}

	if sd, ok := c.tp.(shutdownable); ok {
		if err := sd.Shutdown(context.Background()); err != nil {
			c.logger.Debug("shutdown failed", "error", err)
		}
	}

	if c.closer != nil {
		if err := c.closer.Close(); err != nil {
			c.logger.Debug("close failed", "error", err)
		}
	}
}

func (c *Client) HasError() error {
	return c.err
}

func (c *Client) setError(err error) {
	if c.err == nil && err != nil {
		c.err = err
	}
}

func (c *Client) defaultResource() (*resource.Resource, error) {
	cookieContents, err := c.cookie()
	if err != nil {
		c.logger.Debug("cookie failed", "error", err)
	}

	attr := []attribute.KeyValue{
		semconv.ServiceNameKey.String("cloudquery"),
		semconv.ServiceVersionKey.String("0.0.0"), // TODO insert release version/commit hash/dirty flag
		attribute.Bool("ci", isCI()),
	}
	if cookieContents != "" {
		attr = append(attr, semconv.ServiceInstanceIDKey.String(cookieContents))
	}

	return resource.New(context.Background(),
		resource.WithTelemetrySDK(),
		resource.WithHost(), // exposes hostname
		resource.WithOS(),   // includes os description which has hostname + os version
		resource.WithProcessRuntimeName(),
		resource.WithProcessRuntimeVersion(),
		resource.WithProcessRuntimeDescription(),
		resource.WithAttributes(attr...),
	)
}

func (c *Client) cookie() (string, error) {
	fn := filepath.Join(".", ".cq", "telemetry-cookie")

	exists := true
	fi, err := c.fs.Stat(fn)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
		exists = false
	}
	if exists && fi.IsDir() {
		return "", fmt.Errorf("telemetry-cookie is a directory")
	}

	if exists {
		b, err := c.fs.ReadFile(fn)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	id := uuid.NewV4().String()
	if err := c.fs.WriteFile(fn, []byte(id), 0644); err != nil {
		return "", err
	}
	return id, nil
}

func isCI() bool {
	for _, v := range []string{
		"CI", "BUILD_ID", "BUILDKITE", "CIRCLECI", "CIRCLE_CI", "CIRRUS_CI", "CODEBUILD_BUILD_ID", "GITHUB_ACTIONS", "GITLAB_CI", "HEROKU_TEST_RUN_ID", "TEAMCITY_VERSION", "TF_BUILD", "TRAVIS",
	} {
		if os.Getenv(v) != "" {
			return true
		}
	}

	return false
}

type shutdownable interface {
	Shutdown(context.Context) error
}

var _ shutdownable = (*trace.TracerProvider)(nil)
