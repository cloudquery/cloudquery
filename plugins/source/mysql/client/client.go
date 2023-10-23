package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedDestination
	logger      zerolog.Logger
	tables      schema.Tables
	options     plugin.NewClientOptions
	db          *sql.DB
	tableSchema string
}

var _ schema.ClientMeta = (*Client)(nil)

func (*Client) ID() string {
	return "source-mysql"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger:  logger,
			options: opts,
			tables:  schema.Tables{},
		}, nil
	}
	var mySQLSpec Spec
	err := json.Unmarshal(spec, &mySQLSpec)
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

	c := Client{logger: logger.With().Str("module", "mysql-source").Logger(), db: db, tableSchema: dsn.DBName, options: opts}
	c.tables, err = c.listTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to apply config to tables: %w", err)
	}

	return c, nil
}

func (c Client) Tables(ctx context.Context, opts plugin.TableOptions) (schema.Tables, error) {
	if c.options.NoConnection {
		return schema.Tables{}, nil
	}
	return c.tables.FilterDfs(opts.Tables, opts.SkipTables, opts.SkipDependentTables)
}

func (c Client) Close(_ context.Context) error {
	return c.db.Close()
}
