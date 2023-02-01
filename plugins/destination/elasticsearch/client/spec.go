package client

import "runtime"

type Spec struct {
	Addresses []string `json:"addresses"` // A list of Elasticsearch nodes to use.
	Username  string   `json:"username"`  // Username for HTTP Basic Authentication.
	Password  string   `json:"password"`  // Password for HTTP Basic Authentication.

	CloudID                string `json:"cloud_id"`                // Endpoint for the Elastic Service (https://elastic.co/cloud).
	APIKey                 string `json:"api_key"`                 // Base64-encoded token for authorization; if set, overrides username/password and service token.
	ServiceToken           string `json:"service_token"`           // Service token for authorization; if set, overrides username/password.
	CertificateFingerprint string `json:"certificate_fingerprint"` // SHA256 hex fingerprint given by Elasticsearch on first launch.

	// PEM-encoded certificate authorities.
	// When set, an empty certificate pool will be created, and the certificates will be appended to it.
	CACert string `json:"ca_cert"`

	Concurrency int `json:"concurrency"` // Number of concurrent worker goroutines to use for indexing. (Default: number of CPUs)
}

func (s *Spec) SetDefaults() {
	if len(s.Addresses) == 0 {
		s.Addresses = []string{"http://localhost:9200"}
	}
	if s.Concurrency == 0 {
		s.Concurrency = runtime.NumCPU()
	}
}

func (*Spec) Validate() error {
	return nil
}
