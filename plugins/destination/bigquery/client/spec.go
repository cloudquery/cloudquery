package client

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
)

const (
	batchSize = 10000
	// documented BigQuery limit is 10MB, and we try to keep well below that as the size
	// estimate is not exact and there are also limits on request size, apart from the batch size
	batchSizeBytes = 5 * 1024 * 1024
)

type TimePartitioningOption string

const (
	TimePartitioningOptionNone = TimePartitioningOption("none")
	TimePartitioningOptionHour = TimePartitioningOption("hour")
	TimePartitioningOptionDay  = TimePartitioningOption("day")
)

var TimePartitioningOptions = []TimePartitioningOption{
	TimePartitioningOptionNone,
	TimePartitioningOptionHour,
	TimePartitioningOptionDay,
}

func (t TimePartitioningOption) Validate() error {
	for _, v := range TimePartitioningOptions {
		if t == v {
			return nil
		}
	}
	return fmt.Errorf("%v is not a valid option for time partitioning. Options are: %v", string(t), TimePartitioningOptions)
}

type Spec struct {
	// The id of the project where the destination BigQuery database resides.
	ProjectID string `json:"project_id" jsonschema:"required,minLength=1"`

	//  The name of the BigQuery dataset within the project, e.g. `my_dataset`.
	//  This dataset needs to be created before running a sync or migration.
	DatasetID string `json:"dataset_id" jsonschema:"required,minLength=1"`

	// The data location of the BigQuery dataset. If set, will be used as the default location for job operations.
	// Pro-tip: this can solve "dataset not found" issues for newly created datasets.
	DatasetLocation string `json:"dataset_location"`

	// The time partitioning to use when creating tables. The partition time column used will always be `_cq_sync_time` so that all rows for a sync run will be partitioned on the hour/day the sync started.
	TimePartitioning TimePartitioningOption `json:"time_partitioning"`

	// The duration to keep the partitions. Only applicable if `time_partitioning` is set to a value other than `none`. A value of 0 means no expiration.
	TimePartitioningExpiration configtype.Duration `json:"time_partitioning_expiration" jsonschema:"minimum=0,default=0"`

	// GCP service account key content.
	// This allows for using different service accounts for the GCP source and BigQuery destination.
	// If using service account keys, it is best to use [environment or file variable substitution](/docs/advanced-topics/environment-variable-substitution).
	ServiceAccountKeyJSON string `json:"service_account_key_json"`

	// The BigQuery API endpoint to use. This is useful for testing against a local emulator.
	Endpoint string `json:"endpoint"`

	// Number of records to write before starting a new object.
	BatchSize int64 `json:"batch_size" jsonschema:"minimum=1,default=10000"`

	// Number of bytes (as Arrow buffer size) to write before starting a new object.
	BatchSizeBytes int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=5242880"`

	// Maximum interval between batch writes.
	BatchTimeout configtype.Duration `json:"batch_timeout"`

	// Identifies the project context bq client should execute in. Defaults to the project_id. You can set it to *detect-project-id* to automatically detect project id from credentials in the environment.
	ClientProjectID string `json:"client_project_id"`

	// Optional: Configuration for creating text embeddings for certain tables
	TextEmbeddings *TextEmbeddingsSpec `json:"text_embeddings,omitempty"`
}

type TextEmbeddingsSpec struct {
	RemoteModelName string        `json:"remote_model_name" jsonschema:"required,minLength=1"`
	Tables          []TableConfig `json:"tables,omitempty"`
	TextSplitter    TextSplitter  `json:"text_splitter"`
}

// TableConfig defines per-source-table embedding configuration.
// SourceTableName is the base/source table from which text columns will be embedded.
// TargetTableName is the destination table that will be created to store embeddings
// and metadata columns.
type TableConfig struct {
	SourceTableName string   `json:"source_table_name" jsonschema:"required,minLength=1"`
	TargetTableName string   `json:"target_table_name" jsonschema:"required,minLength=1"`
	EmbedColumns    []string `json:"embed_columns" jsonschema:"required,minItems=1"`
	MetadataColumns []string `json:"metadata_columns,omitempty"`
}

// TextSplitter defines how source text should be split into chunks for embedding.
type TextSplitter struct {
	RecursiveText RecursiveText `json:"recursive_text" jsonschema:"required"`
}

