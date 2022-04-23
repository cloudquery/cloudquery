package analytics

import (
	"time"

	"github.com/cloudquery/cloudquery/internal/persistentdata"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/google/uuid"
	"github.com/spf13/afero"
	"gopkg.in/segmentio/analytics-go.v3"
)

const CQTeamID = "12345678-0000-0000-0000-c1a0dbeef000"

type VersionInfo struct {
	Version   string
	BuildDate string
	CommitId  string
}

// currentHub is the initial Hub with no Client bound and an empty Scope.
var currentHub = New()

type Client struct {
	version    VersionInfo
	env        *Environment
	terminal   bool
	userId     uuid.UUID
	instanceId uuid.UUID

	disabled         bool
	endpoint         string
	insecureEndpoint bool

	client analytics.Client
}

type Option func(c *Client)

func WithDisabled() Option {
	return func(c *Client) {
		c.disabled = true
	}
}

func WithVersionInfo(version, commit, buildDate string) Option {
	return func(c *Client) {
		c.version.Version = version
		c.version.CommitId = commit
		c.version.BuildDate = buildDate
	}
}

func WithEndpoint(endpoint string, insecure bool) Option {
	return func(c *Client) {
		c.endpoint = endpoint
		c.insecureEndpoint = insecure
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
		userId:     getUserId(),
		instanceId: uuid.New(),
		client:     analytics.New("NaEpIFd1mc3i6IT7Jas66hp170gNbHzE"),
	}
	for _, o := range opts {
		o(c)
	}
	if c.env == nil {
		c.env = getEnvironmentAttributes(c.terminal)
	}
	return c
}

// randomId will read or generate a persistent `telemetry-random-id` file and return its value.
// First it will try reading ~/.cq/telemetry-random-id and use that value if found. If not, it will move on to ./cq/telemetry-random-id, first attempting a read and if not found, will create that file filling it with a newly generated ID.
// If a directory with the same name is encountered, process is aborted and an empty string is returned.
// If a new file is generated, c.newRandomId is set.
func getUserId() uuid.UUID {
	fs := afero.Afero{Fs: afero.NewOsFs()}
	v, err := persistentdata.New(fs, "telemetry-random-id", uuid.NewString).Get()
	if err != nil {
		return uuid.New()
	}
	id, err := uuid.Parse(v.Content)
	if err != nil {
		return uuid.New()
	}
	return id
}

type Message interface {
	Properties() map[string]interface{}
}

func Capture(eventType string, providers registry.Providers, data Message, diags diag.Diagnostics) error {
	c := currentHub
	if c.disabled {
		return nil
	}

	return c.client.Enqueue(analytics.Track{
		UserId:    c.userId.String(),
		Event:     eventType,
		Timestamp: time.Now().UTC(),
		Context:   nil,
		Properties: map[string]interface{}{
			"version":     c.version.Version,
			"commitId":    c.version.CommitId,
			"buildDate":   c.version.BuildDate,
			"env":         c.env,
			"data":        data.Properties(),
			"instanceId":  c.instanceId,
			"success":     !diags.HasErrors(),
			"providers":   providers,
			"diagnostics": core.SummarizeDiagnostics(diags),
		},
	})
}

func Close() {
	_ = currentHub.client.Close()
}
