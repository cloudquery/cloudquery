package client

import (
	_ "embed"
	"runtime"
)

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 5 * 1024 * 1024
)

type Spec struct {
	// A list of Elasticsearch nodes to use.
	Addresses []string `json:"addresses" jsonschema:"default=http://localhost:9200"`

	// Username for HTTP Basic Authentication.
	Username string `json:"username" jsonschema_extras:"x-cq-auth=true"`

	// Password for HTTP Basic Authentication.
	Password string `json:"password" jsonschema_extras:"x-cq-auth=true"`

	// Endpoint for the Elastic Service (https://elastic.co/cloud).
	CloudID string `json:"cloud_id"`

	// Base64-encoded token for authorization; if set, overrides username/password and service token.
	APIKey string `json:"api_key" jsonschema_extras:"x-cq-auth=true"`

	// Service token for authorization; if set, overrides username/password.
	ServiceToken string `json:"service_token" jsonschema_extras:"x-cq-auth=true"`

	// SHA256 hex fingerprint given by Elasticsearch on first launch.
	CertificateFingerprint string `json:"certificate_fingerprint" jsonschema_extras:"x-cq-auth=true"`

	// PEM-encoded certificate authorities.
	// When set, an empty certificate pool will be created, and the certificates will be appended to it.
	CACert string `json:"ca_cert" jsonschema_extras:"x-cq-auth=true"`

	// Number of concurrent worker goroutines to use for indexing. (Default: number of CPUs)
	Concurrency int `json:"concurrency" jsonschema:"minimum=1"`

	// Number of documents to batch together per request.
	BatchSize int `json:"batch_size" jsonschema:"minimum=1,default=1000"`

	// Number of bytes to batch together per request.
	BatchSizeBytes int `json:"batch_size_bytes" jsonschema:"minimum=1,default=5242880"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if len(s.Addresses) == 0 {
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
