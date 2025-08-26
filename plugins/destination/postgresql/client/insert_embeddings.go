package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const CQIDColumn = "_cq_id"

// embeddingBatch accumulates rows for a single base table to insert into <table>_embeddings.
type embeddingBatch struct {
	tableSchema *embeddingTable
	rows        [][]any
	cqIDs       []uuid.UUID
}

// embeddingTable mirrors the parts we need from schema.Table.
type embeddingTable struct {
	Name    string
	Columns []embeddingTableColumn
	PkName  string
}

type embeddingTableColumn struct {
	Name       string
	PrimaryKey bool
}

func toEmbeddingTableColumns(cols []schema.Column) []embeddingTableColumn {
	out := make([]embeddingTableColumn, len(cols))
	for i, c := range cols {
		out[i] = embeddingTableColumn{Name: c.Name, PrimaryKey: c.PrimaryKey}
	}
	return out
}

// addEmbeddingRows records a set of transformed rows for a given table into the provided map.
func (c *Client) addEmbeddingRows(ctx context.Context, tableToBatch map[string]*embeddingBatch, tableName string, tbl *schema.Table, r arrow.Record, rows [][]any) error {
	if err := c.ensureConfigColumnsMatchBatchSchema(tbl); err != nil {
		return err
	}

	if err := c.createPgVectorTableIfNotExists(ctx, tbl); err != nil {
		return err
	}

	b := tableToBatch[tableName]
	if b == nil {
		b = &embeddingBatch{tableSchema: &embeddingTable{
			Name:    tbl.Name,
			Columns: toEmbeddingTableColumns(tbl.Columns),
			PkName:  tbl.PkConstraintName,
		}}
	}
	cqIDs, err := c.getCQIDs(r)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to get CQIDs")
		return nil
	}
	b.rows = append(b.rows, rows...)
	b.cqIDs = append(b.cqIDs, cqIDs...)
	tableToBatch[tableName] = b
	return nil
}

