package history

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/helpers"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	listTables        = `SELECT table_name FROM information_schema.tables WHERE table_schema=$1 AND table_type='BASE TABLE' ORDER BY 1`
	getColumnComments = `WITH x AS (SELECT column_name, pg_catalog.col_description(format('%s.%s',table_schema,table_name)::regclass::oid,ordinal_position) AS comment FROM information_schema.columns w
here table_schema=$1 AND table_name=$2) SELECT * FROM x WHERE comment IS NOT NULL;`

	createHyperTable    = `SELECT * FROM create_hypertable($1, 'cq_fetch_date', chunk_time_interval => INTERVAL '%d day', if_not_exists => true);`
	dataRetentionPolicy = `SELECT add_retention_policy($1, INTERVAL '%d day', if_not_exists => true);`

	dropTableView   = `DROP VIEW IF EXISTS "%[1]s"`
	createTableView = `CREATE VIEW "%[1]s" AS SELECT * FROM history."%[1]s" WHERE cq_fetch_date = find_latest('history', '%[1]s')`

	SchemaName = "history"
)

type DDLManager struct {
	log     hclog.Logger
	conn    *pgxpool.Conn
	cfg     *Config
	dialect schema.Dialect
}

func NewDDLManager(l hclog.Logger, conn *pgxpool.Conn, cfg *Config, dt schema.DialectType) (*DDLManager, error) {
	if dt != schema.TSDB {
		return nil, fmt.Errorf("history is only supported on timescaledb")
	}

	return &DDLManager{
		log:     l,
		conn:    conn,
		cfg:     cfg,
		dialect: schema.GetDialect(dt),
	}, nil
}

func (h DDLManager) SetupHistory(ctx context.Context, conn *pgxpool.Conn) error {
	var tables []string
	if err := pgxscan.Get(ctx, conn, &tables, listTables, SchemaName); err != nil {
		return fmt.Errorf("failed to list tables: %w", err)
	}

	for _, table := range tables {
		parentIdCol, parentTable, err := h.getTableParent(ctx, conn, table)
		if err != nil {
			return fmt.Errorf("getTableParent failed for %s: %w", table, err)
		}

		if err := h.createHyperTable(ctx, conn, table, parentTable != ""); err != nil {
			return fmt.Errorf("failed to create hypertable for table: %s: %w", table, err)
		}
		if err := h.recreateView(ctx, conn, table); err != nil {
			return fmt.Errorf("recreateView: %w", err)
		}

		if parentTable != "" {
			if err := h.buildCascadeTrigger(ctx, conn, table, parentIdCol, parentTable); err != nil {
				return fmt.Errorf("table build %s failed: %w", table, err)
			}
		}
	}

	return nil
}

func (h DDLManager) createHyperTable(ctx context.Context, conn *pgxpool.Conn, tableName string, hasParent bool) error {
	var hyperTable struct {
		HypertableId int    `db:"hypertable_id"`
		SchemaName   string `db:"schema_name"`
		TableName    string `db:"table_name"`
		Created      bool   `db:"created"`
	}

	tName := fmt.Sprintf(`"%s"."%s"`, SchemaName, tableName)
	if err := pgxscan.Get(ctx, conn, &hyperTable, fmt.Sprintf(createHyperTable, h.cfg.TimeInterval), tName); err != nil {
		return fmt.Errorf("failed to create hypertable: %w", err)
	}
	h.log.Debug("created hyper table for table", "table", hyperTable.TableName, "id", hyperTable.HypertableId, "created", hyperTable.Created)
	if hasParent { // TODO
		return nil
	}
	if _, err := conn.Exec(ctx, fmt.Sprintf(dataRetentionPolicy, h.cfg.Retention), tName); err != nil {
		return err
	}
	h.log.Debug("created data retention policy", "table", hyperTable.TableName, "days", h.cfg.Retention)
	return nil
}

func (h DDLManager) recreateView(ctx context.Context, conn *pgxpool.Conn, table string) error {
	if err := conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		// Must drop the view first -- CREATE OR REPLACE view won't cut it if columns are changed. PostgreSQL doc states:
		// > The new query must generate the same columns that were generated by the existing view query (that is, the same column names in the same order and with
		// > the same data types), but it may add additional columns to the end of the list.
		// ref: https://www.postgresql.org/docs/14/sql-createview.html

		if _, err := tx.Exec(ctx, fmt.Sprintf(dropTableView, table)); err != nil {
			return fmt.Errorf("failed to drop view for table: %w", err)
		}

		if _, err := tx.Exec(ctx, fmt.Sprintf(createTableView, table)); err != nil {
			return fmt.Errorf("failed to create view for table: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("tx failed for %s: %w", table, err)
	}
	return nil
}

func (h DDLManager) buildCascadeTrigger(ctx context.Context, conn *pgxpool.Conn, table, parentIdColumn, parentTable string) error {
	if _, err := conn.Exec(ctx, "SELECT history.build_trigger($1, $2, $3);", parentTable, table, parentIdColumn); err != nil {
		return fmt.Errorf("failed to create trigger: %w", err)
	}
	return nil
}

func (h DDLManager) getTableParent(ctx context.Context, conn *pgxpool.Conn, tableName string) (parentIdColumn, parentTable string, err error) {
	var comments []struct {
		Col     string `db:"column_name"`
		Comment string `db:"comment"`
	}
	if err := pgxscan.Select(ctx, conn, &comments, getColumnComments, SchemaName, tableName); err != nil {
		return "", "", fmt.Errorf("failed to get column comments: %w", err)
	}

	found := 0
	for _, c := range comments {
		parentTable, _ = schema.GetFKFromComment(c.Comment)
		if parentTable != "" {
			found++
			parentIdColumn = c.Col
		}
	}

	if found > 1 {
		return "", "", fmt.Errorf("multiple FK comments found in table")
	}

	return
}

func AddHistoryFunctions(ctx context.Context, conn *pgxpool.Conn) error {
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

func TransformDSN(dsn string) (string, error) {
	return setDsnElement(dsn, map[string]string{"search_path": SchemaName})
}

func setDsnElement(dsn string, elems map[string]string) (string, error) {
	u, err := helpers.ParseConnectionString(dsn)
	if err != nil {
		return "", err
	}

	vals := u.Query()
	for k, v := range elems {
		vals.Set(k, v)
	}
	u.RawQuery = vals.Encode()
	return u.String(), nil
}
