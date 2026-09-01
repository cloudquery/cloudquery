package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// COPY prepares a statement to learn the column types, and upserts additionally
// create and drop a staging table; below this both cost more than they save.
// A var so tests can reach the COPY path with the small batches they build.
var minCopyRows int64 = 100

// CrateDB has no COPY over the wire protocol and CockroachDB gates temporary
// tables behind an experimental setting. pgvector inserts embeddings after each
// flush, which COPY has no equivalent of.
func (c *Client) useCopyFrom() bool {
	return c.spec.UseCopyFrom && c.pgType == pgTypePostgreSQL && !c.hasPgVectorConfig()
}

type copyGroup struct {
	table *schema.Table
	msgs  message.WriteInserts
	rows  int64
}

func (g *copyGroup) eligible() bool {
	if g.rows < minCopyRows {
		return false
	}
	// PkConstraintName is only populated once migrations have read it back from
	// the database; without it there is nothing to conflict on.
	return len(g.table.PrimaryKeysIndexes()) == 0 || g.table.PkConstraintName != ""
}

func (c *Client) insertBatchCopy(ctx context.Context, messages message.WriteInserts, pgTables map[string]struct{}) error {
	// Grouping preserves arrival order both across and within tables, which is
	// what makes the last row win when a primary key repeats.
	var groups []*copyGroup
	byTable := make(map[string]*copyGroup)

	for _, msg := range messages {
		tableName, ok := msg.Record.Schema().Metadata().GetValue(schema.MetadataTableName)
		if !ok {
			return errors.New("table name not found in metadata")
		}
		if _, ok := pgTables[tableName]; !ok {
			return fmt.Errorf("table %s not found", tableName)
		}
		g := byTable[tableName]
		if g == nil {
			g = &copyGroup{table: c.normalizeTable(msg.GetTable())}
			byTable[tableName] = g
			groups = append(groups, g)
		}
		g.msgs = append(g.msgs, msg)
		g.rows += msg.Record.NumRows()
	}

	// Ineligible groups go into a single pipelined batch: one SendBatch for all
	// of them together, rather than one per table. A table is either wholly
	// eligible or wholly ineligible, so no table's rows are split across the two
	// paths and per-table arrival order still holds.
	var execMsgs message.WriteInserts
	for _, g := range groups {
		if g.eligible() {
			if err := c.copyTable(ctx, g.table, g.msgs); err != nil {
				return err
			}
			continue
		}
		execMsgs = append(execMsgs, g.msgs...)
	}
	if len(execMsgs) == 0 {
		return nil
	}
	return c.insertBatchExec(ctx, execMsgs, pgTables)
}

func (c *Client) copyTable(ctx context.Context, table *schema.Table, msgs message.WriteInserts) error {
	columns := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		columns[i] = col.Name
	}
	src := &copyFromRecords{client: c, msgs: msgs}

	if len(table.PrimaryKeysIndexes()) == 0 {
		return c.retryOnDeadlock(func() error {
			src.reset()
			_, err := c.conn.CopyFrom(ctx, pgx.Identifier{table.Name}, columns, src)
			return err
		}, "failed to copy into "+table.Name)
	}

	return c.retryOnDeadlock(func() error {
		src.reset()
		return c.upsertViaStagingTable(ctx, table, columns, src)
	}, "failed to upsert into "+table.Name)
}

// COPY has no ON CONFLICT, so an upsert stages the rows and merges them. The
// transaction both scopes the temporary table to this connection and drops it on
// either outcome.
func (c *Client) upsertViaStagingTable(ctx context.Context, table *schema.Table, columns []string, src pgx.CopyFromSource) error {
	conn, err := c.conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	staging := stagingTableName()
	// LIKE without INCLUDING carries no constraints, so repeated primary keys
	// land here for the merge to resolve rather than being rejected on the way in.
	createSQL := fmt.Sprintf("create temp table %s (like %s) on commit drop",
		pgx.Identifier{staging}.Sanitize(), pgx.Identifier{table.Name}.Sanitize())
	if _, err := tx.Exec(ctx, createSQL); err != nil {
		return fmt.Errorf("failed to create staging table: %w", err)
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{staging}, columns, src); err != nil {
		return fmt.Errorf("failed to copy into staging table: %w", err)
	}

	if _, err := tx.Exec(ctx, mergeFromStaging(table, staging)); err != nil {
		return fmt.Errorf("failed to merge staging table: %w", err)
	}

	return tx.Commit(ctx)
}

// distinct on is load-bearing: on conflict do update rejects a statement whose
// source repeats a conflict key ("cannot affect row a second time"). ctid desc
// keeps the last row copied, matching insertBatchExec.
func mergeFromStaging(table *schema.Table, staging string) string {
	quotedColumns := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		quotedColumns[i] = pgx.Identifier{col.Name}.Sanitize()
	}
	primaryKeys := table.PrimaryKeys()
	quotedPrimaryKeys := make([]string, len(primaryKeys))
	for i, pk := range primaryKeys {
		quotedPrimaryKeys[i] = pgx.Identifier{pk}.Sanitize()
	}
	columnList := strings.Join(quotedColumns, ",")
	primaryKeyList := strings.Join(quotedPrimaryKeys, ",")

	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
	sb.WriteString(columnList)
	sb.WriteString(") select distinct on (")
	sb.WriteString(primaryKeyList)
	sb.WriteString(") ")
	sb.WriteString(columnList)
	sb.WriteString(" from ")
	sb.WriteString(pgx.Identifier{staging}.Sanitize())
	sb.WriteString(" order by ")
	sb.WriteString(primaryKeyList)
	sb.WriteString(", ctid desc on conflict on constraint ")
	sb.WriteString(pgx.Identifier{table.PkConstraintName}.Sanitize())
	sb.WriteString(" do update set ")
	for i, quoted := range quotedColumns {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(quoted)
		sb.WriteString("=excluded.")
		sb.WriteString(quoted)
	}
	return sb.String()
}

func stagingTableName() string {
	return "cq_staging_" + strings.ReplaceAll(uuid.NewString(), "-", "")
}

// copyFromRecords transforms one record at a time so a batch is never fully
// materialized as Go values.
type copyFromRecords struct {
	client *Client
	msgs   message.WriteInserts

	msgIdx  int
	rows    [][]any
	rowIdx  int
	current []any
}

func (s *copyFromRecords) reset() {
	s.msgIdx, s.rows, s.rowIdx, s.current = 0, nil, 0, nil
}

func (s *copyFromRecords) Next() bool {
	for s.rowIdx >= len(s.rows) { // loops to skip records with no rows
		if s.msgIdx >= len(s.msgs) {
			return false
		}
		s.rows = s.client.transformValues(s.msgs[s.msgIdx].Record)
		s.rowIdx = 0
		s.msgIdx++
	}
	s.current = s.rows[s.rowIdx]
	s.rowIdx++
	return true
}

func (s *copyFromRecords) Values() ([]any, error) { return s.current, nil }

func (*copyFromRecords) Err() error { return nil }
