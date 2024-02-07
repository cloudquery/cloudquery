package client

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/invopop/jsonschema"
)

// Spec is the (nested) spec used by Firestore Source Plugin
type Spec struct {
	// The ID of the project to use for this client. If not specified, the project id will be auto-detected from the credentials.
	ProjectID string `json:"project_id"`
	// If `true` the `service_account_json` content will be treated as base64-encoded.
	UseBase64 bool `json:"use_base64" jsonschema:"default=false"`
	//  Service account JSON content.
	ServiceAccountJSON string `json:"service_account_json" jsonschema_extras:"x-cq-auth=true"`
	// Maximum batch size for each request when reading Firestore data.
	MaxBatchSize int `json:"max_batch_size" jsonschema:"minimum=1"`
	// List of fields to order the results by.
	OrderBy string `json:"order_by"`
	// The order direction used when `order_by` is `true`.
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
