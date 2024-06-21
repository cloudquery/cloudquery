package client

import (
	_ "embed"
	"runtime"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 5 * 1024 * 1024
)

type Spec struct {
	// A list of Elasticsearch nodes to use.
	// Mutually exclusive with `cloud_id`.
	Addresses []string `json:"addresses" jsonschema:"default=http://localhost:9200"`

	// Username for HTTP Basic Authentication.
	Username string `json:"username"`

	// Password for HTTP Basic Authentication.
	Password string `json:"password"`

	// Endpoint for the Elasticsearch Service (https://elastic.co/cloud).
	// Mutually exclusive with `addresses`.
	CloudID string `json:"cloud_id" jsonschema:"example=MyDeployment:abcdefgh"`

	// Base64-encoded token for authorization; if set, overrides username/password and service token.
	APIKey string `json:"api_key"`

	// Service token for authorization; if set, overrides username/password.
	ServiceToken string `json:"service_token"`

	// SHA256 hex fingerprint given by Elasticsearch on first launch.
	CertificateFingerprint string `json:"certificate_fingerprint"`

	// PEM-encoded certificate authorities.
	// When set, an empty certificate pool will be created, and the certificates will be appended to it.
	CACert string `json:"ca_cert"`

	// Number of concurrent worker goroutines to use for indexing. (Default: number of CPUs)
	Concurrency int `json:"concurrency" jsonschema:"minimum=1"`

	// Number of documents to batch together per request.
	BatchSize int64 `json:"batch_size" jsonschema:"minimum=1,default=1000"`

	// Number of bytes to batch together per request.
	BatchSizeBytes int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=5242880"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if len(s.Addresses) == 0 && s.CloudID == "" {
		s.Addresses = []string{"http://localhost:9200"}
	}
	if s.Concurrency == 0 {
		s.Concurrency = runtime.NumCPU()
	}
	if s.BatchSize == 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
}

func (*Spec) Validate() error {
	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Not = &jsonschema.Schema{
		Description: "Either addresses or cloud_id must be set, but not both.",
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			one := uint64(1)
			properties := jsonschema.NewProperties()

			addresses := *sc.Properties.Value("addresses")
			addresses.MinLength = &one
			properties.Set("addresses", &addresses)

			cloudID := *sc.Properties.Value("cloud_id")
			cloudID.MinLength = &one

			properties.Set("cloud_id", &cloudID)

			return properties
		}(),
		Required: []string{"addresses", "cloud_id"},
	}
}
