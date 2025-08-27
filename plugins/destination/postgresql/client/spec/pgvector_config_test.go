package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func baseSpec() *Spec {
	return &Spec{ConnectionString: "abc"}
}

func validPgVectorConfig() *PgVectorConfig {
	return &PgVectorConfig{
		Tables: []PgVectorTableConfig{
			{
				SourceTableName: "box_file_contents",
				TargetTableName: "box_file_contents_embeddings",
				EmbedColumns:    []string{"content"},
				MetadataColumns: []string{"id"},
			},
		},
		OpenAIEmbedding: OpenAIEmbedding{
			Dimensions: 1536,
			APIKey:     "key",
			ModelName:  "text-embedding-3-small",
		},
	}
}

func TestPgVectorConfig_Omitted_OK(t *testing.T) {
	s := baseSpec()
	require.NoError(t, s.Validate())
}

func TestPgVectorConfig_EmptyTables_Error(t *testing.T) {
	s := baseSpec()
	s.PgVectorConfig = &PgVectorConfig{OpenAIEmbedding: OpenAIEmbedding{Dimensions: 1, APIKey: "k", ModelName: "text-embedding-3-small"}}
	err := s.Validate()
	require.Error(t, err)
}

func TestPgVectorConfig_DuplicateTables_Error(t *testing.T) {
	s := baseSpec()
	cfg := validPgVectorConfig()
	cfg.Tables = append(cfg.Tables, PgVectorTableConfig{SourceTableName: cfg.Tables[0].SourceTableName, TargetTableName: "another", EmbedColumns: []string{"c"}, MetadataColumns: []string{"m"}})
	s.PgVectorConfig = cfg
	err := s.Validate()
	require.Error(t, err)

	cfg = validPgVectorConfig()
	cfg.Tables = append(cfg.Tables, PgVectorTableConfig{SourceTableName: "another", TargetTableName: cfg.Tables[0].TargetTableName, EmbedColumns: []string{"c"}, MetadataColumns: []string{"m"}})
	s.PgVectorConfig = cfg
	err = s.Validate()
	require.Error(t, err)
}

func TestPgVectorConfig_TableMissingColumns_Error(t *testing.T) {
	s := baseSpec()
	cfg := validPgVectorConfig()
	cfg.Tables[0].EmbedColumns = nil
	s.PgVectorConfig = cfg
	require.Error(t, s.Validate())

	cfg = validPgVectorConfig()
	cfg.Tables[0].MetadataColumns = nil // Now optional, should not error
	s.PgVectorConfig = cfg
	require.NoError(t, s.Validate())
}

func TestPgVectorConfig_EmbeddingMissingFields_Error(t *testing.T) {
	s := baseSpec()
	cfg := validPgVectorConfig()
	cfg.OpenAIEmbedding.APIKey = ""
	s.PgVectorConfig = cfg
	require.Error(t, s.Validate())

	cfg = validPgVectorConfig()
	cfg.OpenAIEmbedding.ModelName = ""
	s.PgVectorConfig = cfg
	require.Error(t, s.Validate())

	cfg = validPgVectorConfig()
	cfg.OpenAIEmbedding.Dimensions = 0
	s.PgVectorConfig = cfg
	require.Error(t, s.Validate())
}

func TestPgVectorConfig_TextSplitter_AllOrNothing_Error(t *testing.T) {
	s := baseSpec()
	cfg := validPgVectorConfig()
	cfg.TextSplitter = &PgVectorTextSplitter{RecursiveText: PgVectorRecursiveText{ChunkSize: 0, ChunkOverlap: 0}}
	s.PgVectorConfig = cfg
	require.Error(t, s.Validate())

	cfg = validPgVectorConfig()
	cfg.TextSplitter = &PgVectorTextSplitter{RecursiveText: PgVectorRecursiveText{ChunkSize: 1000, ChunkOverlap: -1}}
	s.PgVectorConfig = cfg
	require.Error(t, s.Validate())
}

func TestPgVectorConfig_Valid_OK(t *testing.T) {
	s := baseSpec()
	cfg := validPgVectorConfig()
	cfg.TextSplitter = &PgVectorTextSplitter{RecursiveText: PgVectorRecursiveText{ChunkSize: 1000, ChunkOverlap: 500}}
	s.PgVectorConfig = cfg
	require.NoError(t, s.Validate())
	// Ensure dimensions align to model choice
	require.Equal(t, 1536, s.PgVectorConfig.OpenAIEmbedding.Dimensions)

	cfg = validPgVectorConfig()
	cfg.OpenAIEmbedding.ModelName = "text-embedding-3-large"
	s.PgVectorConfig = cfg
	require.NoError(t, s.Validate())
	require.Equal(t, 3072, s.PgVectorConfig.OpenAIEmbedding.Dimensions)
}

func TestPgVectorConfig_DefaultTextSplitter_OnSetDefaults(t *testing.T) {
	s := baseSpec()
	cfg := validPgVectorConfig()
	cfg.TextSplitter = nil
	s.PgVectorConfig = cfg
	s.SetDefaults()
	require.NotNil(t, s.PgVectorConfig.TextSplitter)
	require.Equal(t, 1000, s.PgVectorConfig.TextSplitter.RecursiveText.ChunkSize)
	require.Equal(t, 500, s.PgVectorConfig.TextSplitter.RecursiveText.ChunkOverlap)
}
