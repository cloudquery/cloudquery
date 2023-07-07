package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

const (
	batchSize = 10000
	// documented BigQuery limit is 10MB, and we try to keep well below that as the size
	// estimate is not exact and there are also limits on request size, apart from the batch size
	batchSizeBytes = 5 * 1024 * 1024
)

type TimePartitioningOption string

const (
	TimePartitioningOptionNone = "none"
	TimePartitioningOptionHour = "hour"
	TimePartitioningOptionDay  = "day"
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
	ProjectID             string                 `json:"project_id"`
	DatasetID             string                 `json:"dataset_id"`
	DatasetLocation       string                 `json:"dataset_location"`
	TimePartitioning      TimePartitioningOption `json:"time_partitioning"`
	ServiceAccountKeyJSON string                 `json:"service_account_key_json"`
	BatchSize             int                    `json:"batch_size"`
	BatchSizeBytes        int                    `json:"batch_size_bytes"`
	BatchTimeout          configtype.Duration    `json:"batch_timeout"`
}

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
		return fmt.Errorf("project_id is required")
	}
	if s.DatasetID == "" {
		return fmt.Errorf("dataset_id is required")
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
