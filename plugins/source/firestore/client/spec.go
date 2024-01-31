package client

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"github.com/invopop/jsonschema"
	"strings"

	"cloud.google.com/go/firestore"
)

// Spec is the (nested) spec used by Firestore Source Plugin
type Spec struct {
	// the ID of the project to use for this client. If not specified, the project id will be auto-detected from the credentials
	ProjectID string `json:"project_id"`
	// if true the service account JSON content will be treated as base64 encoded
	UseBase64 bool `json:"use_base64" jsonschema:"default=false"`
	//  service account JSON content
	ServiceAccountJSON string `json:"service_account_json"`
	// maximum batch size for each request when reading Firestore data
	MaxBatchSize int `json:"max_batch_size" jsonschema:"minimum=1"`
	// field(s) to order the results by
	OrderBy string `json:"order_by"`
	// the direction to order the results by when order_by is specified - accepts either asc or desc
	OrderDirection string `json:"order_direction" jsonschema:"enum=asc,enum=desc,default=asc"`
}

func (s *Spec) Validate() error {
	// decode base64 if needed - note if the Validate function is removed from the spec, this will need to be done
	// elsewhere in the application
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

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	// Configure defaults
	sc.Properties.Value("project_id").Default = firestore.DetectProjectID
	sc.Properties.Value("max_batch_size").Default = 50_000
}

func (s *Spec) SetDefaults() {
	if s.MaxBatchSize == 0 {
		s.MaxBatchSize = 50_000
	}
	if s.ProjectID == "" {
		s.ProjectID = firestore.DetectProjectID
	}
}

//go:embed schema.json
var JSONSchema string
