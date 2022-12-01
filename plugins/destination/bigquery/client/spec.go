package client

import (
	"fmt"
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
	ProjectID        string                 `json:"project_id"`
	DatasetID        string                 `json:"dataset_id"`
	TimePartitioning TimePartitioningOption `json:"time_partitioning"`
}

func (s *Spec) SetDefaults() {
	if s.TimePartitioning == "" {
		s.TimePartitioning = TimePartitioningOptionNone
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
	return nil
}
