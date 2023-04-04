package client

import (
	"fmt"

	"cloud.google.com/go/firestore"
)

type Spec struct {
	ProjectID          string `json:"project_id"`
	ServiceAccountJSON string `json:"service_account_json"`
	MaxBatchSize       int    `json:"max_batch_size"`
	OrderByField       string `json:"order_by_field"`
	OrderByDirection   string `json:"order_by_direction"`
}

func (s *Spec) Validate() error {
	if s.OrderByDirection != "" && s.OrderByDirection != "asc" && s.OrderByDirection != "desc" {
		return fmt.Errorf("invalid order_by_direction %s", s.OrderByDirection)
	}
	if s.MaxBatchSize < 0 {
		return fmt.Errorf("max_batch_size must be greater than 0")
	}
	return nil
}

func (s *Spec) SetDefaults() {
	if s.MaxBatchSize == 0 {
		s.MaxBatchSize = 50_000
	}
	if s.ProjectID == "" {
		s.ProjectID = firestore.DetectProjectID
	}
}
