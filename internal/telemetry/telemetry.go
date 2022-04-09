package telemetry

import (
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/afero"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	otrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/cloudquery/cloudquery/internal/persistentdata"
	"github.com/cloudquery/cloudquery/pkg/ui"
)

const timeout = 2 * time.Second

// Client is the telemetry client.
type Client struct {
	// OpenTelemetry resource entry. Used in optional args.
	ores *resource.Resource

	// The TracerProvider we create
	tp otrace.TracerProvider

	// Exporter to use. Used in optional args.
	exporter trace.SpanExporter

	// This is closed on shutdown. Used with the file exporter to close the file.
	closer io.Closer

	logger hclog.Logger
	fs     afero.Afero

	// If we encountered an error during Telemetry init, it's set here.
	err error

	// Build info. These are set as resource attributes in the default resource.
	version, commit, buildDate string

	// Whether we're in debug mode or not. In debug mode, error messages from the OpenTelemetry SDK is bumped to a higher level
	debug bool

	// Whether telemetry collection is disabled. If so, a NoopTracerProvider is set, and we don't initialize the default resource
	disabled bool

	// true if we created a new telemetry-random-id file. This is used to enable the warning message in the console client.
	newRandomId bool

	// Contents of the generated/persisted random id
	randomIdValue string

	// endpoint to send data to
	endpoint string

	// allow insecure connection to endpoint
	insecureEndpoint bool
}

type Option func(*Client)

const CQTeamID = "12345678-0000-0000-0000-c1a0dbeef000"

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
		)
		c.setError(err)
		c.exporter = exp
		c.closer = w
	}
}

func WithDebug() Option {
	return func(c *Client) {
		c.debug = true
	}
}

func WithDisabled() Option {
	return func(c *Client) {
		c.disabled = true
	}
}

func WithVersionInfo(version, commit, buildDate string) Option {
	return func(c *Client) {
		c.version = version
		c.commit = commit
		c.buildDate = buildDate
	}
}

func WithEndpoint(endpoint string, insecure bool) Option {
	return func(c *Client) {
		c.endpoint = endpoint
		c.insecureEndpoint = insecure
	}
}

func New(ctx context.Context, options ...Option) *Client {
	c := &Client{
		fs:     afero.Afero{Fs: afero.NewOsFs()},
		logger: hclog.NewNullLogger(),
	}
	for _, opt := range options {
		opt(c)
	}

	otel.SetErrorHandler(&errorHandler{l: c.logger, debug: c.debug})

	if c.ores == nil {
		var err error
		c.ores, err = c.defaultResource(ctx)
		c.setError(err)
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
		exp, err := c.httpExporter(ctx)
		if err != nil {
			c.setError(err)
		} else {
			opts = append(opts, trace.WithBatcher(exp))
		}
	}

	c.tp = trace.NewTracerProvider(opts...)

	return c
}

func (c *Client) Tracer(ctx context.Context) (context.Context, Tracer) {
	tw := &wrappedTracer{
		Tracer: c.tp.Tracer("cloudquery.io/internal/telemetry"),
	}
	return ContextWithTracer(ctx, tw), tw
}

