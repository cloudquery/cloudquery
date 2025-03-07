package client

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
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
	if len(s.ServiceAccountKeyJSON) > 0 {
		if err := isValidJson(s.ServiceAccountKeyJSON); err != nil {
			return fmt.Errorf("invalid json for service_account_key_json: %w", err)
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
