package client

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"

	mysql "github.com/go-sql-driver/mysql"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	destination.DefaultReverseTransformer
	logger zerolog.Logger

	spec      specs.Destination
	mySQLSpec Spec

	db *sql.DB
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var mySQLSpec Spec
	err := spec.UnmarshalSpec(&mySQLSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	mySQLSpec.SetDefaults()
	if err := mySQLSpec.Validate(); err != nil {
		return nil, err
	}

	dsn, err := mysql.ParseDSN(mySQLSpec.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("invalid MySQL connection string: %w", err)
	}
	if dsn.Params == nil {
		dsn.Params = map[string]string{}
	}
	dsn.Params["parseTime"] = "true"
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Client{logger: logger.With().Str("module", "mysql").Logger(), db: db, spec: spec, mySQLSpec: mySQLSpec}, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.db.Close()
}

func (c *Client) pkEnabled() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