type RecursiveText struct {
	ChunkSize    int `json:"chunk_size" jsonschema:"required,minimum=1"`
	ChunkOverlap int `json:"chunk_overlap" jsonschema:"required,minimum=0"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if s.TimePartitioning == "" {
		s.TimePartitioning = TimePartitioningOptionNone
	}
	if s.BatchSize == 0 {
		s.BatchSize = batchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = batchSizeBytes
	}
	if s.BatchTimeout.Duration() <= 0 {
		s.BatchTimeout = configtype.NewDuration(10 * time.Second)
	}
	if s.ClientProjectID == "" {
		s.ClientProjectID = s.ProjectID
	}
	if s.TextEmbeddings != nil {
		s.TextEmbeddings.SetDefaults()
	}
}

func (s *Spec) Validate() error {
	if s.ProjectID == "" {
		return errors.New("project_id is required")
	}
	if s.DatasetID == "" {
		return errors.New("dataset_id is required")
	}
	if err := s.TimePartitioning.Validate(); err != nil {
		return fmt.Errorf("time_partitioning: %w", err)
	}
	if s.TimePartitioning != TimePartitioningOptionNone && s.TimePartitioningExpiration.Duration() > 0 {
		return errors.New("time_partitioning_expiration option requires time_partitioning to be set")
	}
	if len(s.ServiceAccountKeyJSON) > 0 {
		if err := isValidJson(s.ServiceAccountKeyJSON); err != nil {
			return fmt.Errorf("invalid json for service_account_key_json: %w", err)
		}
	}
	if s.TextEmbeddings != nil {
		if err := s.TextEmbeddings.Validate(); err != nil {
			return fmt.Errorf("text_embeddings: %w", err)
		}
	}
	return nil
}

func isValidJson(content string) error {
	var v map[string]any
	err := json.Unmarshal([]byte(content), &v)
	if err != nil {
		return err
	}
	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Value("batch_timeout").Default = "10s"
}

func (TimePartitioningOption) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Default = TimePartitioningOptionNone
	sc.Enum = make([]any, len(TimePartitioningOptions))
	for i := range TimePartitioningOptions {
		sc.Enum[i] = TimePartitioningOptions[i]
	}
}

// SetDefaults sets default values for TextEmbeddingsSpec
func (t *TextEmbeddingsSpec) SetDefaults() {
	// Set default values for TextSplitter if not set
	if t.TextSplitter.RecursiveText.ChunkSize == 0 {
		t.TextSplitter.RecursiveText.ChunkSize = 1000
	}
	if t.TextSplitter.RecursiveText.ChunkOverlap == 0 {
		t.TextSplitter.RecursiveText.ChunkOverlap = 100
	}

	// Ensure _cq_id is in metadata columns for each table if not already present
	for i := range t.Tables {
		hasCqID := slices.Contains(t.Tables[i].MetadataColumns, "_cq_id")
		if !hasCqID {
			t.Tables[i].MetadataColumns = append(t.Tables[i].MetadataColumns, "_cq_id")
		}
	}
}

// Validate validates the TextEmbeddingsSpec configuration
func (t *TextEmbeddingsSpec) Validate() error {
	if t.RemoteModelName == "" {
		return errors.New("remote_model_name is required when text_embeddings is set")
	}

	if len(t.Tables) == 0 {
		return errors.New("at least one table must be defined when text_embeddings is set")
	}

	// Validate each table configuration
	for i, table := range t.Tables {
		if table.SourceTableName == "" {
			return fmt.Errorf("table[%d]: source_table_name cannot be empty", i)
		}
		if table.TargetTableName == "" {
			return fmt.Errorf("table[%d]: target_table_name cannot be empty", i)
		}
		if len(table.EmbedColumns) == 0 {
			return fmt.Errorf("table[%d]: at least one embed_column must be set", i)
		}
	}

	// Validate TextSplitter if set
	if t.TextSplitter.RecursiveText.ChunkSize <= 0 {
		return errors.New("text_splitter.recursive_text.chunk_size must be greater than 0")
	}
	if t.TextSplitter.RecursiveText.ChunkOverlap < 0 {
		return errors.New("text_splitter.recursive_text.chunk_overlap must be non-negative")
	}
	if t.TextSplitter.RecursiveText.ChunkOverlap >= t.TextSplitter.RecursiveText.ChunkSize {
		return errors.New("text_splitter.recursive_text.chunk_overlap must be less than chunk_size")
	}

	return nil
}
