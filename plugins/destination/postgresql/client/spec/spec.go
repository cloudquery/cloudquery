package spec

import (
	_ "embed"
	"errors"
	"slices"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
	"github.com/jackc/pgx/v5/tracelog"
)

const (
	defaultBatchSize      = 10000
	defaultBatchSizeBytes = 100000000
	defaultBatchTimeout   = 60 * time.Second
	CQIDColumn            = "_cq_id"
)

type Spec struct {
	// Connection string to connect to the database. This can be a DSN or a URI, for example:
	//
	// - `"user=user password=pass host=localhost port=5432 dbname=mydb sslmode=disable"` DSN format
	// - `"postgres://user:pass@localhost:5432/mydb?sslmode=prefer"` connect with tcp and prefer TLS
	// - `"postgres://user:pass@localhost:5432/mydb?sslmode=disable&search_path=myschema"` connect with tcp, disable TLS and use a custom schema
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1,example=${POSTGRESQL_CONNECTION_STRING}"`

	// Available: `error`, `warn`, `info`, `debug`, `trace`.
	// Defines what [`pgx`](https://github.com/jackc/pgx) call events should be logged.
	PgxLogLevel LogLevel `json:"pgx_log_level,omitempty" jsonschema:"default=error"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=100000000"`

	// Maximum interval between batch writes.
	BatchTimeout configtype.Duration `json:"batch_timeout,omitempty"`

	// Option to create specific indexes to improve deletion performance
	CreatePerformanceIndexes bool `json:"create_performance_indexes,omitempty" jsonschema:"default=false"`

	// Optional configuration to enable PgVector embedding support.
	PgVectorConfig *PgVectorConfig `json:"pgvector_config,omitempty"`
}

func (s *Spec) HasPgVectorConfig() bool {
	return s.PgVectorConfig != nil
}

func (s *Spec) SetDefaults() {
	if s.BatchSize <= 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes <= 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
	if s.BatchTimeout.Duration() <= 0 {
		s.BatchTimeout = configtype.NewDuration(defaultBatchTimeout)
	}
	if s.PgxLogLevel == 0 {
		s.PgxLogLevel = LogLevel(tracelog.LogLevelError)
	}
	if s.PgVectorConfig != nil {
		for i := range s.PgVectorConfig.Tables {
			s.PgVectorConfig.Tables[i].MetadataColumns = ensureCQIDPresent(s.PgVectorConfig.Tables[i].MetadataColumns)
		}
		if s.PgVectorConfig.TextSplitter == nil {
			s.PgVectorConfig.TextSplitter = &PgVectorTextSplitter{
				RecursiveText: PgVectorRecursiveText{
					ChunkSize:    1000,
					ChunkOverlap: 500,
				},
			}
		}
	}
}

func ensureCQIDPresent(metadataColumns []string) []string {
	if slices.Contains(metadataColumns, CQIDColumn) {
		return metadataColumns
	}
	return append([]string{CQIDColumn}, metadataColumns...)
}

func embeddingDimensionsForModel(model string) (int, error) {
	switch model {
	case "text-embedding-3-small":
		return 1536, nil
	case "text-embedding-3-large":
		return 3072, nil
	default:
		return 0, errors.New("`pgvector_config.embedding.model_name` must be one of: text-embedding-3-small, text-embedding-3-large")
	}
}

func (s *Spec) Validate() error {
	if len(s.ConnectionString) == 0 {
		return errors.New("`connection_string` is required")
	}
	if s.PgVectorConfig != nil {
		if len(s.PgVectorConfig.Tables) == 0 {
			return errors.New("`pgvector_config.tables` must contain at least 1 table")
		}
		seenNames := make(map[string]struct{}, len(s.PgVectorConfig.Tables))
		for _, tbl := range s.PgVectorConfig.Tables {
			if len(tbl.TableName) == 0 {
				return errors.New("`pgvector_config.tables.table_name` is required")
			}
			if _, ok := seenNames[tbl.TableName]; ok {
				return errors.New("`pgvector_config.tables` contains duplicate table names: " + tbl.TableName)
			}
			seenNames[tbl.TableName] = struct{}{}
			if len(tbl.EmbedColumns) == 0 {
				return errors.New("`pgvector_config.tables.embed_columns` must contain at least 1 column")
			}
			if len(tbl.MetadataColumns) == 0 {
				return errors.New("`pgvector_config.tables.metadata_columns` must contain at least 1 column")
			}
		}
		emb := s.PgVectorConfig.Embedding
		if emb.Dimensions <= 0 || len(emb.APIKey) == 0 || len(emb.ModelName) == 0 {
			return errors.New("`pgvector_config.embedding` must have `dimensions`, `api_key`, and `model_name` set")
		}
		// Enforce model support and sync dimensions to the selected model
		dims, err := embeddingDimensionsForModel(emb.ModelName)
		if err != nil {
			return err
		}
		s.PgVectorConfig.Embedding.Dimensions = dims
		if s.PgVectorConfig.TextSplitter != nil {
			ts := s.PgVectorConfig.TextSplitter
			if ts.RecursiveText.ChunkSize <= 0 {
				return errors.New("`pgvector_config.text_splitter.recursive_text.chunk_size` must be > 0")
			}
			if ts.RecursiveText.ChunkOverlap < 0 {
				return errors.New("`pgvector_config.text_splitter.recursive_text.chunk_overlap` must be >= 0")
			}
		}
	}
	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Value("batch_timeout").Default = "60s"
}

//go:embed schema.json
var JSONSchema string

// PgVectorConfig holds configuration for creating embeddings and storing them with pgvector.
type PgVectorConfig struct {
	// Tables to create embeddings for.
	Tables []PgVectorTableConfig `json:"tables,omitempty" jsonschema:"required,minItems=1"`
	// Optional text splitting configuration. If set, all sub-configurations must be set.
	TextSplitter *PgVectorTextSplitter `json:"text_splitter,omitempty"`
	// Embedding provider configuration. Required if PgVectorConfig is set.
	Embedding PgVectorEmbedding `json:"embedding" jsonschema:"required"`
}

// PgVectorTableConfig defines per-source-table embedding configuration.
type PgVectorTableConfig struct {
	TableName       string   `json:"table_name" jsonschema:"required,minLength=1"`
	EmbedColumns    []string `json:"embed_columns" jsonschema:"required,minItems=1"`
	MetadataColumns []string `json:"metadata_columns" jsonschema:"required,minItems=1"`
}

// PgVectorTextSplitter defines how source text should be split into chunks for embedding.
type PgVectorTextSplitter struct {
	RecursiveText PgVectorRecursiveText `json:"recursive_text" jsonschema:"required"`
}

type PgVectorRecursiveText struct {
	ChunkSize    int `json:"chunk_size" jsonschema:"required,minimum=1"`
	ChunkOverlap int `json:"chunk_overlap" jsonschema:"required,minimum=0"`
}

// PgVectorEmbedding holds embedding provider settings.
type PgVectorEmbedding struct {
	Dimensions int    `json:"dimensions" jsonschema:"minimum=1"`
	APIKey     string `json:"api_key" jsonschema:"required,minLength=1,title=OpenAI API Key"`
	ModelName  string `json:"model_name" jsonschema:"required,minLength=1"`
}

func (s *Spec) GetPgVectorTableConfig(tableName string) *PgVectorTableConfig {
	if s.PgVectorConfig == nil {
		return nil
	}
	for _, tbl := range s.PgVectorConfig.Tables {
		if tbl.TableName == tableName {
			return &tbl
		}
	}
	return nil
}
