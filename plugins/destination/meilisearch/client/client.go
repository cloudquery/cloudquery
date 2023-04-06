package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/zerolog"
)

type Client struct {
	Meilisearch *meilisearch.Client

	logger   zerolog.Logger
	dstSpec  specs.Destination
	pkColumn string

	destination.UnimplementedUnmanagedWriter
}

var _ destination.Client = (*Client)(nil)

func (c *Client) Close(context.Context) error {
	c.Meilisearch = nil
	return nil
}

func (*Client) DeleteStale(context.Context, schema.Tables, string, time.Time) error {
	return fmt.Errorf("DeleteStale not supported")
}

func (c *Client) verifyVersion() error {
	version, err := c.Meilisearch.Version()
	if err != nil {
		return err
	}

	parts := strings.Split(version.PkgVersion, ".")
	if len(parts) < 3 {
		return fmt.Errorf("malformed version %q (expected \"major.minor.patch\"", version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse major version (%q): %w", parts[0], err)
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("failed to parse minor version (%q): %w", parts[1], err)
	}

	const (
		minMajor = 1
		minMinor = 1
	)

	if (major > minMajor) || (major == minMajor && minor >= minMinor) {
		return nil
	}

	return fmt.Errorf("unsupported Meilisearch version %s (must be >= 1.1)", version.PkgVersion)
}

func New(_ context.Context, logger zerolog.Logger, dstSpec specs.Destination) (destination.Client, error) {
	var pkColumn string
	switch dstSpec.WriteMode {
	case specs.WriteModeAppend:
		pkColumn = schema.CqIDColumn.Name
	case specs.WriteModeOverwrite:
		pkColumn = hashColumnName
	default:
		return nil, fmt.Errorf("%q write_mode is not supported", dstSpec.WriteMode)
	}

	spec := new(Spec)
	if err := dstSpec.UnmarshalSpec(spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	spec.setDefaults()
	if err := spec.validate(); err != nil {
		return nil, err
	}

	mClient, err := spec.getClient()
	if err != nil {
		return nil, err
	}

	client := &Client{
		Meilisearch: mClient,
		logger:      logger.With().Str("module", "dest-meilisearch").Str("host", spec.Host).Logger(),
		dstSpec:     dstSpec,
		pkColumn:    pkColumn,
	}

	return client, client.verifyVersion()
}
