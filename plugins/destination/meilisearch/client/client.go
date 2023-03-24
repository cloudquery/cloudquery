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

	major, _, ok := strings.Cut(version.PkgVersion, ".")
	if !ok {
		return fmt.Errorf("malformed version %q (expected \"major.minor.patch\"", version)
	}

	num, err := strconv.ParseInt(major, 10, 32)
	if err != nil {
		return fmt.Errorf("failed to parse %q as major version: %w", major, err)
	}

	const minVersion = 1
	if num < minVersion {
		return fmt.Errorf("unsupported Meilisearch version %d (must be >= %d)", num, minVersion)
	}

	return nil
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
