package client

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/stretchr/testify/require"
)

func newTestEmbeddingsClient() *ConcreteEmbeddingsClient {
	spec := &TextEmbeddingsSpec{RemoteModelName: "embedding_model"}
	spec.SetDefaults()
	return &ConcreteEmbeddingsClient{
		spec:      spec,
		ProjectID: "my-project",
		DatasetID: "my_dataset",
	}
}

func newTestTableConfig() *TableConfig {
	return &TableConfig{
		SourceTableName: "source_table",
		TargetTableName: "target_table",
		EmbedColumns:    []string{"title", "body"},
		MetadataColumns: []string{"_cq_id", "name"},
	}
}

func TestBuildEmbeddingQueryBindsCqIDsAsParameter(t *testing.T) {
	cqIDs := []string{
		"7b0e6b0c-1f5c-4c53-9d5e-0a8f4f6f1c11",
		"x'); DROP TABLE `my-project.my_dataset.source_table`; --",
	}

	sql, parameters, err := newTestEmbeddingsClient().buildEmbeddingQuery(newTestTableConfig(), cqIDs)
	require.NoError(t, err)

	for _, cqID := range cqIDs {
		require.NotContains(t, sql, cqID)
	}
	require.NotContains(t, sql, "DROP TABLE")
	require.Contains(t, sql, "WHERE _cq_id IN UNNEST(@cq_ids)")
	require.Equal(t, []bigquery.QueryParameter{{Name: "cq_ids", Value: cqIDs}}, parameters)
}

func TestBuildEmbeddingQueryKeepsIdentifiers(t *testing.T) {
	sql, _, err := newTestEmbeddingsClient().buildEmbeddingQuery(newTestTableConfig(), []string{"id"})
	require.NoError(t, err)

	require.Contains(t, sql, "INSERT INTO `my-project.my_dataset.target_table` (_cq_id, name, chunk_id, chunk_text, embedding)")
	require.Contains(t, sql, "FROM `my-project.my_dataset.source_table`")
	require.Contains(t, sql, "MODEL `my-project.my_dataset.embedding_model`")
	require.Contains(t, sql, `title || '\n' || body AS content`)
	require.Contains(t, sql, "1000 AS chunk_size")
	require.Contains(t, sql, "100 AS overlap")
}

func TestBuildEmbeddingQueryRejectsUnsafeIdentifiers(t *testing.T) {
	const injection = "x` WHERE FALSE; DROP TABLE t; --"
	cases := map[string]func(c *ConcreteEmbeddingsClient, tableConfig *TableConfig){
		"source table":    func(_ *ConcreteEmbeddingsClient, tc *TableConfig) { tc.SourceTableName = injection },
		"target table":    func(_ *ConcreteEmbeddingsClient, tc *TableConfig) { tc.TargetTableName = injection },
		"embed column":    func(_ *ConcreteEmbeddingsClient, tc *TableConfig) { tc.EmbedColumns = []string{"title", "(SELECT 1)"} },
		"metadata column": func(_ *ConcreteEmbeddingsClient, tc *TableConfig) { tc.MetadataColumns = []string{"_cq_id, 1"} },
		"project":         func(c *ConcreteEmbeddingsClient, _ *TableConfig) { c.ProjectID = injection },
		"dataset":         func(c *ConcreteEmbeddingsClient, _ *TableConfig) { c.DatasetID = injection },
		"remote model":    func(c *ConcreteEmbeddingsClient, _ *TableConfig) { c.spec.RemoteModelName = injection },
	}
	for description, mutate := range cases {
		t.Run(description, func(t *testing.T) {
			c := newTestEmbeddingsClient()
			tableConfig := newTestTableConfig()
			mutate(c, tableConfig)

			sql, parameters, err := c.buildEmbeddingQuery(tableConfig, []string{"id"})
			require.ErrorContains(t, err, "invalid")
			require.Empty(t, sql)
			require.Nil(t, parameters)
		})
	}
}
