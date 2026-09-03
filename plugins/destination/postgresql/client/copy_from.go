package client

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// Below this, preparing the COPY statement and staging an upsert cost more than
// they save. A var so tests can reach the COPY path with small batches.
var minCopyRows int64 = 100

// CrateDB has no COPY over the wire protocol and CockroachDB gates temporary
// tables behind an experimental setting. pgvector inserts embeddings after each
// flush, which COPY has no equivalent of.
func (c *Client) useCopyFrom() bool {
	return c.spec.UseCopyFrom && c.pgType == pgTypePostgreSQL && !c.hasPgVectorConfig()
}

type copyGroup struct {
	table     *schema.Table
	pkColumns []string // the database's, which are not always the source schema's
	msgs      message.WriteInserts
	rows      int64
}

func (g *copyGroup) eligible() bool {
	if g.rows < minCopyRows {
		return false
	}
	if len(g.table.PrimaryKeysIndexes()) == 0 {
		return true
	}
	return g.table.PkConstraintName != "" && len(g.pkColumns) > 0
}

// The database's primary key columns, which the merge dedupes on: deduping on the
// source schema's instead lets two rows share a constraint key once they drift.
func (c *Client) pkConstraintColumns(tableName string) []string {
	c.pgTablesToPKConstraintsMu.RLock()
	defer c.pgTablesToPKConstraintsMu.RUnlock()
	if entry := c.pgTablesToPKConstraints[tableName]; entry != nil {
		return slices.Clone(entry.columns)
	}
	return nil
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
			g = &copyGroup{
				table:     c.normalizeTable(msg.GetTable()),
				pkColumns: c.pkConstraintColumns(tableName),
			}
			byTable[tableName] = g
			groups = append(groups, g)
		}
		g.msgs = append(g.msgs, msg)
		g.rows += msg.Record.NumRows()
	}

	var execMsgs message.WriteInserts
	for _, g := range groups {
		if g.eligible() {
			if err := c.copyTable(ctx, g); err != nil {
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

func (c *Client) copyTable(ctx context.Context, g *copyGroup) error {
	table := g.table
	columns := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		columns[i] = col.Name
	}
	src := &copyFromRecords{client: c, msgs: g.msgs}

	// Postgres rejects COPY FROM outright on a table with row-level security
	// enabled, so this branch fails where insertBatchExec would have worked.
	if len(table.PrimaryKeysIndexes()) == 0 {
		return c.retryOnDeadlock(func() error {
			src.reset()
			_, err := c.conn.CopyFrom(ctx, pgx.Identifier{table.Name}, columns, src)
			return err
		}, "failed to copy into "+table.Name)
	}

	return c.retryOnDeadlock(func() error {
		src.reset()
		return c.upsertViaStagingTable(ctx, table, columns, g.pkColumns, src)
	}, "failed to upsert into "+table.Name)
}

func (c *Client) upsertViaStagingTable(ctx context.Context, table *schema.Table, columns, pkColumns []string, src pgx.CopyFromSource) error {
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
	// The staging table has no policies of its own; the merge enforces the target's.
	createSQL := fmt.Sprintf("create temp table %s (like %s) on commit drop",
		pgx.Identifier{staging}.Sanitize(), pgx.Identifier{table.Name}.Sanitize())
	if _, err := tx.Exec(ctx, createSQL); err != nil {
		return fmt.Errorf("failed to create staging table: %w", err)
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{staging}, columns, src); err != nil {
		return fmt.Errorf("failed to copy into staging table: %w", err)
	}

	if _, err := tx.Exec(ctx, mergeFromStaging(table, staging, pkColumns)); err != nil {
		return fmt.Errorf("failed to merge staging table: %w", err)
	}

	return tx.Commit(ctx)
}

// distinct on is load-bearing: on conflict do update rejects a statement whose
// source repeats a conflict key ("cannot affect row a second time"). ctid desc
// keeps the last row copied, matching insertBatchExec.
func mergeFromStaging(table *schema.Table, staging string, pkColumns []string) string {
	quotedColumns := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		quotedColumns[i] = pgx.Identifier{col.Name}.Sanitize()
	}
	quotedPKColumns := make([]string, len(pkColumns))
	for i, pk := range pkColumns {
		quotedPKColumns[i] = pgx.Identifier{pk}.Sanitize()
	}
	columnList := strings.Join(quotedColumns, ",")
	pkColumnList := strings.Join(quotedPKColumns, ",")

	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
	sb.WriteString(columnList)
	sb.WriteString(") select distinct on (")
	sb.WriteString(pkColumnList)
	sb.WriteString(") ")
	sb.WriteString(columnList)
	sb.WriteString(" from ")
	sb.WriteString(pgx.Identifier{staging}.Sanitize())
	sb.WriteString(" order by ")
	sb.WriteString(pkColumnList)
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
	for s.rowIdx >= len(s.rows) {
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
