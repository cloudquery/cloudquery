package timescale

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/client/database/postgres"
	pgsdk "github.com/cloudquery/cq-provider-sdk/database/postgres"
	"github.com/cloudquery/cq-provider-sdk/helpers"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	createHyperTable    = `SELECT * FROM create_hypertable($1, 'cq_fetch_date', chunk_time_interval => INTERVAL '%d day', if_not_exists => true);`
	dataRetentionPolicy = `SELECT add_retention_policy($1, INTERVAL '%d day', if_not_exists => true);`

	dropTableView   = `DROP VIEW IF EXISTS "%[1]s"`
	createTableView = `CREATE VIEW "%[1]s" AS SELECT * FROM history."%[1]s" WHERE cq_fetch_date = find_latest('history', '%[1]s')`

	historySchema = "history"
	viewSchema    = "public"
)

type Executor struct {
	logger hclog.Logger
	dsn    string
}

func New(logger hclog.Logger, dsn string) Executor {
	return Executor{
		logger: logger,
		dsn:    dsn,
	}
}

// Setup sets all required history functions and validation checks that it can run cleanly.
func (e Executor) Setup(ctx context.Context) (string, error) {
	pool, err := pgsdk.Connect(ctx, e.dsn)
	if err != nil {
		return e.dsn, err
	}
	defer pool.Close()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return e.dsn, err
	}
	defer conn.Release()

	if err := addHistoryFunctions(ctx, conn); err != nil {
		return e.dsn, fmt.Errorf("failed to create history functions: %w", err)
	}

	return setDsnElement(e.dsn, map[string]string{"search_path": historySchema}), nil
}

func (e Executor) Validate(ctx context.Context) (bool, error) {
	const (
		validateTimescaleInstalled = `SELECT EXISTS(SELECT 1 FROM pg_extension where extname = 'timescaledb')`
	)

	pool, err := pgsdk.Connect(ctx, e.dsn)
	if err != nil {
		return false, err
	}
	defer pool.Close()

	if err := postgres.ValidatePostgresVersion(ctx, pool, postgres.MinPostgresVersion); err != nil {
		return false, err
	}

	var installed bool
	if err := pgxscan.Get(ctx, pool, &installed, validateTimescaleInstalled); err != nil {
		return false, err
	}
	if !installed {
		return false, fmt.Errorf("timescaledb extension not installed, `CREATE EXTENSION IF NOT EXISTS timescaledb;`")
	}

	return true, nil
}

func (e Executor) Finalize(ctx context.Context) error {
	// TODO do view stuff
	return nil
}

func addHistoryFunctions(ctx context.Context, conn *pgxpool.Conn) error {
	const (
		createHistorySchema   = `CREATE SCHEMA IF NOT EXISTS history;`
		cascadeDeleteFunction = `
				CREATE OR REPLACE FUNCTION history.cascade_delete()
					RETURNS trigger
					LANGUAGE 'plpgsql'
					COST 100
					VOLATILE NOT LEAKPROOF
				AS $BODY$
				BEGIN
					BEGIN
						IF (TG_OP = 'DELETE') THEN
							EXECUTE format('DELETE FROM history.%I where %I = %L AND cq_fetch_date = %L', TG_ARGV[0], TG_ARGV[1], OLD.cq_id, OLD.cq_fetch_date);
							RETURN OLD;
						END IF;
						RETURN NULL; -- result is ignored since this is an AFTER trigger
					END;
				END;
				$BODY$;`
		buildTriggerFunction = `
				CREATE OR REPLACE FUNCTION history.build_trigger(_table_name text, _child_table_name text, _parent_id text)
					RETURNS integer
					LANGUAGE 'plpgsql'
					COST 100
					VOLATILE PARALLEL UNSAFE
				AS $BODY$
				BEGIN
					IF NOT EXISTS ( SELECT 1 FROM pg_trigger WHERE tgname = _child_table_name )  then
					EXECUTE format(
						'CREATE TRIGGER %I BEFORE DELETE ON history.%I FOR EACH ROW EXECUTE PROCEDURE history.cascade_delete(%s, %s)'::text,
						_child_table_name, _table_name, _child_table_name, _parent_id);
					return 0;
					ELSE
						return 1;
					END IF;
				END;
				$BODY$;`

		findLatestFetchDate = `
			CREATE OR REPLACE FUNCTION find_latest(schema TEXT, _table_name TEXT) 
			RETURNS timestamp without time zone AS $body$
			DECLARE
			 fetchDate timestamp without time zone;
			BEGIN
				EXECUTE format('SELECT cq_fetch_date FROM %I.%I order by cq_fetch_date desc limit 1', schema, _table_name) into fetchDate;
				return fetchDate;
			END;
			$body$  LANGUAGE plpgsql IMMUTABLE`
	)

	return conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, createHistorySchema); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, buildTriggerFunction); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, cascadeDeleteFunction); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, findLatestFetchDate); err != nil {
			return err
		}
		return nil
	})
}

type createHyperTableResult struct {
	HypertableId int    `db:"hypertable_id"`
	SchemaName   string `db:"schema_name"`
	TableName    string `db:"table_name"`
	Created      bool   `db:"created"`
}

func setPath(ctx context.Context, conn *pgx.Conn, schemaName string) error {
	_, err := conn.Exec(ctx, "SET search_path TO $1", schemaName)
	return err
}

func setDsnElement(dsn string, elems map[string]string) string {
	u, err := helpers.ParseConnectionString(dsn)
	if err != nil {
		panic(err.Error())
	}

	vals := u.Query()
	for k, v := range elems {
		vals.Set(k, v)
	}
	u.RawQuery = vals.Encode()
	return u.String()
}