// insertEmbeddingsBatch requests embeddings for accumulated rows and inserts/upserts them into <table>_embeddings.
func (c *Client) insertEmbeddingsBatch(ctx context.Context, tableToBatch map[string]*embeddingBatch) error {
	if len(tableToBatch) == 0 {
		return nil
	}
	if !c.hasPgVectorConfig() || c.embeddingsRequester == nil {
		return nil
	}
	if err := c.ensurePgVectorExtensionInstalled(ctx); err != nil {
		return err
	}

	for tableName, rowBatch := range tableToBatch {
		// Ensure table has a pgvector config
		tblCfg := c.spec.GetPgVectorTableConfig(tableName)
		if tblCfg == nil {
			continue
		}

		// Build inputs from embed columns per row (concatenate with a separator)
		columnIdxs, err := columnIndexes(rowBatch.tableSchema.Columns, tblCfg.EmbedColumns)
		if err != nil {
			return err
		}
		inputs := buildInputs(rowBatch.rows, columnIdxs)

		// Request embeddings for all inputs at once, preserving order
		inputToEmbeddingResponses, err := c.embeddingsRequester.CreateEmbeddings(ctx, inputs)
		if err != nil {
			return err
		}
		if len(inputToEmbeddingResponses) != len(rowBatch.rows) {
			return fmt.Errorf("embeddings count mismatch: got %d want %d", len(inputToEmbeddingResponses), len(rowBatch.rows))
		}

		// Prepare insert SQL for <table>_embeddings
		embTableName := tableName + "_embeddings"
		sql := c.buildEmbeddingsInsertSQL(embTableName, tblCfg)

		// We cannot upsert in this case, because the number of rows per `_cq_id` might have changed
		// so we need to delete the rows with the cqIDs before inserting
		if deleteSQL := c.buildDeleteCQIDsSQL(embTableName, rowBatch.cqIDs); deleteSQL != "" {
			if _, err := c.conn.Exec(ctx, deleteSQL); err != nil {
				return err
			}
		}

		batch := new(pgx.Batch)
		for i, row := range rowBatch.rows {
			for _, vals := range c.buildEmbeddingsRowValues(row, rowBatch.tableSchema.Columns, tblCfg, inputToEmbeddingResponses[i]) {
				batch.Queue(sql, vals...)
			}
			if int64(batch.Len()) >= c.batchSize {
				if err := c.flushBatch(ctx, batch); err != nil {
					return err
				}
				batch = new(pgx.Batch)
			}
		}
		if err := c.flushBatch(ctx, batch); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ensureConfigColumnsMatchBatchSchema(tbl *schema.Table) error {
	tblCfg := c.spec.GetPgVectorTableConfig(tbl.Name)
	if tblCfg == nil {
		return nil
	}

	// Build a set of base table column names for quick membership checks
	colSet := make(map[string]struct{}, len(tbl.Columns))
	for _, c := range tbl.Columns {
		colSet[c.Name] = struct{}{}
	}

	// Check both metadata and embed columns exist on the base table
	missing := make([]string, 0)
	for _, name := range tblCfg.MetadataColumns {
		if _, ok := colSet[name]; !ok {
			missing = append(missing, name)
		}
	}
	for _, name := range tblCfg.EmbedColumns {
		if _, ok := colSet[name]; !ok {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("pgvector config for %s references missing columns: %s", tbl.Name, strings.Join(missing, ", "))
	}

	return nil
}

func columnIndexes(cols []embeddingTableColumn, names []string) ([]int, error) {
	idxByName := make(map[string]int, len(cols))
	for i, c := range cols {
		idxByName[c.Name] = i
	}
	out := make([]int, len(names))
	for i, n := range names {
		pos, ok := idxByName[n]
		if !ok {
			return nil, fmt.Errorf("column %s not found in table", n)
		}
		out[i] = pos
	}
	return out, nil
}

func (*Client) buildDeleteCQIDsSQL(embTable string, cqIDs []uuid.UUID) string {
	ids := make([]string, len(cqIDs))
	for i, id := range cqIDs {
		ids[i] = fmt.Sprintf("'%s'", id.String())
	}
	return fmt.Sprintf("delete from %s where %s in (%s)", embTable, pgx.Identifier{CQIDColumn}.Sanitize(), strings.Join(ids, ","))
}

func (*Client) buildEmbeddingsInsertSQL(embTable string, cfg *spec.PgVectorTableConfig) string {
	return fmt.Sprintf(
		"insert into %s (%s) values (%s)",
		pgx.Identifier{embTable}.Sanitize(),
		strings.Join(sanitizeIdentifiers(makeSlice(cfg.MetadataColumns, []string{"chunk", "embedding"})), ","),
		strings.Join(makePlaceholders(len(cfg.MetadataColumns)+2), ","),
	)
}

func (*Client) buildEmbeddingsRowValues(row []any, cols []embeddingTableColumn, cfg *spec.PgVectorTableConfig, embeddingResponses []EmbeddingResponse) [][]any {
	// Precompute values for metadata columns in the same order as metadataColumns
	metaVals := make([]any, 0, len(cfg.MetadataColumns))
	idxByColName := make(map[string]int, len(cols))
	colNames := make([]string, len(cols))
	for i, ccol := range cols {
		idxByColName[ccol.Name] = i
		colNames[i] = ccol.Name
	}
	for _, name := range cfg.MetadataColumns {
		idx := idxByColName[name]
		var val any
		if idx >= 0 && idx < len(row) {
			val = row[idx]
		}
		metaVals = append(metaVals, val)
	}

	// For each embedding response, clone metadata values and append chunk and vector
	out := make([][]any, 0, len(embeddingResponses))
	for _, er := range embeddingResponses {
		vals := make([]any, 0, len(metaVals)+2)
		vals = append(vals, metaVals...)
		// chunk text, then embedding vector as string literal (cast to vector in SQL)
		vals = append(vals, er.Chunk)
		vals = append(vals, embeddingLiteral(er.Vector))
		out = append(out, vals)
	}
	return out
}

func buildInputs(rows [][]any, embedColIdx []int) []string {
	inputs := make([]string, 0, len(rows))
	for _, row := range rows {
		parts := make([]string, 0, len(embedColIdx))
		for _, idx := range embedColIdx {
			if idx >= 0 && idx < len(row) {
				if v := row[idx]; v != nil {
					parts = append(parts, scalarToString(v))
				}
			}
		}
		inputs = append(inputs, strings.Join(parts, "\n"))
	}
	return inputs
}

func scalarToString(v any) string {
	switch t := v.(type) {
	case pgtype.Text:
		if t.Valid {
			return t.String
		}
		return ""
	case *pgtype.Text:
		if t != nil && t.Valid {
			return t.String
		}
		return ""
	case string:
		return t
	default:
		return fmt.Sprint(v)
	}
}

func embeddingLiteral(vec []float32) string {
	if len(vec) == 0 {
		return "[]"
	}
	parts := make([]string, len(vec))
	for i, f := range vec {
		parts[i] = strconv.FormatFloat(float64(f), 'f', -1, 32)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func sanitizeIdentifiers(idents []string) []string {
	out := make([]string, len(idents))
	for i, ident := range idents {
		out[i] = pgx.Identifier{ident}.Sanitize()
	}
	return out
}

func makeSlice[T any](slices ...[]T) []T {
	out := make([]T, 0)
	for _, slice := range slices {
		out = append(out, slice...)
	}
	return out
}

func makePlaceholders(n int) []string {
	out := make([]string, 0, n)
	for i := range n {
		out = append(out, fmt.Sprintf("$%d", i+1))
	}
	return out
}
