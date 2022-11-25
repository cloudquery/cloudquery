package client

import "fmt"

type Spec struct {
	ProjectID string `json:"project_id"`
	DatasetID string `json:"dataset_id"`
}

func (*Spec) SetDefaults() {
	// stub for any future defaults
}

func (s *Spec) Validate() error {
	if s.ProjectID == "" {
		return fmt.Errorf("project_id is required")
	}
	if s.DatasetID == "" {
		return fmt.Errorf("dataset_id is required")
	}
	return nil
}
