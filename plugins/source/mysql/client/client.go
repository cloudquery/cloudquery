package client

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

type Client struct {
	logger      zerolog.Logger
	metrics     *source.Metrics
	Tables      schema.Tables
	db          *sql.DB
	tableSchema string
}

var _ schema.ClientMeta = (*Client)(nil)

func (*Client) ID() string {
	return "source-mysql"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
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

	c := &Client{logger: logger.With().Str("module", "mysql-source").Logger(), db: db, tableSchema: dsn.DBName}
	c.Tables, err = c.listTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	c.Tables, err = c.Tables.FilterDfs(spec.Tables, spec.SkipTables, spec.SkipDependentTables)
	if err != nil {
		return nil, fmt.Errorf("failed to apply config to tables: %w", err)
	}

	return c, nil
}