func (c *Client) Shutdown(ctx context.Context) {
	if err := c.HasError(); err != nil {
		c.logger.Debug("telemetry error", "error", err)
	}

	if c.disabled {
		return
	}

	if sd, ok := c.tp.(shutdownable); ok {
		if err := sd.Shutdown(ctx); err != nil {
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

func (c *Client) Enabled() bool {
	return !c.disabled
}

func (c *Client) setError(err error) {
	if err != nil {
		c.logger.Debug("telemetry error occurred", "error", err)
	}
	if c.err == nil && err != nil {
		c.err = err
	}
}

func (c *Client) defaultResource(ctx context.Context) (*resource.Resource, error) {
	randId, err := c.randomId()
	if err != nil {
		c.logger.Debug("randomId failed", "error", err)
	}

	attr := []attribute.KeyValue{
		semconv.ServiceNameKey.String("cloudquery"),
		semconv.ServiceVersionKey.String(c.version),
		attribute.String("commit", c.commit),
		attribute.String("build_date", c.buildDate),
		attribute.Bool("ci", IsCI()),
		attribute.Bool("faas", IsFaaS()),
		attribute.Bool("terminal", ui.IsTerminal()),
	}
	if !c.newRandomId && randId != "" {
		attr = append(attr, attribute.Bool("random_id_persisted", true))
	}
	if randId == "" {
		randId = genRandomId() // generate ephemeral random ID on error
	}
	c.randomIdValue = randId

	attr = append(attr, semconv.ServiceInstanceIDKey.String(randId))

	if hn, err := os.Hostname(); err == nil && hn != "" {
		attr = append(attr, semconv.HostNameKey.String(HashAttribute(hn)))
	}
	attr = append(attr, osInfo()...)
	attr = append(attr, macHost()...)

	return resource.New(ctx,
		resource.WithTelemetrySDK(),
		resource.WithProcessRuntimeName(),
		resource.WithProcessRuntimeVersion(),
		resource.WithProcessRuntimeDescription(),
		resource.WithAttributes(attr...),
	)
}

// randomId will read or generate a persistent `telemetry-random-id` file and return its value.
// First it will try reading ~/.cq/telemetry-random-id and use that value if found. If not, it will move on to ./cq/telemetry-random-id, first attempting a read and if not found, will create that file filling it with a newly generated ID.
// If a directory with the same name is encountered, process is aborted and an empty string is returned.
// If a new file is generated, c.newRandomId is set.
func (c *Client) randomId() (string, error) {
	v, err := persistentdata.New(c.fs, "telemetry-random-id", genRandomId).Get()
	c.newRandomId = v.Created
	return v.Content, err
}

// NewRandomId returns true if we created a new random id in this session
func (c *Client) NewRandomId() bool {
	return c.newRandomId
}

// RandomId returns the generated ID
func (c *Client) RandomId() string {
	return c.randomIdValue
}

// httpExporter creates the default HTTP SpanExporter
func (c *Client) httpExporter(ctx context.Context) (trace.SpanExporter, error) {
	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(c.endpoint),
		otlptracehttp.WithTimeout(timeout),
	}
	if c.insecureEndpoint {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	return otlptracehttp.New(ctx, opts...)
}

// grpcExporter creates the default gRPC SpanExporter
func (c *Client) grpcExporter(ctx context.Context) (trace.SpanExporter, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(c.endpoint),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
		otlptracegrpc.WithDialOption(grpc.WithContextDialer(func(_ context.Context, addr string) (net.Conn, error) {
			return net.DialTimeout("tcp", addr, timeout)
		})),
		// otlptracegrpc.WithDialOption(grpc.WithReturnConnectionError()),
		// otlptracegrpc.WithDialOption(grpc.FailOnNonTempDialError(true)),
		otlptracegrpc.WithTimeout(timeout),
	}
	if c.insecureEndpoint {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	return otlptracegrpc.New(ctx, opts...)
}

// IsCI determines if we're running under a CI env by checking CI-specific env vars
func IsCI() bool {
	for _, v := range []string{
		"CI", "BUILD_ID", "BUILDKITE", "CIRCLECI", "CIRCLE_CI", "CIRRUS_CI", "CODEBUILD_BUILD_ID", "GITHUB_ACTIONS", "GITLAB_CI", "HEROKU_TEST_RUN_ID", "TEAMCITY_VERSION", "TF_BUILD", "TRAVIS",
	} {
		if os.Getenv(v) != "" {
			return true
		}
	}

	return false
}

// IsFaaS determines if we're running under a Lambda env by checking Lambda-specific env vars
func IsFaaS() bool {
	for _, v := range []string{
		"LAMBDA_TASK_ROOT", "AWS_LAMBDA_FUNCTION_NAME", // AWS
		"FUNCTION_TARGET",             // GCP
		"AZURE_FUNCTIONS_ENVIRONMENT", // Azure
	} {
		if os.Getenv(v) != "" {
			return true
		}
	}

	return false
}

// HashAttribute creates a one-way hash from an attribute
func HashAttribute(value string) string {
	s := sha1.New()
	_, _ = s.Write([]byte(value))
	return fmt.Sprintf("%0x", s.Sum(nil))
}

type shutdownable interface {
	Shutdown(context.Context) error
}

// Make sure TracerProvider is shutdownable. This would fail if the OpenTelemetry API changes.
// Client.Shutdown() doesn't require this to be the case, because it has to work with NoopTracerProvider as well, which doesn't have a Shutdown method.
var _ shutdownable = (*trace.TracerProvider)(nil)

// errorHandler is used to set the global OpenTelemetry error handler and suppress otel errors to debug level.
type errorHandler struct {
	l     hclog.Logger
	debug bool
}

func (e *errorHandler) Handle(err error) {
	if e.debug {
		// Upgrade error severity
		e.l.Warn("otel error occurred", "error", err)
	} else {
		e.l.Debug("otel error occurred", "error", err)
	}
}

func genRandomId() string {
	return uuid.NewString()
}
