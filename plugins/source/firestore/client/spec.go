package client

import (
	"encoding/base64"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
)

type Spec struct {
	ProjectID          string `json:"project_id"`
	UseBase64          bool   `json:"use_base64"`
	ServiceAccountJSON string `json:"service_account_json"`
	MaxBatchSize       int    `json:"max_batch_size"`
	OrderBy            string `json:"order_by"`
	OrderDirection     string `json:"order_direction"`
}

func (s *Spec) Validate() error {
	if s.UseBase64 {
		data, err := base64.StdEncoding.DecodeString(s.ServiceAccountJSON)
		if err != nil {
			return fmt.Errorf("failed to decode service_account_json: %w", err)
		}
		s.ServiceAccountJSON = string(data)
	}
	s.OrderDirection = strings.ToLower(s.OrderDirection)
	if s.OrderDirection != "" && s.OrderDirection != "asc" && s.OrderDirection != "desc" {
		return fmt.Errorf("invalid order_by_direction %s", s.OrderDirection)
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
