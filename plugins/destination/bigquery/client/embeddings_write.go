package client

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

const (
	CQIDColumn = "_cq_id"
)

func (c *ConcreteEmbeddingsClient) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	if !c.hasMigrated {
		if err := c.MigrateTables(ctx); err != nil {
			return err
		}
		c.hasMigrated = true
	}

	tableConfig, err := c.findTableConfig(name)
	if err != nil {
		return err
	}

	cqIDs, err := extractCqIDs(msgs)
	if err != nil {
		return fmt.Errorf("failed to extract CQ IDs: %w", err)
	}

	query, err := c.buildEmbeddingQuery(tableConfig, cqIDs)
	if err != nil {
		return fmt.Errorf("failed to build embedding query: %w", err)
	}

	job, err := c.client.client.Query(query).Run(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute embedding query: %w", err)
	}

	status, err := job.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for embedding query: %w", err)
	}

	if status.Err() != nil {
		return fmt.Errorf("embedding query failed: %w", status.Err())
	}

	return nil
}

func (c *ConcreteEmbeddingsClient) buildEmbeddingQuery(tableConfig *TableConfig, cqIDs []string) (string, error) {
	// Build embed columns concatenated with newlines
	embedColumnsExpr := strings.Join(tableConfig.EmbedColumns, " || '\\n' || ")

	// Build metadata columns for SELECT
	metadataColumns := make([]string, len(tableConfig.MetadataColumns))
	copy(metadataColumns, tableConfig.MetadataColumns)
	metadataColumnsStr := strings.Join(metadataColumns, ", ")

	// Build WHERE clause for CQ IDs
	cqIDList := "'" + strings.Join(cqIDs, "', '") + "'"

	// Build the query template
	queryTemplate := `
INSERT INTO ` + "`{{.ProjectID}}.{{.DatasetID}}.{{.TargetTableName}}`" + ` ({{.MetadataColumns}}, chunk_id, chunk_text, embedding)
WITH docs AS (
  SELECT 
    {{.MetadataColumns}},
    {{.EmbedColumnsExpr}} AS content
  FROM ` + "`{{.ProjectID}}.{{.DatasetID}}.{{.SourceTableName}}`" + `
  WHERE _cq_id IN ({{.CqIDList}})
),
params AS (
  SELECT 
    {{.ChunkSize}} AS chunk_size,
    {{.ChunkOverlap}} AS overlap
),
offsets AS (
  SELECT 
    {{.MetadataColumns}},
    GENERATE_ARRAY(0, LENGTH(d.content) - 1, p.chunk_size - p.overlap) AS starts,
    d.content,
    p.chunk_size
  FROM docs d, params p
),
chunks AS (
  SELECT
    {{.MetadataColumns}},
    i AS chunk_id,
    SUBSTR(content, start + 1, chunk_size) AS chunk_text
  FROM offsets, UNNEST(starts) AS start WITH OFFSET i
),
embeddings AS (
  SELECT
    {{.MetadataColumns}},
    chunk_id,
    chunk_text,
    ml_generate_embedding_result AS embedding
  FROM ML.GENERATE_EMBEDDING(
    MODEL ` + "`{{.ProjectID}}.{{.DatasetID}}.{{.RemoteModelName}}`" + `,
    (
      SELECT 
        {{.MetadataColumns}},
        chunk_id,
        chunk_text,
        chunk_text AS content
      FROM chunks
    )
  )
)
SELECT
  {{.MetadataColumns}},
  chunk_id,
  chunk_text,
  embedding
FROM embeddings`

	// Prepare template data
	tmplData := struct {
		ProjectID        string
		DatasetID        string
		SourceTableName  string
		TargetTableName  string
		MetadataColumns  string
		EmbedColumnsExpr string
		CqIDList         string
		ChunkSize        int
		ChunkOverlap     int
		RemoteModelName  string
	}{
		ProjectID:        c.ProjectID,
		DatasetID:        c.DatasetID,
		SourceTableName:  tableConfig.SourceTableName,
		TargetTableName:  tableConfig.TargetTableName,
		MetadataColumns:  metadataColumnsStr,
		EmbedColumnsExpr: embedColumnsExpr,
		CqIDList:         cqIDList,
		ChunkSize:        c.spec.TextSplitter.RecursiveText.ChunkSize,
		ChunkOverlap:     c.spec.TextSplitter.RecursiveText.ChunkOverlap,
		RemoteModelName:  c.spec.RemoteModelName,
	}

	// Execute template
	tmpl, err := template.New("embedding_query").Parse(queryTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse query template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, tmplData); err != nil {
		return "", fmt.Errorf("failed to execute query template: %w", err)
	}

	return buf.String(), nil
}
