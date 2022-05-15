package analytics

import (
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/persistentdata"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/google/uuid"
	"github.com/modern-go/reflect2"
	"github.com/rs/zerolog/log"
	"github.com/rudderlabs/analytics-go"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
)

const (
	CQTeamID = "12345678-0000-0000-0000-c1a0dbeef000"
)

type VersionInfo struct {
	Version   string `json:"version,omitempty"`
	BuildDate string `json:"build_date,omitempty"`
	CommitId  string `json:"commit_id,omitempty"`
}

// currentHub is the initial Hub with no Client bound and an empty Scope.
var currentHub = New()

type Client struct {
	version    VersionInfo
	env        *Environment
	terminal   bool
	userId     uuid.UUID
	instanceId string

	disabled bool
	debug    bool
	inspect  bool

	properties map[string]interface{}

	client analytics.Client
	apikey string
}

type Option func(c *Client)

func WithProperties(properties map[string]interface{}) Option {
	return func(c *Client) {
		for k, v := range properties {
			c.properties[k] = v
		}
	}
}

func WithApiKey(apikey string) Option {
	return func(c *Client) {
		c.apikey = apikey
	}
}

func WithInstanceId(instanceId string) Option {
	return func(c *Client) {
		c.instanceId = instanceId
	}
}

func WithDisabled() Option {
	return func(c *Client) {
		c.disabled = true
	}
}

func WithDebug() Option {
	return func(c *Client) {
		c.debug = true
	}
}

func WithInspect() Option {
	return func(c *Client) {
		c.inspect = true
	}
}

func WithVersionInfo(version, commit, buildDate string) Option {
	return func(c *Client) {
		c.version.Version = version
		c.version.CommitId = commit
		c.version.BuildDate = buildDate
	}
}

func WithTerminal(terminal bool) Option {
	return func(c *Client) {
		c.terminal = terminal
	}
}

// Init initializes the Analytics Client with options. The returned error is non-nil if
// options is invalid, for instance if a malformed DSN is provided.
func Init(opts ...Option) error {
	currentHub = New(opts...)
	return nil
}

func New(opts ...Option) *Client {
	c := &Client{
		version:    VersionInfo{},
		userId:     GetUserId(),
		instanceId: uuid.New().String(),
		properties: make(map[string]interface{}),
		debug:      false,
		inspect:    false,
	}

	for _, o := range opts {
		o(c)
	}
	if c.env == nil {
		c.env = getEnvironmentAttributes(c.terminal)
	}
	cfg := analytics.Config{}
	if c.debug {
		cfg.Verbose = true
		cfg.Logger = logging.NewSimple(&log.Logger, "analytics")
	}

	ac, err := analytics.NewWithConfig(c.apikey, "https://cloudquerypgm.dataplane.rudderstack.com", cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to initialize analytics client, client is disabled")
		c.disabled = true
	}
	c.client = ac
	return c
}

// GetUserId will read or generate a persistent `telemetry-random-id` file and return its value.
// First it will try reading ~/.cq/telemetry-random-id and use that value if found. If not, it will move on to ./cq/telemetry-random-id, first attempting a read and if not found, will create that file filling it with a newly generated ID.
// If a directory with the same name is encountered, process is aborted and an empty string is returned.
// If a new file is generated, c.newRandomId is set.
func GetUserId() uuid.UUID {
	fs := afero.Afero{Fs: afero.NewOsFs()}
	v, err := persistentdata.New(fs, "telemetry-random-id", uuid.NewString).Get()
	if err != nil {
		return uuid.New()
	}
	id, err := uuid.Parse(strings.TrimSuffix(v.Content, "\r\n"))
	if err != nil {
		return uuid.New()
	}
	return id
}

type Message interface {
	Properties() map[string]interface{}
}

func Capture(eventType string, providers registry.Providers, data Message, diags diag.Diagnostics, extra ...interface{}) {
	c := currentHub
	if c.disabled || c.apikey == "" {
		return
	}
	pp := make([]string, len(providers))
	for i, p := range providers {
		pp[i] = p.String()
	}

	eventProps := map[string]interface{}{
		"version":             c.version.Version,
		"commit_id":           c.version.CommitId,
		"build_date":          c.version.BuildDate,
		"env":                 c.env,
		"instance_id":         c.instanceId,
		"success":             !diags.HasErrors(),
		"installed_providers": pp,
		"diagnostics":         core.SummarizeDiagnostics(diags),
	}

	if !reflect2.IsNil(data) {
		for k, v := range data.Properties() {
			eventProps[k] = v
		}
	}

	for i := 0; i < len(extra); i += 2 {
		eventProps[cast.ToString(extra[i])] = extra[i+1]
	}

	// add any global properties
	for k, v := range c.properties {
		eventProps[k] = v
	}
	if c.inspect {
		log.Info().Interface("data", eventProps).Str("event", eventType).Msg("inspect analytics event")
		// if inspect is turned on we only return inspect messages
		return
	}

	event := analytics.Track{UserId: c.userId.String(), Event: eventType, Timestamp: time.Now().UTC(), Properties: eventProps}
	if err := c.client.Enqueue(event); err != nil {
		if c.debug {
			log.Error().Err(err).Msg("failed to send analytics")
		}
	}
}

func SetGlobalProperty(k string, v interface{}) {
	c := currentHub
	c.properties[k] = v
}

func Enabled() bool {
	return !currentHub.disabled
}

func Close() {
	if currentHub.client == nil {
		return
	}
	_ = currentHub.client.Close()
}
