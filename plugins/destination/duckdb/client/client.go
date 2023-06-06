package client

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/rs/zerolog"

	// import duckdb driver
	"github.com/marcboeker/go-duckdb"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	db        *sql.DB
	connector driver.Connector
	logger    zerolog.Logger
	spec      specs.Destination
	duckSpec  Spec
	metrics   destination.Metrics
}

var _ destination.Client = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, dstSpec specs.Destination) (destination.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "duckdb-dest").Logger(),
		spec:   dstSpec,
	}

	var spec Spec
	if err := dstSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal duckdb spec: %w", err)
	}

	c.duckSpec = spec

	// The connection string parsing here is a temporary workaround for external
	// databases like MotherDuck that can fail to close ATTACH-ed databases and
	// need to first connect with a label and then issue a USE command
	// to select the database. This is a known bug and might not be necessary in the future,
	// and it is not needed for local databases.
	parsedConnStr := parseConnectionString(spec.ConnectionString)
	connStr := parsedConnStr.path
	if parsedConnStr.label != "" {
		connStr = parsedConnStr.label + ":"
	}
	if parsedConnStr.query != "" {
		connStr += "?" + parsedConnStr.query
	}

	c.connector, err = duckdb.NewConnector(connStr, nil)
	if err != nil {
		return nil, err
	}
	c.db = sql.OpenDB(c.connector)

	if parsedConnStr.label != "" {
		// part of the temporary workaround for external databases mentioned above
		c.logger.Info().Str("label", parsedConnStr.label).Str("database", parsedConnStr.path).Msg("using database for label")
		err = c.exec(ctx, "USE "+parsedConnStr.path+";")
		if err != nil && strings.Contains(err.Error(), "does not exist") {
			// the database doesn't exist yet, so use the default connection string to create it
			// for the first time
			c.connector, err = duckdb.NewConnector(connStr, nil)
			if err != nil {
				return nil, err
			}
			c.db = sql.OpenDB(c.connector)
		} else if err != nil {
			return nil, err
		}
	}

	err = c.exec(ctx, "INSTALL 'json'; LOAD 'json';")
	if err != nil {
		return nil, err
	}
	err = c.exec(ctx, "INSTALL 'parquet'; LOAD 'parquet';")
	if err != nil {
		return nil, err
	}

	return c, nil
}

type parsedConnectionString struct {
	label string // e.g. "md" or "motherduck"
	path  string // e.g. "/path/to/file" or "mydatabase"
	query string // e.g. "token=bla&log_level=debug"
}

func parseConnectionString(conn string) parsedConnectionString {
	reConnString := regexp.MustCompile(`^(?:(?P<label>[a-zA-Z]+):)?(?P<path>[^\s?]*)(?:\?(?P<params>.*))?$`)
	matches := reConnString.FindStringSubmatch(conn)
	if len(matches) == 0 {
		return parsedConnectionString{
			path: conn,
		}
	}
	// get named groups
	groupNames := reConnString.SubexpNames()
	groups := make(map[string]string)
	for i, match := range matches {
		if i == 0 {
			continue
		}
		groups[groupNames[i]] = match
	}
	return parsedConnectionString{
		label: groups["label"],
		path:  groups["path"],
		query: groups["params"],
	}
}

func (c *Client) Close(ctx context.Context) error {
	var err error

	if c.db == nil {
		return fmt.Errorf("client already closed or not initialized")
	}
	//
	//connString := c.duckSpec.ConnectionString
	//var label string
	//switch {
	//case strings.HasPrefix(connString, "md:"):
	//	label = strings.TrimPrefix(connString, "md:")
	//case strings.HasPrefix(connString, "motherduck:"):
	//	label = strings.TrimPrefix(connString, "motherduck:")
	//}
	//if label != "" {
	//	err = c.exec(ctx, fmt.Sprintf("DETACH DATABASE IF EXISTS '%s';", label))
	//	if err != nil {
	//		return err
	//	}
	//}

	err = c.db.Close()
	c.db = nil
	return err
}

func (c *Client) Metrics() destination.Metrics {
	return c.metrics
}

func (c *Client) exec(ctx context.Context, query string, args ...any) error {
	_, err := c.db.ExecContext(ctx, query, args...)
	return err
}
