package history

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	createHistorySchema        = `CREATE SCHEMA IF NOT EXISTS history;`
	validateTimescaleInstalled = `SELECT EXISTS(SELECT 1 FROM pg_extension where extname = 'timescaledb')`
	cascadeDeleteFunction      = `
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

	createHyperTable    = `SELECT * FROM create_hypertable($1, 'cq_fetch_date', chunk_time_interval => INTERVAL '%d day', if_not_exists => true);`
	dataRetentionPolicy = `SELECT add_retention_policy($1, INTERVAL '%d day', if_not_exists => true);`
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

	dropTableView   = `DROP VIEW IF EXISTS "%[1]s"`
	createTableView = `CREATE VIEW "%[1]s" AS SELECT * FROM history."%[1]s" WHERE cq_fetch_date = find_latest('history', '%[1]s')`
)

type Config struct {
	// Retention of data in days, defaults to 7
	Retention int `default:"7" hcl:"retention,optional"`
	// TimeInterval defines how chunks are split by time defaults to one chunk per 24 hours.
	TimeInterval int `default:"24" hcl:"interval,optional"`
	// TimeTruncation truncates fetch time by hour, for example if we fetch with TimeTruncation = 1 at 11:25 the fetch date will truncate to 11:00
	// defaults to 24 hours, which means one set of fetch data per day.
	TimeTruncation int `default:"24" hcl:"truncation,optional"`
}

func (c Config) FetchDate() time.Time {
	return time.Now().UTC().Truncate(time.Duration(c.TimeTruncation) * time.Hour)
}

// SetupHistory sets all required history functions and validation checks that it can run cleanly.
func SetupHistory(ctx context.Context, conn *pgxpool.Conn) error {
	installed, err := validateInstalled(ctx, conn)
	if err != nil {
		return fmt.Errorf("failed to validate timescale installed: %w", err)
	}
	if !installed {
		return errors.New("timescaledb extension not installed, `CREATE EXTENSION IF NOT EXISTS timescaledb;`")
	}
	if err := addHistoryFunctions(ctx, conn); err != nil {
		return fmt.Errorf("failed to create history functions: %w", err)
	}
	return nil
}

func addHistoryFunctions(ctx context.Context, conn *pgxpool.Conn) error {
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

func validateInstalled(ctx context.Context, conn *pgxpool.Conn) (bool, error) {
	var installed bool
	if err := pgxscan.Get(ctx, conn, &installed, validateTimescaleInstalled); err != nil {
		return false, err
	}
	return installed, nil
}
